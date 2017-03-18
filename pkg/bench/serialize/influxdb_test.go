package serialize

import "testing"

func TestInfluxDBSerialize_WriteInt(t *testing.T) {
	influxs := InfluxDBSerialize{}
	log.Info(string(influxs.WriteInt(createDummyIntPoints())))
}
