package serialize

import "testing"

func TestXephonKSerialize_WriteInt(t *testing.T) {
	xks := XephonKSerialize{}
	xks.Start()
	xks.WriteInt(createDummyIntPoints())
	xks.WriteInt(createDummyIntPoints())
	xks.End()
	log.Info(string(xks.Data()))
}
