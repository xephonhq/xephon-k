package memory

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
)

var initPointsLength = 100

// Data is a map using SeriesID as key
type Data struct {
	intSeries map[common.SeriesID]*IntSeriesStore
	//doubleSeries map[common.SeriesID]*DoubleSeriesStore
	series map[common.SeriesID]SeriesStore
}

func NewData(capacity int) *Data {
	return &Data{
		intSeries: make(map[common.SeriesID]*IntSeriesStore, capacity),
		//doubleSeries: make(map[common.SeriesID]*DoubleSeriesStore, capacity),
		series: make(map[common.SeriesID]SeriesStore, capacity),
	}
}

func (data *Data) WriteIntSeries(id common.SeriesID, series common.IntSeries) error {
	store, ok := data.series[id]
	if !ok {
		// create the new store and put it in the map
		log.Debugf("mem:data create new entry %s %v in map", series.Name, series.Tags)
		data.series[id] = NewIntSeriesStore()
		store = data.series[id]
	}
	intStore, ok := store.(*IntSeriesStore)
	if !ok {
		return errors.Errorf("%s %v is %s but tried to write int", store.GetName(), store.GetTags(), store.GetSeriesType())
	}
	err := intStore.WriteSeries(series)
	if err != nil {
		return err
	}
	return nil
}

func (data *Data) ReadSeries(id common.SeriesID, startTime int64, endTime int64) (common.Series, bool, error) {
	// TODO: auto convert the start and end time based on series precision
	store, ok := data.series[id]
	if !ok {
		return &common.MetaSeries{}, false, nil
	}
	switch store.GetSeriesType() {
	case common.TypeStringSeries:
		intStore, ok := store.(*IntSeriesStore)
		if !ok {
			return &common.MetaSeries{}, false, errors.Errorf("%s %v is marked as int but actually %s",
				store.GetName(), store.GetTags(), common.SeriesTypeString(store.GetSeriesType()))
		}
		// TODO: this should also return error
		return intStore.ReadByStartEndTime(startTime, endTime), true, nil
		//case common.TypeDoubleSeries:
		// TODO
	default:
		return &common.MetaSeries{}, false, errors.Errorf("%s %v has unsupported type %s",
			store.GetName(), store.GetTags(), common.SeriesTypeString(store.GetSeriesType()))
	}
}

// WriteIntSeries create the entry if it does not exist, otherwise merge with existing
// TODO: shouldn't we use pointer here?
func (data *Data) WriteIntSeriesOld(id common.SeriesID, series common.IntSeries) error {
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
