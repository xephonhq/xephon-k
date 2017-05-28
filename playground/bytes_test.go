package playground

import "testing"

func TestBytes_NegativeIndex(t *testing.T) {
	b := make([]byte, 10)
	b[9] = 1
	// no, you can't use negative index
	// b[-1]
}
