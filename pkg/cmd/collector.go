package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/collector"
	"github.com/xephonhq/xephon-k/pkg/collector/system"
	"os"
	"os/signal"
	"time"
	"fmt"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/bench/serialize"
	"net/http"
	"io"
	"io/ioutil"
)

var CollectorCmd = &cobra.Command{
	Use:   "xkc",
	Short: "Xephon K Collector",
	Long:  "xkc is the metrics collector for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		// let's just assume we only report to Xephon-K

		config := collector.NewConfig()
		currentBatchSize := 0
		batchSize := config.BatchSize
		tickChan := time.NewTicker(config.Interval).C

		hostInfo := system.NewHostInfo()
		cpuCollector := system.StatCollector{}
		memCollector := system.MeminfoCollector{}

		// prepare the series
		cpuMetrics := []string{
			"user", "nice", "system", "idle", "iowait", "irq", "softirq", "steal", "guest", "guestnice",
		}
		metricNames := []string{
			"mem.total", "mem.free",
			// "mem.buffers", "mem.cached", "mem.active", "mem.inactive", "mem.dirty",
			// "mem.writeback", "mem.mapped", "mem.shmem",
			// "mem.slab", "mem.sreclaimable", "mem.sunreclaim",
			// "mem.kernelstack", "mem.pagetables", "mem.writebacktmp",
			// "mem.swapcached", "mem.swaptotal", "mem.swapfree",
		}
		// FIXME: re enable CPU metrics after figured out a better way
		for _, m := range cpuMetrics {
			metricNames = append(metricNames, fmt.Sprintf("cpu.total.%s", m))
			for i := 0; i < hostInfo.NumCores; i++ {
				metricNames = append(metricNames, fmt.Sprintf("cpu.%d.%s", i, m))
			}
		}
		// log.Info(metricNames)

		seriesMap := make(map[string]*common.IntSeries, len(metricNames))
		for _, m := range metricNames {
			seriesMap[m] = common.NewIntSeries(m)
		}

		// client and serializer
		client := http.Client{}
		serializer := serialize.XephonKSerialize{}
		baseReq, err := http.NewRequest("POST", "http://localhost:23333/write", nil)
		if err != nil {
			log.Fatalf("can't create base request %v", err)
			return
		}

		// catch CTRL + C
		// http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)

		for {
			select {
			case <-sigChan:
				log.Info("you pressed ctrl + c")
				log.Info("this is dummy clean up")
				os.Exit(0)
			case <-tickChan:
				log.Info("tick")
				// FIXME: IntPoint use UnixNano, but it is actually using millisecond
				currentTime := time.Now().Unix() * 1000
				//log.Info(currentTime)
				if currentBatchSize >= batchSize {
					// flush
					// send the data to xephon
					log.Info("I should flush now!")
					serializer.Start()
					for _, s := range seriesMap {
						s.Tags["host"] = hostInfo.Hostname
						// TODO: why WriteInt is pass by value instead of passing pointer
						serializer.WriteInt(*s)
					}
					serializer.End()
					req := new(http.Request)
					*req = *baseReq
					req.Body = serializer.ReadCloser()
					res, err := client.Do(req)
					if err != nil {
						log.Warn(err)
					} else {
						io.Copy(ioutil.Discard, res.Body)
						res.Body.Close()
						log.Info("flushed")
					}
					currentBatchSize = 0
				} else {
					currentBatchSize++
				}
				// TODO: actually they should be updated concurrently
				cpuCollector.Update()
				memCollector.Update()
				// FIXME: this should not be human work ....
				//for i := 0; i < hostInfo.NumCores; i++ {
				// FIXME: .... it's too hard to add points to metrics, need a lot of copy and paste code
				//	series := seriesMap[fmt.Sprintf("cpu.%s")]
				//}
				s := seriesMap["mem.total"]
				s.Points = append(s.Points, common.IntPoint{TimeNano: currentTime, V: int(memCollector.MemTotal)})
				s = seriesMap["mem.free"]
				s.Points = append(s.Points, common.IntPoint{TimeNano: currentTime, V: int(memCollector.MemFree)})
			}
		}
	},
}

func ExecuteCollector() {
	if CollectorCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	// global flags
	CollectorCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
}
