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

func TestSlice_Equal(t *testing.T) {
	s := []int{1, 2}
	// can't compare two slice
	//s2 := s
	//fmt.Println(s == s2)
	s1 := &s
	s2 := &s
	// we can compare pointer to slice though
	fmt.Println(s1 == s2)
}