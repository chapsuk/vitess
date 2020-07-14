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

package discovery

import (
	"testing"
	"time"

	"vitess.io/vitess/go/vt/log"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	querypb "vitess.io/vitess/go/vt/proto/query"
	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
	"vitess.io/vitess/go/vt/topo"
	"vitess.io/vitess/go/vt/topo/memorytopo"
)

func TestPickSimple(t *testing.T) {
	te := newPickerTestEnv(t, []string{"cell"})
	want := addTablet(te, 100, topodatapb.TabletType_REPLICA, "cell", true, true)
	defer deleteTablet(te, want)

	tp, err := NewTabletPicker(te.topoServ, te.cells, te.keyspace, te.shard, "replica")
	require.NoError(t, err)

	tablet, err := tp.PickForStreaming(context.Background())
	require.NoError(t, err)
	if !proto.Equal(want, tablet) {
		t.Errorf("Pick: %v, want %v", tablet, want)
	}
}

func TestPickFromOtherCell(t *testing.T) {
	te := newPickerTestEnv(t, []string{"cell", "otherCell"})
	want := addTablet(te, 100, topodatapb.TabletType_REPLICA, "otherCell", true, true)
	defer deleteTablet(te, want)

	tp, err := NewTabletPicker(te.topoServ, te.cells, te.keyspace, te.shard, "replica")
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	tablet, err := tp.PickForStreaming(ctx)
	require.NoError(t, err)
	if !proto.Equal(want, tablet) {
		t.Errorf("Pick: %v, want %v", tablet, want)
	}
}

func TestPickFromTwoHealthy(t *testing.T) {
	te := newPickerTestEnv(t, []string{"cell"})
	want1 := addTablet(te, 100, topodatapb.TabletType_REPLICA, "cell", true, true)
	defer deleteTablet(te, want1)
	want2 := addTablet(te, 101, topodatapb.TabletType_RDONLY, "cell", true, true)
	defer deleteTablet(te, want2)

	tp, err := NewTabletPicker(te.topoServ, te.cells, te.keyspace, te.shard, "replica,rdonly")
	require.NoError(t, err)

	// In 20 attempts, both tablet types must be picked at least once.
	var picked1, picked2 bool
	for i := 0; i < 20; i++ {
		tablet, err := tp.PickForStreaming(context.Background())
		require.NoError(t, err)
		if proto.Equal(tablet, want1) {
			picked1 = true
		}
		if proto.Equal(tablet, want2) {
			picked2 = true
		}
	}
	assert.True(t, picked1)
	assert.True(t, picked2)
}

func TestPickRespectsTabletType(t *testing.T) {
	te := newPickerTestEnv(t, []string{"cell"})
	want := addTablet(te, 100, topodatapb.TabletType_REPLICA, "cell", true, true)
	defer deleteTablet(te, want)
	dont := addTablet(te, 101, topodatapb.TabletType_MASTER, "cell", true, true)
	defer deleteTablet(te, dont)

	tp, err := NewTabletPicker(te.topoServ, te.cells, te.keyspace, te.shard, "replica,rdonly")
	require.NoError(t, err)

	// In 20 attempts, master tablet must be never picked
	for i := 0; i < 20; i++ {
		tablet, err := tp.PickForStreaming(context.Background())
		require.NoError(t, err)
		require.NotNil(t, tablet)
		require.True(t, proto.Equal(tablet, want), "picked wrong tablet type")
	}
}

func TestPickUsingCellAlias(t *testing.T) {
	// test env puts all cells into an alias called "cella"
	te := newPickerTestEnv(t, []string{"cell"})
	want := addTablet(te, 100, topodatapb.TabletType_REPLICA, "cell", true, true)
	defer deleteTablet(te, want)

	tp, err := NewTabletPicker(te.topoServ, []string{"cella"}, te.keyspace, te.shard, "replica")
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	tablet, err := tp.PickForStreaming(ctx)
	require.NoError(t, err)
	if !proto.Equal(want, tablet) {
		t.Errorf("Pick: %v, want %v", tablet, want)
	}
}

func TestPickError(t *testing.T) {
	te := newPickerTestEnv(t, []string{"cell"})
	defer deleteTablet(te, addTablet(te, 100, topodatapb.TabletType_REPLICA, "cell", false, false))

	_, err := NewTabletPicker(te.topoServ, te.cells, te.keyspace, te.shard, "badtype")
	assert.EqualError(t, err, "failed to parse list of tablet types: badtype")

	_, err = NewTabletPicker(te.topoServ, te.cells, te.keyspace, te.shard, "replica,rdonly")
	require.NoError(t, err)
}

type pickerTestEnv struct {
	t        *testing.T
	keyspace string
	shard    string
	cells    []string

	topoServ *topo.Server
}

func newPickerTestEnv(t *testing.T, cells []string) *pickerTestEnv {
	ctx := context.Background()

	te := &pickerTestEnv{
		t:        t,
		keyspace: "ks",
		shard:    "0",
		cells:    cells,
		topoServ: memorytopo.NewServer(cells...),
	}
	// create cell alias
	err := te.topoServ.CreateCellsAlias(ctx, "cella", &topodatapb.CellsAlias{
		Cells: cells,
	})
	require.NoError(t, err)
	err = te.topoServ.CreateKeyspace(ctx, te.keyspace, &topodatapb.Keyspace{})
	require.NoError(t, err)
	err = te.topoServ.CreateShard(ctx, te.keyspace, te.shard)
	require.NoError(t, err)
	return te
}

func addTablet(te *pickerTestEnv, id int, tabletType topodatapb.TabletType, cell string, serving, healthy bool) *topodatapb.Tablet {
	tablet := &topodatapb.Tablet{
		Alias: &topodatapb.TabletAlias{
			Cell: cell,
			Uid:  uint32(id),
		},
		Keyspace: te.keyspace,
		Shard:    te.shard,
		KeyRange: &topodatapb.KeyRange{},
		Type:     tabletType,
		PortMap: map[string]int32{
			"test": int32(id),
		},
	}
	err := te.topoServ.CreateTablet(context.Background(), tablet)
	require.NoError(te.t, err)

	var herr string
	if !healthy {
		herr = "err"
	}
	_ = createFixedHealthConn(tablet, &querypb.StreamHealthResponse{
		Serving: serving,
		Target: &querypb.Target{
			Keyspace:   te.keyspace,
			Shard:      te.shard,
			TabletType: tabletType,
		},
		RealtimeStats: &querypb.RealtimeStats{HealthError: herr},
	})

	return tablet
}

func deleteTablet(te *pickerTestEnv, tablet *topodatapb.Tablet) {

	//log error
	if err := te.topoServ.DeleteTablet(context.Background(), tablet.Alias); err != nil {
		log.Errorf("failed to DeleteTablet with alias : %v", err)
	}

	//This is not automatically removed from shard replication, which results in log spam and log error
	if err := topo.DeleteTabletReplicationData(context.Background(), te.topoServ, tablet); err != nil {
		log.Errorf("failed to automatically remove from shard replication: %v", err)
	}
}
