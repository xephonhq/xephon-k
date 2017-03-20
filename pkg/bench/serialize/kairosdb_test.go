package serialize

import "testing"

func TestKairosDBSerialize_WriteInt(t *testing.T) {
	kdbs := KairosDBSerialize{}
	log.Info(string(kdbs.WriteInt(createDummyIntPoints())))
	// 1489891475000
	//    2147483647
}
