package memory

type Index []IndexRow

type IndexRow struct {
	key      string
	value    string
	seriesID SeriesID
}
