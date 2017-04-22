package playground

import (
	"testing"
	"fmt"
)

func TestNestedMap(t *testing.T) {
	var m map[string]map[string]bool
	m = make(map[string]map[string]bool, 10)
	m["app"] = make(map[string]bool, 10)
	fmt.Println(len(m["app"]))
	// NOTE: you can't use cap on map http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#map_cap
	// fmt.Println(cap(m["app"]))
}
