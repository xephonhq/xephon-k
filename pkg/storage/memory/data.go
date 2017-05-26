package memory

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	"reflect"
)

var initPointsLength = 100

// Data is a map using SeriesID as key
type Data struct {
	series map[common.SeriesID]SeriesStore
}

// check interface
var _ SeriesStore = (*IntSeriesStore)(nil)
var _ SeriesStore = (*DoubleSeriesStore)(nil)

type SeriesStore interface {
	common.Hashable
	GetSeriesType() int64
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

func (data *Data) WriteDoubleSeries(id common.SeriesID, series common.DoubleSeries) error {
	store, ok := data.series[id]
	if !ok {
		// create the new store and put it in the map
		log.Debugf("mem:data create new entry %s %v in map", series.Name, series.Tags)
		data.series[id] = NewDoubleSeriesStore(series)
		store = data.series[id]
	}
	doubleStore, ok := store.(*DoubleSeriesStore)
	if !ok {
		return errors.Errorf("%s %v is %s but tried to write double", store.GetName(), store.GetTags(), store.GetSeriesType())
	}
	err := doubleStore.WriteSeries(series)
	if err != nil {
		return err
	}
	return nil
}

func (data *Data) ReadSeries(id common.SeriesID, startTime int64, endTime int64) (common.Series, bool, error) {
	// TODO: auto convert the start and end time based on series precision
	store, ok := data.series[id]
	if !ok {
		// FIXME: use EmptySeries
		return &common.RawSeries{}, false, nil
	}
	switch store.GetSeriesType() {
	case common.TypeIntSeries:
		intStore, ok := store.(*IntSeriesStore)
		if !ok {
			return &common.RawSeries{}, false, errors.Errorf("%s %v is marked as int but actually %s",
				store.GetName(), store.GetTags(), reflect.TypeOf(store))
		}
		// TODO: this should also return error
		return intStore.ReadByStartEndTime(startTime, endTime), true, nil
	case common.TypeDoubleSeries:
		doubleStore, ok := store.(*DoubleSeriesStore)
		if !ok {
			return &common.RawSeries{}, false, errors.Errorf("%s %v is marked as double but actually %s",
				store.GetName(), store.GetTags(), reflect.TypeOf(store))
		}
		// TODO: this should also return error
		return doubleStore.ReadByStartEndTime(startTime, endTime), true, nil
	default:
		return &common.RawSeries{}, false, errors.Errorf("%s %v has unsupported type %s",
			store.GetName(), store.GetTags(), common.SeriesTypeString(store.GetSeriesType()))
	}
}
