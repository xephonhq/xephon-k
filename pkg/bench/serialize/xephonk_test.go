package serialize

import "testing"

func TestXephonKSerialize_WriteInt(t *testing.T) {
	xks := XephonKSerialize{}
	log.Info(string(xks.WriteInt(createDummyIntPoints())))
}
