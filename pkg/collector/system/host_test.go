package system

import (
	"testing"
	"fmt"
)

func TestNewHostInfo(t *testing.T) {
	host := NewHostInfo()
	fmt.Println(host.Hostname)
}
