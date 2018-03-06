package mem

import "sync"

// Ring is a simple cache without compression
type Ring struct {
	numPartitions uint64
	// TODO: benchmark performance of using pointer and struct https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/
	partitions []*Partition
}

// TODO: value for n?
func NewRing(n int) *Ring {
	r := &Ring{
		numPartitions: uint64(n),
		partitions:    make([]*Partition, n),
	}
	for i := 0; i < n; i++ {
		r.partitions[i] = &Partition{
			store: make(map[uint64]*Store),
		}
	}
	return r
}

func (r *Ring) getPartition(hash uint64) *Partition {
	return r.partitions[int(hash%r.numPartitions)]
}

type Partition struct {
	mu    sync.RWMutex
	store map[uint64]*Store
}

type Store struct {
	// TODO: metrics tags
	// FIXME: only support float64? what about int, bool etc.?
	mu     sync.RWMutex
	times  []int64
	values []float64
	size   int
}

func NewStore(len int) *Store {
	return &Store{
		times:  make([]int64, 0, len),
		values: make([]float64, 0, len),
	}
}

func (p *Partition) WritePoints(hash uint64, times []int64, values []float64) {
	p.mu.RLock()
	s := p.store[hash]
	if s != nil {
		s.mu.Lock()
		p.mu.RUnlock()
		s.size += len(times)
		s.times = append(s.times, times...)
		s.values = append(s.values, values...)
		s.mu.Unlock()
		return
	}
	p.mu.RUnlock()
	p.mu.Lock()
	s = NewStore(len(times))
	s.size += len(times)
	s.times = append(s.times, times...)
	s.values = append(s.values, values...)
	p.store[hash] = s
	p.mu.Unlock()
	return
}
