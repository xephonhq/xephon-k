package playground

import "testing"

// TODO: test if using range on objects would have extra copy

type Query struct {
	StartTime int
}

func TestRange_Modify(t *testing.T) {
	queries := []Query{{StartTime: 0}, {StartTime: 1}}
	// NOTE: you can't, it makes a copy of the value
	for _, q := range queries {
		q.StartTime = 1
	}
	t.Log(queries)

	for i := 0; i < len(queries); i++ {
		// NOTE: this still makes a copy
		q := queries[i]
		q.StartTime = 1
	}
	t.Log(queries)

	// NOTE: this the only working way
	for i := 0; i < len(queries); i++ {
		q := &queries[i]
		q.StartTime = 1
	}
	t.Log(queries)
}
