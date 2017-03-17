package memory

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
		p1 := common.IntPoint{TimeNano: 1359788400002, V: 1}
		ps1 := []common.IntPoint{p1}
		s1 := common.IntSeries{Name: "cpi", Tags: tags, Points: ps1}
		multipleSeries = append(multipleSeries, s1)
	}
	return multipleSeries
}

func TestStore_WriteIntSeries(t *testing.T) {
	store := NewMemStore()
	store.WriteIntSeries(createDummySeries())
}

func TestStore_QueryIntSeries(t *testing.T) {
	asst := assert.New(t)
	store := NewMemStore()
	store.WriteIntSeries(createDummySeries())
	tags := make(map[string]string)
	tags["os"] = "ubuntu"
	tags["machine"] = "machine-1"
	qExact := common.Query{Tags: tags, Name: "cpi", MatchPolicy: "exact"}
	returnedSeries, err := store.QueryIntSeries(qExact)
	asst.Nil(err)
	asst.Equal(1, len(returnedSeries))
	// FIXME: the store length is zero
	asst.Equal(1, len(returnedSeries[0].Points))
	log.Info(returnedSeries[0].Points[0].TimeNano)
}
