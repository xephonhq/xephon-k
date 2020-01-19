package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/common"
)

func TestIntSeriesStore_WriteSeries(t *testing.T) {
	asst := assert.New(t)
	store := NewIntSeriesStore(common.IntSeries{})
	asst.Equal(0, store.length)
	tags := map[string]string{"os": "ubuntu"}
	p1 := common.IntPoint{T: 1359788400002, V: 1}
	p2 := common.IntPoint{T: 1359788400003, V: 2}
	s1 := common.IntSeries{SeriesMeta: common.SeriesMeta{Name: "cpi", Tags: tags}, Points: []common.IntPoint{p1, p2}}
	store.WriteSeries(s1)
	asst.Equal(2, store.length)
	p3 := common.IntPoint{T: 1359788400001, V: 3}
	p4 := common.IntPoint{T: 1359788400004, V: 4}
	s2 := common.IntSeries{SeriesMeta: common.SeriesMeta{Name: "cpi", Tags: tags}, Points: []common.IntPoint{p3, p4}}
	store.WriteSeries(s2)
	asst.Equal(4, store.length)
	p5 := common.IntPoint{T: 1359788400005, V: 5}
	p6 := common.IntPoint{T: 1359788400006, V: 6}
	s3 := common.IntSeries{SeriesMeta: common.SeriesMeta{Name: "cpi", Tags: tags}, Points: []common.IntPoint{p5, p6}}
	store.WriteSeries(s3)
	asst.Equal(6, store.length)
}
