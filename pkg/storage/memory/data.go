package memory

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
)

var initPointsLength = 100

// Data is a map using SeriesID as key
type Data struct {
	series map[common.SeriesID]SeriesStore
}

func NewData(capacity int) *Data {
	return &Data{
		series: make(map[common.SeriesID]SeriesStore, capacity),
	}
}

func (data *Data) WriteIntSeries(id common.SeriesID, series common.IntSeries) error {
	store, ok := data.series[id]
	if !ok {
		// create the new store and put it in the map
		log.Debugf("mem:data create new entry %s %v in map", series.Name, series.Tags)
		data.series[id] = NewIntSeriesStore(series)
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
	case common.TypeIntSeries:
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
