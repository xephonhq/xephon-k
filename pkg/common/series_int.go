package common

import "time"

// TODO: int series of other precision, maybe we should add millisecond to the default function as well
func NewIntSeries(name string) *IntSeries {
	return &IntSeries{
		Meta: SeriesMeta{
			Name:      name,
			Tags:      map[string]string{},
			Type:      TypeIntSeries,
			Precision: time.Millisecond.Nanoseconds(),
		},
	}
}

func (series *IntSeries) GetName() string {
	return series.Meta.Name
}

func (series *IntSeries) GetTags() map[string]string {
	return series.Meta.Tags
}

func (series *IntSeries) GetSeriesType() int64 {
	return series.Meta.Type
}

func (series *IntSeries) GetSeriesID() SeriesID {
	if series.Meta.Id == 0 {
		series.Meta.Id = uint64(Hash(series))
	}
	return SeriesID(series.Meta.Id)
}
