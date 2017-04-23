package playground

import (
	"testing"
	"fmt"
)

// Golang support swap element like python, amazing /w\
func TestSlice_SwapElement(t *testing.T) {
	s := []string{"a", "b"}
	s[0], s[1] = s[1], s[0]
	fmt.Println(s)
}