package system

import (
	"os"
	"runtime"
)

// TODO: /etc/os-release

type HostInfo struct {
	NumCores int
	Hostname string
}

func NewHostInfo() HostInfo {
	info := HostInfo{}

	info.NumCores = runtime.NumCPU()
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}
	info.Hostname = hostname

	return info
}
