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

func DBString(db int) string {
	switch db {
	case DBXephonK:
		return "xephonk"
	case DBInfluxDB:
		return "influxdb"
	case DBKairosDB:
		return "kairosdb"
	}
	return "unknown"
}
