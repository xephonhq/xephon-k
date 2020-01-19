package playground

import (
	"testing"
	"fmt"
)

func TestSort_Selection(t *testing.T) {
	// selection sort
	// TODO: bechmark with stl's sort
	arr := []int{3, 1, 8, 9, 2}
	var smallestVal int
	var smallestIndex int
	for i := 0; i < len(arr)-1; i++ {
		smallestVal = arr[i]
		smallestIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < smallestVal {
				smallestIndex = j
				smallestVal = arr[j]
			}
		}
		// swap
		if i != smallestIndex {
			arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
		}
	}
	fmt.Println(arr)
}

// https://www.cs.cmu.edu/~adamchik/15-121/lectures/Binary%20Heaps/heaps.html
type heap struct {
	data []int
}

// TODO: percolation down

func (h *heap) insert(x int) {
	h.data = append(h.data, x)
	// percolation up
	pos := len(h.data) - 1
	for ; pos > 1 && x < h.data[pos/2]; pos = pos / 2 {
		h.data[pos] = h.data[pos/2]
	}
	h.data[pos] = x
}

func TestHeap_Insert(t *testing.T) {
	h := heap{data: []int{0, 6, 7, 12, 10, 15, 17}}
	h.insert(5)
	fmt.Println(h.data)
}
