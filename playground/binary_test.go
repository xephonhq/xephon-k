package playground

import (
	"testing"
	"encoding/binary"
)

func TestBinary_Int64(t *testing.T) {
	v := int64(-1)
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(v))
	vd := binary.BigEndian.Uint64(b[:])
	t.Log(vd)
	t.Log(int64(vd))
}

func TestBinary_Int64Uint64(t *testing.T) {
	a := int64(1)
	b := uint64(a)
	t.Log(a, b)
}
