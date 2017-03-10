package memory

import (
	"sort"
	"sync"

	"github.com/xephonhq/xephon-k/pkg/common"
)

var initPointsLength = 100

// SeriesID is hash result of metric name and (sorted) tags
type SeriesID string

// Data is a map using SeriesID as key
// TODO: should be able to allow double etc later
type Data map[SeriesID]*IntSeriesStore

// WriteIntSeries create the entry if it does not exist, otherwise merge with existing
// TODO: return error
func (data Data) WriteIntSeries(id SeriesID, series common.IntSeries) {
	seriesStore, ok := data[id]
	if ok {
		log.Info("mem:data merge with existing series")
		seriesStore.WriteSeries(series)
	} else {
		log.Info("mem:data create new entry in map")
		data[id] = NewIntSeriesStore()
		// FIXED: http://stackoverflow.com/questions/32751537/why-do-i-get-a-cannot-assign-error-when-setting-value-to-a-struct-as-a-value-i
		// store.data[id].series = oneSeries
		seriesStore = data[id]
		seriesStore.series = series
	}
}

// IntSeriesStore protects the underlying IntSeries with a RWMutex
type IntSeriesStore struct {
	mu     sync.RWMutex
	series common.IntSeries
	length int
}

// NewIntSeriesStore creates a IntSeriesStore
func NewIntSeriesStore() *IntSeriesStore {
	series := common.IntSeries{}
	series.Points = make([]common.IntPoint, initPointsLength)
	return &IntSeriesStore{series: series, length: 0}
}

// WriteSeries merges the new series with existing one and replace old points with new points if their timestamp matches
// TODO: what happens when no memory is available? maybe this function should return error
func (store *IntSeriesStore) WriteSeries(newSeries common.IntSeries) {
	store.mu.Lock()
	defer store.mu.Unlock()

	// merge the old series with new series
	// TODO: check if they are the same series
	// TODO: efficient merge sort
	// TODO: actually we can remove duplicate when merge by comparing with previous point

	// store.series should already be sorted, so we only sort the newSeries
	sort.Sort(common.ByTime(newSeries.Points))
	i := 0
	j := 0
	k := 0
	// NOTE: we can't use len(store.series.Points) because there might be duplicate
	oldLength := store.length
	newLength := len(newSeries.Points)
	// log.Infof("ol %d nl %d", oldLength, newLength)
	points := make([]common.IntPoint, oldLength+newLength)
	for i < oldLength && j < newLength {
		if store.series.Points[i].TimeNano < newSeries.Points[j].TimeNano {
			points[k] = store.series.Points[i]
			i++
		} else if store.series.Points[i].TimeNano == newSeries.Points[j].TimeNano {
			// if there is duplicate, overwrite the old point with new point
			points[k] = newSeries.Points[j]
			i++
			j++
		} else {
			points[k] = newSeries.Points[j]
			j++
		}
		// log.Infof("value in loop is %v", points[k].TimeNano)
		k++
	}
	// log.Infof("i %d j %d k %d", i, j, k)
	store.length = k

	// copy what is left, should only have one array left
	// https://github.com/golang/go/wiki/SliceTricks
	if i < oldLength {
		// FIXED: should cut ... instead of simply append
		points = append(points[:k], store.series.Points[i:]...)
		store.length = k + oldLength - i
	}
	if j < newLength {
		points = append(points[:k], newSeries.Points[j:]...)
		store.length = k + newLength - j
	}
	// log.Infof("length %d", store.length)

	// TODO: maybe using pool is a good idea since there are a lot of merge when inserting series
	store.series.Points = points

	//n := 0
	//for n < store.length {
	//	log.Infof("time in store %v", store.series.Points[n].TimeNano)
	//	n++
	//}
}
