package system

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
	//"fmt"
)

var meminfoPath = "/proc/meminfo"

type VirtualMemoryStat struct {
	MemTotal     uint64
	MemFree      uint64
	MemAvailable uint64
	Buffers      uint64
	Cached       uint64
	Active       uint64
	Inactive     uint64
	Dirty        uint64
	Writeback    uint64 // Memory which is actively being written back to the disk
	Mapped       uint64 // Files which have been mapped into memory (with mmap(2)), such as librarie
	Shmem        uint64 // Amount of memory consumed in tmpfs(5) filesystems
	Slab         uint64 // In-kernel data structures cache
	SReclaimable uint64
	SUnreclaim   uint64
	KernelStack  uint64 // Amount of memory allocated to kernel stacks
	PageTables   uint64
	WritebackTmp uint64 // Memory used by FUSE for temporary writeback buffers
	HugePagesize uint64
	DirectMap4k  uint64
	DirectMap2M  uint64
	DirectMap1G  uint64
}

type SwapMemoryStat struct {
	SwapCached uint64
	SwapTotal  uint64
	SwapFree   uint64
}

type MeminfoCollector struct {
	VirtualMemoryStat
	SwapMemoryStat
}

func (collector *MeminfoCollector) Update() error {
	file, err := os.Open(meminfoPath)
	if err != nil {
		return errors.Wrapf(err, "can't open %s", meminfoPath)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		// omit the `:`, i.e. `MemTotal:` -> `MemTotal`
		head := parts[0][:len(parts[0])-1]
		// assume all unit is KB TODO: are there any counter example?
		value, err := strconv.ParseUint(parts[1], 10, 64)
		//fmt.Println(value)
		if err != nil {
			return errors.Wrapf(err, "can't parse %s", head)
		}
		//fmt.Println(head)
		switch head {
		case "MemTotal":
			collector.MemTotal = value
		case "MemFree":
			collector.MemFree = value
		case "MemAvailable":
			collector.MemAvailable = value
		case "Buffers":
			collector.Buffers = value
		case "Cached":
			collector.Cached = value
		case "Active":
			collector.Active = value
		case "Inactive":
			collector.Inactive = value
		case "Dirty":
			collector.Dirty = value
		case "Writeback":
			collector.Writeback = value
		case "Mapped":
			collector.Mapped = value
		case "Shmem":
			collector.Shmem = value
		case "Slab":
			collector.Slab = value
		case "SReclaimable":
			collector.SReclaimable = value
		case "SUnreclaim":
			collector.SUnreclaim = value
		case "KernelStack":
			collector.KernelStack = value
		case "PageTables":
			collector.PageTables = value
		case "WritebackTmp":
			collector.WritebackTmp = value
		case "HugePagesize":
			collector.HugePagesize = value
		case "DirectMap4k":
			collector.DirectMap4k = value
		case "DirectMap2M":
			collector.DirectMap2M = value
		case "DirectMap1G":
			collector.DirectMap1G = value
			// Swap
		case "SwapCached":
			collector.SwapCached = value
		case "SwapTotal":
			collector.SwapTotal = value
		case "SwapFree":
			collector.SwapFree = value
		default:
			// do nothing
			//fmt.Println(head)
		}
	}
	return nil
}
