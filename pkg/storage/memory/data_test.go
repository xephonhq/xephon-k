package memory

import (
	"testing"

	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestIntSeriesStore_WriteSeries(t *testing.T) {
	asst := assert.New(t)
	store := NewIntSeriesStore()
	asst.Equal(0, store.length)
	tags := make(map[string]string)
	tags["os"] = "ubuntu"
	p1 := common.IntPoint{TimeNano: 1359788400002, V: 1}
	p2 := common.IntPoint{TimeNano: 1359788400003, V: 2}
	ps1 := []common.IntPoint{p1, p2}
	s1 := common.IntSeries{Name: "cpi", Tags: tags, Points: ps1}
	store.WriteSeries(s1)
	asst.Equal(2, store.length)
	p3 := common.IntPoint{TimeNano: 1359788400001, V: 1}
	p4 := common.IntPoint{TimeNano: 1359788400004, V: 2}
	ps2 := []common.IntPoint{p3, p4}
	s2 := common.IntSeries{Name: "cpi", Tags: tags, Points: ps2}
	store.WriteSeries(s2)
	// FIXME: this test broke
	asst.Equal(4, store.length)
}
