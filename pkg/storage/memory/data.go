package memory

import (
	"github.com/xephonhq/xephon-k/pkg/common"
)

var initPointsLength = 100

// Data is a map using SeriesID as key
type Data struct {
	intSeries    map[common.SeriesID]*IntSeriesStore
	doubleSeries map[common.SeriesID]*DoubleSeriesStore
}

func NewData(capacity int) *Data {
	return &Data{
		intSeries:    make(map[common.SeriesID]*IntSeriesStore, capacity),
		doubleSeries: make(map[common.SeriesID]*DoubleSeriesStore, capacity),
	}
}

// WriteIntSeries create the entry if it does not exist, otherwise merge with existing
// TODO: shouldn't we use pointer here?
func (data *Data) WriteIntSeries(id common.SeriesID, series common.IntSeries) error {
	seriesStore, ok := data.intSeries[id]
	if ok {
		log.Debugf("mem:data merge with entry %s %v in map", series.Name, series.Tags)
		err := seriesStore.WriteSeries(series)
		if err != nil {
			return err
		}
	} else {
		log.Debugf("mem:data create new entry %s %v in map", series.Name, series.Tags)
		data.intSeries[id] = NewIntSeriesStore()
		// FIXED: http://stackoverflow.com/questions/32751537/why-do-i-get-a-cannot-assign-error-when-setting-value-to-a-struct-as-a-value-i
		// store.data[id].series = oneSeries
		seriesStore = data.intSeries[id]
		seriesStore.series = series
		// FIXED, we should set the length as well
		seriesStore.length = len(series.Points)
	}
	return nil
}
