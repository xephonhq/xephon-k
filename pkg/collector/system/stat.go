package system

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

var statPath = "/proc/stat"

const (
	CPUUserHz = 100 // TODO: this is always 100? or at least for x86
)

/*
/proc/stat contains
- cpu
- TODO: intr
- ctxt context switches
- btime	the time system boot in Unix timestamp, not time since boot
- process number of forks since boot
- procs_running
- procs_blocked blocked waiting for I/O to complete
- TODO: softirq
*/

type CPUStat struct {
	User      float64
	Nice      float64
	System    float64
	Idle      float64
	IOWait    float64
	Irq       float64
	SoftIrq   float64
	Steal     float64
	Guest     float64
	GuestNice float64
}

type ExtraStat struct {
	ContextSwitches uint64
	BootTime        uint64
	Processes       uint64
	ProcessRunning  uint64
	ProcessBlocked  uint64
}

type GlobalStat struct {
	ExtraStat
	CPUs     []CPUStat
	CPUTotal CPUStat
}

func (globalStat *GlobalStat) Update() error {
	file, err := os.Open(statPath)
	defer file.Close()
	if err != nil {
		return errors.Wrap(err, "can't open /proc/stat")
	}
	// NOTE: http://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		head := parts[0]
		if strings.HasPrefix(head, "cpu") {
			// TODO: the first line `cpu` does not equals to the add up of each cpu `cpu0, ... cpu7`, though the diff is little
			stat := CPUStat{}
			values := make([]float64, 10)
			for i, v := range parts[1:11] {
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return errors.Wrap(err, "can't parse CPU time")
				}
				// value is The amount of time, measured in units of USER_HZ http://man7.org/linux/man-pages/man5/proc.5.html
				values[i] = value / CPUUserHz
			}
			stat.User = values[0]
			stat.Nice = values[1]
			stat.System = values[2]
			stat.Idle = values[3]
			stat.IOWait = values[4]
			stat.Irq = values[5]
			stat.SoftIrq = values[6]
			stat.Steal = values[7]
			stat.Guest = values[8]
			stat.GuestNice = values[9]
			if head == "cpu" {
				globalStat.CPUTotal = stat
			} else {
				globalStat.CPUs = append(globalStat.CPUs, stat)
			}
		} else {
			switch head {
			case "ctxt":
				value, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					return errors.Wrap(err, "can't parse context switches")
				}
				globalStat.ContextSwitches = value
			case "processes":
				value, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					return errors.Wrap(err, "can't pare processes")
				}
				globalStat.Processes = value
			case "btime":
				value, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					return errors.Wrap(err, "can't parse boot time")
				}
				// boot time, in seconds since the Epoch
				globalStat.BootTime = value
			case "procs_running":
				value, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					return errors.Wrap(err, "can't parse process running")
				}
				globalStat.ProcessRunning = value
			case "procs_blocked":
				value, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					return errors.Wrap(err, "can't parse process blocked")
				}
				globalStat.ProcessBlocked = value
			default:
				// TODO: log? only `intr` and `softirq` is left
			}
		}
	}
	return nil
}
