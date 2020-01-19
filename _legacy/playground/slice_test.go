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

func TestSlice_Sub(t *testing.T) {
	s := make([]byte, 10)
	s2 := s[1:10]
	t.Log(s2)
}


// it's quite strange, I used to think range is slower, but the result is opposite
func BenchmarkSlice_IteratePrimitive(b *testing.B) {
	var t int
	arr := make([]int, 1000)
	// 5000000	       373 ns/op
	b.Run("range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, j := range arr {
				t = j
			}
		}
	})
	// 3000000	       467 ns/op
	b.Run("index", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(arr); j++ {
				t = arr[j]
			}
		}
	})
	// 3000000	       557 ns/op
	b.Run("index without len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < 1000; j++ {
				t = arr[j]
			}
		}
	})
	_ = t
}

// for struct it seems using index is faster due to copy
func BenchmarkSlice_IterateStruct(b *testing.B) {
	type ts struct {
		a int
		b string
	}
	var t ts
	arr := make([]ts, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = ts{a: 1, b: "I am a short string"}
	}
	// 1000000	      1421 ns/op
	b.Run("range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, j := range arr {
				t = j
			}
		}
	})
	// 1000000	      1132 ns/op
	b.Run("index", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(arr); j++ {
				t = arr[j]
			}
		}
	})
	// 1000000	      1177 ns/op
	b.Run("index without len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < 1000; j++ {
				t = arr[j]
			}
		}
	})
	_ = t
}
