package playground

import (
	"fmt"
	"testing"
)

type index struct {
	m map[string]string
}

func TestNestedMap(t *testing.T) {
	var m map[string]map[string]bool
	m = make(map[string]map[string]bool, 10)
	m["app"] = make(map[string]bool, 10)
	fmt.Println(len(m["app"]))
	// NOTE: you can't use cap on map http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#map_cap
	// fmt.Println(cap(m["app"]))
}

func TestMap_Init(t *testing.T) {
	i := index{}
	// i.m["ha"] = "hahah" // assignment to entry in nil map [recovered]
	i.m = make(map[string]string)
	i.m["ha"] = "hahah"
}

type foo struct {
	bar string
}

func TestMap_Range(t *testing.T) {
	m := make(map[string]foo)
	m["a"] = foo{bar: "a"}
	// this can't modify the value
	for k := range m {
		t := m[k]
		t.bar = "t"
	}
	t.Log(m)
	// you have to assign the modified value
	for k := range m {
		m[k] = foo{bar: "t"}
	}
	t.Log(m)
}
