package memory

import (
	"github.com/xephonhq/xephon-k/pkg/common"
)

var initPointsLength = 100

// Data is a map using SeriesID as key
// TODO: should be able to allow double etc later, maybe wrap it using a struct
type Data map[common.SeriesID]*IntSeriesStore

// WriteIntSeries create the entry if it does not exist, otherwise merge with existing
// TODO: return error
// TODO: shouldn't we use pointer here?
func (data Data) WriteIntSeries(id common.SeriesID, series common.IntSeries) {
	seriesStore, ok := data[id]
	if ok {
		log.Debugf("mem:data merge with existing series %s", series.Name)
		seriesStore.WriteSeries(series)
	} else {
		// TODO: log the tags
		log.Debugf("mem:data create new entry %s %v in map", series.Name, series.Tags)
		data[id] = NewIntSeriesStore()
		// FIXED: http://stackoverflow.com/questions/32751537/why-do-i-get-a-cannot-assign-error-when-setting-value-to-a-struct-as-a-value-i
		// store.data[id].series = oneSeries
		seriesStore = data[id]
		// TODO: I think this does not work
		seriesStore.series = series
		// FIXED, we should set the length as well
		seriesStore.length = len(series.Points)
	}
}
