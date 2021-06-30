/*
Copyright 2019 The Vitess Authors.

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
	"vitess.io/vitess/go/sqltypes"

	"vitess.io/vitess/go/vt/vtgate/semantics"

	vtrpcpb "vitess.io/vitess/go/vt/proto/vtrpc"
	"vitess.io/vitess/go/vt/vterrors"

	"vitess.io/vitess/go/vt/sqlparser"
	"vitess.io/vitess/go/vt/vtgate/engine"
)

func pushProjection(expr *sqlparser.AliasedExpr, plan logicalPlan, semTable *semantics.SemTable, inner bool) (int, error) {
	switch node := plan.(type) {
	case *route:
		value, err := makePlanValue(expr.Expr)
		if err != nil {
			return 0, err
		}
		_, isColName := expr.Expr.(*sqlparser.ColName)
		badExpr := value == nil && !isColName
		if !inner && badExpr {
			return 0, vterrors.New(vtrpcpb.Code_UNIMPLEMENTED, "unsupported: cross-shard left join and column expressions")
		}
		sel := node.Select.(*sqlparser.Select)
		i := checkIfAlreadyExists(expr, sel)
		if i != -1 {
			return i, nil
		}
		expr = removeQualifierFromColName(expr)

		offset := len(sel.SelectExprs)
		sel.SelectExprs = append(sel.SelectExprs, expr)
		return offset, nil
	case *joinGen4:
		lhsSolves := node.Left.ContainsTables()
		rhsSolves := node.Right.ContainsTables()
		deps := semTable.Dependencies(expr.Expr)
		switch {
		case deps.IsSolvedBy(lhsSolves):
			offset, err := pushProjection(expr, node.Left, semTable, inner)
			if err != nil {
				return 0, err
			}
			node.Cols = append(node.Cols, -(offset + 1))
		case deps.IsSolvedBy(rhsSolves):
			offset, err := pushProjection(expr, node.Right, semTable, inner && node.Opcode != engine.LeftJoin)
			if err != nil {
				return 0, err
			}
			node.Cols = append(node.Cols, offset+1)
		default:
			return 0, vterrors.Errorf(vtrpcpb.Code_INTERNAL, "unknown dependencies for %s", sqlparser.String(expr))
		}
		return len(node.Cols) - 1, nil
	default:
		return 0, vterrors.Errorf(vtrpcpb.Code_UNIMPLEMENTED, "%T not yet supported", node)
	}
}

func removeQualifierFromColName(expr *sqlparser.AliasedExpr) *sqlparser.AliasedExpr {
	if _, ok := expr.Expr.(*sqlparser.ColName); ok {
		expr = sqlparser.CloneRefOfAliasedExpr(expr)
		col := expr.Expr.(*sqlparser.ColName)
		col.Qualifier.Qualifier = sqlparser.NewTableIdent("")
	}
	return expr
}

func checkIfAlreadyExists(expr *sqlparser.AliasedExpr, sel *sqlparser.Select) int {
	for i, selectExpr := range sel.SelectExprs {
		if selectExpr, ok := selectExpr.(*sqlparser.AliasedExpr); ok {
			if sqlparser.EqualsExpr(selectExpr.Expr, expr.Expr) {
				return i
			}
		}
	}
	return -1
}

func planAggregations(qp *queryProjection, plan logicalPlan, semTable *semantics.SemTable) (logicalPlan, error) {
	eaggr := &engine.OrderedAggregate{}
	oa := &orderedAggregate{
		resultsBuilder: newResultsBuilder(plan, eaggr),
		eaggr:          eaggr,
	}
	for _, e := range qp.aggrExprs {
		offset, err := pushProjection(e, plan, semTable, true)
		if err != nil {
			return nil, err
		}
		fExpr := e.Expr.(*sqlparser.FuncExpr)
		opcode := engine.SupportedAggregates[fExpr.Name.Lowered()]
		oa.eaggr.Aggregates = append(oa.eaggr.Aggregates, engine.AggregateParams{
			Opcode: opcode,
			Col:    offset,
		})
	}
	return oa, nil
}

func planOrderBy(qp *queryProjection, orderExprs sqlparser.OrderBy, plan logicalPlan, semTable *semantics.SemTable) (logicalPlan, error) {
	switch plan := plan.(type) {
	case *route:
		return planOrderByForRoute(qp, orderExprs, plan, semTable)
	case *joinGen4:
		return planOrderByForJoin(qp, orderExprs, plan, semTable)
	default:
		return nil, semantics.Gen4NotSupportedF("ordering on complex query")
	}
}

func planOrderByForRoute(qp *queryProjection, orderExprs sqlparser.OrderBy, plan *route, semTable *semantics.SemTable) (logicalPlan, error) {
	origColCount := plan.Select.GetColumnCount()
	for _, order := range orderExprs {
		offset, expr, err := getOrProjectExpr(qp, plan, semTable, order)
		if err != nil {
			return nil, err
		}
		colName, ok := expr.(*sqlparser.ColName)
		if !ok {
			return nil, semantics.Gen4NotSupportedF("order by non-column expression")
		}

		table := semTable.Dependencies(colName)
		tbl, err := semTable.TableInfoFor(table)
		if err != nil {
			return nil, err
		}
		weightStringNeeded := needsWeightString(tbl, colName)

		weightStringOffset := -1
		if weightStringNeeded {
			weightStringOffset, err = pushExpression(plan, expr, semTable)
			if err != nil {
				return nil, err
			}
		}

		plan.eroute.OrderBy = append(plan.eroute.OrderBy, engine.OrderbyParams{
			Col:             offset,
			WeightStringCol: weightStringOffset,
			Desc:            order.Direction == sqlparser.DescOrder,
		})
		plan.Select.AddOrder(order)
	}
	if origColCount != plan.Select.GetColumnCount() {
		plan.eroute.TruncateColumnCount = origColCount
	}

	return plan, nil
}

func pushExpression(plan *route, expr sqlparser.Expr, semTable *semantics.SemTable) (int, error) {
	aliasedExpr := &sqlparser.AliasedExpr{
		Expr: &sqlparser.FuncExpr{
			Name: sqlparser.NewColIdent("weight_string"),
			Exprs: []sqlparser.SelectExpr{
				&sqlparser.AliasedExpr{
					Expr: expr,
				},
			},
		},
	}
	return pushProjection(aliasedExpr, plan, semTable, true)
}

func needsWeightString(tbl semantics.TableInfo, colName *sqlparser.ColName) bool {
	for _, c := range tbl.GetColumns() {
		if colName.Name.String() == c.Name {
			return !sqltypes.IsNumber(c.Type)
		}
	}
	return true // we didn't find the column. better to add just to be safe1
}

// getOrProjectExpr either gets the offset to the expression if it is already projected,
// or pushes the projection if needed
func getOrProjectExpr(qp *queryProjection, plan *route, semTable *semantics.SemTable, order *sqlparser.Order) (offset int, expr sqlparser.Expr, err error) {
	offset, exists := qp.orderExprColMap[order]
	if exists {
		// we are ordering by an expression that is already part of the output
		return offset, qp.selectExprs[offset].Expr, nil
	}

	aliasedExpr := &sqlparser.AliasedExpr{Expr: order.Expr}
	offset, err = pushProjection(aliasedExpr, plan, semTable, true)
	if err != nil {
		return 0, nil, err
	}

	return offset, order.Expr, nil
}

func planOrderByForJoin(qp *queryProjection, orderExprs sqlparser.OrderBy, plan *joinGen4, semTable *semantics.SemTable) (logicalPlan, error) {
	isAllLeft := true
	var err error
	for _, expr := range orderExprs {
		exprDependencies := semTable.Dependencies(expr.Expr)
		if !exprDependencies.IsSolvedBy(plan.Left.ContainsTables()) {
			isAllLeft = false
			break
		}
	}
	if isAllLeft {
		plan.Left, err = planOrderBy(qp, orderExprs, plan.Left, semTable)
		if err != nil {
			return nil, err
		}
		return plan, nil
	}

	return &memorySortGen4{
		orderBy:             nil,
		input:               plan,
		truncateColumnCount: 0,
	}, nil

}
