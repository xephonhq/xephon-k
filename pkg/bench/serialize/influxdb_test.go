package serialize

import "testing"

func TestInfluxDBSerialize_WriteInt(t *testing.T) {
	influxs := InfluxDBSerialize{}
	influxs.Start()
	influxs.WriteInt(createDummyIntPoints())
	influxs.End()
	log.Info(string(influxs.Data()))
}
