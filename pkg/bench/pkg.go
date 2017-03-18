package bench

import "time"

const (
	DBXephonK = iota
	DBInfluxDB
	DBPrometheus
	DBKairosDB
)

type RequestMetric struct {
	Err   error
	Code  int
	Start time.Time
	End   time.Time
}
