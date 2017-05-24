package common

// SeriesID is hash result of metric name and (sorted) tags
// TODO:
// - locality sensitive hashing https://github.com/xephonhq/xephon-k/issues/25
// - distributed hashing
// - use integer instead of string https://github.com/xephonhq/xephon-k/issues/36
type SeriesID uint64

type SeriesIDs []SeriesID

func (ids SeriesIDs) Len() int {
	return len(ids)
}

func (ids SeriesIDs) Swap(i int, j int) {
	ids[i], ids[j] = ids[j], ids[i]
}

func (ids SeriesIDs) Less(i int, j int) bool {
	return ids[i] < ids[j]
}
