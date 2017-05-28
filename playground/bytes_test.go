package playground

import "testing"

func TestBytes_NegativeIndex(t *testing.T) {
	b := make([]byte, 10)
	b[9] = 1

	// this won't compile
	// b[-1]

	// this will panic: runtime error: index out of range
	//i := -1
	//_ = b[i]
}
