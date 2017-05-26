package memory

import (
	"fmt"
	"github.com/xephonhq/xephon-k/pkg/common"
	"testing"
)

func createDummySeries() []common.IntSeries {
	// create a bunch of series
	machineNumber := 10
	multipleSeries := make([]common.IntSeries, 0, 10)
	for i := 0; i < machineNumber; i++ {
		tags := make(map[string]string)
		tags["os"] = "ubuntu"
		tags["machine"] = fmt.Sprintf("machine-%d", i)
		p1 := common.IntPoint{T: 1359788400002, V: 1}
		ps1 := []common.IntPoint{p1}
		s1 := common.IntSeries{SeriesMeta: common.SeriesMeta{Name: "cpi", Tags: tags}, Points: ps1}
		multipleSeries = append(multipleSeries, s1)
	}
	return multipleSeries
}

// TODO: validate the write
func TestStore_WriteIntSeries(t *testing.T) {
	store := NewMemStore()
	store.WriteIntSeries(createDummySeries())
}

// TODO: test query in batch
