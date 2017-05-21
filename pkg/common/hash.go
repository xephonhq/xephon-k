package common

import "sort"

const (
	prime64  = 1099511628211
	offset64 = 14695981039346656037
)

// check interface
var _ Hashable = (*IntSeries)(nil)
var _ Hashable = (*DoubleSeries)(nil)
var _ Hashable = (*Query)(nil)

type Hashable interface {
	GetName() string
	GetTags() map[string]string
}

// Hash returns one result for series/query have same name and tags
func Hash(src Hashable) SeriesID {
	h := NewInlineFNV64a()
	h.Write([]byte(src.GetName()))
	tags := src.GetTags()
	keys := make([]string, len(tags))
	// TODO: more efficient way for hashing, every time we hash, we sort it. And using []byte might be better than string
	// NOTE: we need to sort the key to get deterministic result
	i := 0
	for k := range tags {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		h.Write([]byte(k))
		h.Write([]byte(tags[k]))
	}
	return SeriesID(h.Sum64())
}

// InlineFNV64a is copied from influxdb/models, which is a alloc-free version of pkg/hash/fnv
// FNV is non crypto hash, so it's faster than MD5 see playground/hash_test.go for benchmark
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
