package cassandra

import (
	"testing"
	"fmt"
	"github.com/xephonhq/xephon-k/pkg/common"
)

// TODO: only run this test when cassandra is presented

// FIXME: this code is copied from memory test
func createDummySeries() []common.IntSeries {
	// create a bunch of series
	machineNumber := 10
	multipleSeries := make([]common.IntSeries, 0, 10)
	for i := 0; i < machineNumber; i++ {
		tags := make(map[string]string)
		tags["os"] = "ubuntu"
		tags["machine"] = fmt.Sprintf("machine-%d", i)
		p1 := common.IntPoint{TimeNano: 1359788400002, V: 1}
		p2 := common.IntPoint{TimeNano: 1359788500002, V: 2}
		ps1 := []common.IntPoint{p1, p2}
		s1 := common.IntSeries{Name: "cpi", Tags: tags, Points: ps1}
		multipleSeries = append(multipleSeries, s1)
	}
	return multipleSeries
}

func TestStore_WriteIntSeries(t *testing.T) {
	store := GetDefaultCassandraStore()
	store.WriteIntSeries(createDummySeries())
}
