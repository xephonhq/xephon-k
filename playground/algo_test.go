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
