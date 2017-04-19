package system

import (
	"fmt"
	"testing"
)

func TestNewHostInfo(t *testing.T) {
	host := NewHostInfo()
	fmt.Println(host.Hostname)
}
