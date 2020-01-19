package common

// SeriesID is hash result of metric name and (sorted) tags
// TODO:
// - locality sensitive hashing https://github.com/xephonhq/xephon-k/issues/25
// - distributed hashing
// FIXED:
// - use integer instead of string https://github.com/xephonhq/xephon-k/issues/36
type SeriesID uint64
