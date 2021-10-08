/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package planbuilder

import (
	vtrpcpb "vitess.io/vitess/go/vt/proto/vtrpc"
	"vitess.io/vitess/go/vt/sqlparser"
	"vitess.io/vitess/go/vt/vterrors"
	"vitess.io/vitess/go/vt/vtgate/engine"
	"vitess.io/vitess/go/vt/vtgate/semantics"
)

type rewriter struct {
	semTable     *semantics.SemTable
	reservedVars *sqlparser.ReservedVars
	isInSubquery int
	err          error
}

func queryRewrite(semTable *semantics.SemTable, reservedVars *sqlparser.ReservedVars, statement sqlparser.SelectStatement) error {
	r := rewriter{
		semTable:     semTable,
		reservedVars: reservedVars,
	}
	sqlparser.Rewrite(statement, r.rewriteDown, r.rewriteUp)
	return nil
}

func (r *rewriter) rewriteDown(cursor *sqlparser.Cursor) bool {
	switch node := cursor.Node().(type) {
	case *sqlparser.Select:
		rewriteHavingClause(node)
	case *sqlparser.ComparisonExpr:
		err := rewriteInSubquery(cursor, r, node)
		if err != nil {
			r.err = err
		}
	case *sqlparser.ExistsExpr:
		err := r.rewriteExistsSubquery(cursor, node)
		if err != nil {
			r.err = err
		}
		return false
	case *sqlparser.AliasedTableExpr:
		// rewrite names of the routed tables for the subquery
		// We only need to do this for non-derived tables and if they are in a subquery
		if _, isDerived := node.Expr.(*sqlparser.DerivedTable); isDerived || r.isInSubquery == 0 {
			break
		}
		// find the tableSet and tableInfo that this table points to
		// tableInfo should contain the information for the original table that the routed table points to
		tableSet := r.semTable.TableSetFor(node)
		tableInfo, err := r.semTable.TableInfoFor(tableSet)
		if err != nil {
			// Fail-safe code, should never happen
			break
		}
		// vindexTable is the original table
		vindexTable := tableInfo.GetVindexTable()
		if vindexTable == nil {
			break
		}
		tableName := node.Expr.(sqlparser.TableName)
		// if the table name matches what the original is, then we do not need to rewrite
		if sqlparser.EqualsTableIdent(vindexTable.Name, tableName.Name) {
			break
		}
		// if there is no as clause, then move the routed table to the as clause.
		// i.e
		// routed as x -> original as x
		// routed -> original as routed
		if node.As.IsEmpty() {
			node.As = tableName.Name
		}
		// replace the table name with the original table
		tableName.Name = vindexTable.Name
		node.Expr = tableName
	}
	return true
}

func (r *rewriter) rewriteUp(cursor *sqlparser.Cursor) bool {
	switch cursor.Node().(type) {
	case *sqlparser.Subquery:
		r.isInSubquery--
	}
	return r.err == nil
}

func rewriteInSubquery(cursor *sqlparser.Cursor, r *rewriter, node *sqlparser.ComparisonExpr) error {
	if node.Operator != sqlparser.InOp &&
		node.Operator != sqlparser.NotInOp &&
		node.Operator != sqlparser.EqualOp &&
		node.Operator != sqlparser.NotEqualOp {
		return nil
	}
	subq, exp := getSubqueryAndOtherSide(node)
	if subq == nil || exp == nil {
		return nil
	}

	semTableSQ, found := r.semTable.SubqueryRef[subq]
	if !found {
		return vterrors.Errorf(vtrpcpb.Code_INTERNAL, "BUG: came across subquery that was not in the subq map")
	}

	argName, hasValuesArg := r.reservedVars.ReserveSubQueryWithHasValues()
	semTableSQ.ArgName = argName
	semTableSQ.HasValuesArg = hasValuesArg
	semTableSQ.OtherSide = exp
	cursor.Replace(semTableSQ)
	return nil
}

func getSubqueryAndOtherSide(node *sqlparser.ComparisonExpr) (*sqlparser.Subquery, sqlparser.Expr) {
	var subq *sqlparser.Subquery
	var exp sqlparser.Expr
	if lSubq, lIsSubq := node.Left.(*sqlparser.Subquery); lIsSubq {
		subq = lSubq
		exp = node.Right
	} else if rSubq, rIsSubq := node.Right.(*sqlparser.Subquery); rIsSubq {
		subq = rSubq
		exp = node.Left
	}
	return subq, exp
}

func rewriteSubquery(cursor *sqlparser.Cursor, r *rewriter, node *sqlparser.Subquery) error {
	semTableSQ, found := r.semTable.SubqueryRef[node]
	if !found {
		return vterrors.Errorf(vtrpcpb.Code_INTERNAL, "BUG: came across subquery that was not in the subq map")
	}
	if engine.PulloutOpcode(semTableSQ.OpCode) != engine.PulloutValue {
		return nil
	}

	argName := r.reservedVars.ReserveSubQuery()
	semTableSQ.ArgName = argName
	cursor.Replace(semTableSQ)
	return nil
}

func (r *rewriter) rewriteExistsSubquery(cursor *sqlparser.Cursor, node *sqlparser.ExistsExpr) error {
	semTableSQ, found := r.semTable.SubqueryRef[node.Subquery]
	if !found {
		return vterrors.Errorf(vtrpcpb.Code_INTERNAL, "BUG: came across subquery that was not in the subq map")
	}

	argName := r.reservedVars.ReserveHasValuesSubQuery()
	semTableSQ.ArgName = argName
	cursor.Replace(semTableSQ)
	return nil
}

func rewriteHavingClause(node *sqlparser.Select) {
	if node.Having == nil {
		return
	}

	selectExprMap := map[string]sqlparser.Expr{}
	for _, selectExpr := range node.SelectExprs {
		aliasedExpr, isAliased := selectExpr.(*sqlparser.AliasedExpr)
		if !isAliased || aliasedExpr.As.IsEmpty() {
			continue
		}
		selectExprMap[aliasedExpr.As.Lowered()] = aliasedExpr.Expr
	}

	sqlparser.Rewrite(node.Having.Expr, func(cursor *sqlparser.Cursor) bool {
		switch x := cursor.Node().(type) {
		case *sqlparser.ColName:
			if !x.Qualifier.IsEmpty() {
				return false
			}
			originalExpr, isInMap := selectExprMap[x.Name.Lowered()]
			if isInMap {
				cursor.Replace(originalExpr)
				return false
			}
			return false
		}
		return true
	}, nil)

	exprs := sqlparser.SplitAndExpression(nil, node.Having.Expr)
	node.Having = nil
	for _, expr := range exprs {
		if sqlparser.ContainsAggregation(expr) {
			node.AddHaving(expr)
		} else {
			node.AddWhere(expr)
		}
	}
}
