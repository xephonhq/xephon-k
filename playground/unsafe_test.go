package playground

import (
	"testing"
	"unsafe"
	"bytes"
	"encoding/binary"
)

func TestUnsafe_Ptr(t *testing.T) {
	// covert int to a byte slice
	i := int64(1)
	b := *(*[4]byte)(unsafe.Pointer(&i)) // unsafe_test.go:12: [1 0 0 0]
	t.Log(b)
}

// go test -benchmem -bench=. unsafe_test.go
//BenchmarkUnsafe_EncodeInt64-8             200000             12061 ns/op            9440 B/op          7 allocs/op
//BenchmarkBinary_EncodeInt64-8             100000             15835 ns/op           18912 B/op          8 allocs/op

func BenchmarkUnsafe_EncodeInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := uint64(0)
		b := bytes.NewBuffer(nil)
		for ; j < 1000; j++ {
			b.Write((*(*[4]byte)(unsafe.Pointer(&i)))[:])
		}
	}
}

func BenchmarkBinary_EncodeInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := uint64(0)
		b := bytes.NewBuffer(nil)
		b8 := make([]byte, 8)
		for ; j < 1000; j++ {
			binary.BigEndian.PutUint64(b8, j)
			b.Write(b8)
		}
	}
}
