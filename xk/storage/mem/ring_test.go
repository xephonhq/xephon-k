package mem

import (
	"testing"
	"runtime"
	"time"
	"github.com/dyweb/gommon/util/testutil"
	"sync"
)

func TestNewRing(t *testing.T) {
	r := NewRing(runtime.NumCPU())
	var i uint64
	var j int64
	var batchSize int64 = 10
	for i = 0; i < 100; i++ {
		now := time.Now().UnixNano()
		times := make([]int64, batchSize)
		values := make([]float64, batchSize)
		for j = 0; j < batchSize; j++ {
			times[j] = now + j
			values[j] = float64(j)
		}
		r.getPartition(i).WritePoints(i, times[:], values[:])
	}
	if testutil.Dump().B() {
		for i := 0; i < runtime.NumCPU(); i++ {
			t.Logf("partition %d size %d", i, len(r.partitions[i].store))
			for j, s := range r.partitions[i].store {
				t.Logf("hash %d size %d", j, s.size)
			}
		}
	}
}

// NOTE: race
// FIXME: fatal error: all goroutines are asleep - deadlock!
func TestPartition_WritePoints(t *testing.T) {
	var wg sync.WaitGroup
	var batchSize int64 = 10
	concurrency := runtime.NumCPU()
	r := NewRing(runtime.NumCPU())
	wg.Add(concurrency)
	go func() {
		var i uint64
		var j int64

		for i = 0; i < 100; i++ {
			now := time.Now().UnixNano()
			times := make([]int64, batchSize)
			values := make([]float64, batchSize)
			for j = 0; j < batchSize; j++ {
				times[j] = now + j
				values[j] = float64(j)
			}
			r.getPartition(i).WritePoints(i, times[:], values[:])
		}
		wg.Done()
	}()
	wg.Wait()
}
