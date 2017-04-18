package serialize

import "testing"

func TestKairosDBSerialize_WriteInt(t *testing.T) {
	kdbs := KairosDBSerialize{}
	kdbs.Start()
	kdbs.WriteInt(createDummyIntPoints())
	kdbs.WriteInt(createDummyIntPoints())
	kdbs.End()
	log.Info(string(kdbs.Data()))
	// 1489891475000
	//    2147483647
}
