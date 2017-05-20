package playground

import (
	"testing"
	"hash/fnv"
	"crypto/md5"
	"encoding/binary"
)

// test the speed of non crypto and crypto hash function

// InlineFNV64a is copied from influxdb/models
const (
	prime64  = 1099511628211
	offset64 = 14695981039346656037
)

type InlineFNV64a uint64

// NewInlineFNV64a returns a new instance of InlineFNV64a.
func NewInlineFNV64a() InlineFNV64a {
	return offset64
}

// Write adds data to the running hash.
func (s *InlineFNV64a) Write(data []byte) (int, error) {
	hash := uint64(*s)
	for _, c := range data {
		hash ^= uint64(c)
		hash *= prime64
	}
	*s = InlineFNV64a(hash)
	return len(data), nil
}

// Sum64 returns the uint64 of the current resulting hash.
func (s *InlineFNV64a) Sum64() uint64 {
	return uint64(*s)
}

func TestHash_FNV(t *testing.T) {
	h := fnv.New64a()
	h.Write([]byte("hahah"))
	t.Log(h.Sum(nil)) // 8 byte, in big endian
	t.Log(h.Sum64())
	t.Log(binary.BigEndian.Uint64(h.Sum(nil)))
}

func TestHash_FNVInline(t *testing.T) {
	h := NewInlineFNV64a()
	h.Write([]byte("hahah"))
	// FIXME: cannot use h (type InlineFNV64a) as type io.Writer in argument to io.WriteString:
	// InlineFNV64a does not implement io.Writer (Write method has pointer receiver)
	//io.WriteString(h, "hahah")
	t.Log(h.Sum64()) // 8102845894955527583
}

func TestHash_MD5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("hahah"))
	t.Log(h.Sum(nil)) // 16 byte
}

func fnvHash(s string) {
	h := fnv.New64a()
	h.Write([]byte(s))
	h.Sum(nil)
}

func fnvInline(s string) {
	h := NewInlineFNV64a()
	h.Write([]byte(s))
	h.Sum64()
}

func md5Hash(s string) {
	h := md5.New()
	h.Write([]byte(s))
	h.Sum(nil)
}

// 30000000	        40.2 ns/op ... forgot to call Sum
// 20000000	        63.2 ns/op
func BenchmarkHash_FNV(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		fnvHash("hahaha")
	}
}

// 100000000	        12.0 ns/op (much faster, guess because we don't have any object?)
func BenchmarkHash_FNVInline(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		fnvInline("hahaha")
	}
}

// 20000000	        61.7 ns/op ... forgot to call Sum
// 10000000	       223 ns/op
func BenchmarkHash_MD5(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		md5Hash("hahaha")
	}
}
