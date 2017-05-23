package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/bench/serialize"
	"github.com/xephonhq/xephon-k/pkg/collector"
	"github.com/xephonhq/xephon-k/pkg/collector/system"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/server"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var CollectorCmd = &cobra.Command{
	Use:   "xkc",
	Short: "Xephon K Collector",
	Long:  "xkc is the metrics collector for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		targetDB := bench.DBXephonK
		// get the database
		if strings.HasPrefix(db, "x") {
			targetDB = bench.DBXephonK
		} else if strings.HasPrefix(db, "i") {
			targetDB = bench.DBInfluxDB
		} else if strings.HasPrefix(db, "k") {
			targetDB = bench.DBKairosDB
		} else {
			log.Fatalf("unsupported target db %s", db)
			return
		}

		// client and serializer
		client := http.Client{}
		var baseReq *http.Request
		var serializer serialize.Serializer
		switch targetDB {
		case bench.DBInfluxDB:
			req, err := http.NewRequest("POST", "http://localhost:8086/write?db=sb", nil)
			if err != nil {
				log.Panic(err)
				return
			}
			baseReq = req
			serializer = &serialize.InfluxDBSerialize{}
		case bench.DBXephonK:
			url := fmt.Sprintf("http://localhost:%d/write", server.DefaultPort)
			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				log.Panic(err)
				return
			}
			baseReq = req
			serializer = &serialize.XephonKSerialize{}
		case bench.DBKairosDB:
			req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/datapoints", nil)
			if err != nil {
				log.Panic(err)
				return
			}
			baseReq = req
			serializer = &serialize.KairosDBSerialize{}
		default:
			log.Panic("unsupported database, no base request avaliable")
			return
		}

		config := collector.NewConfig()
		currentBatchSize := 0
		batchSize := config.BatchSize
		tickChan := time.NewTicker(config.Interval).C

		hostInfo := system.NewHostInfo()
		cpuCollector := system.StatCollector{}
		memCollector := system.MeminfoCollector{}

		// prepare the series
		enableCPU := true

		metricNames := []string{
			"mem.total", "mem.free",
			// "mem.buffers", "mem.cached", "mem.active", "mem.inactive", "mem.dirty",
			// "mem.writeback", "mem.mapped", "mem.shmem",
			// "mem.slab", "mem.sreclaimable", "mem.sunreclaim",
			// "mem.kernelstack", "mem.pagetables", "mem.writebacktmp",
			// "mem.swapcached", "mem.swaptotal", "mem.swapfree",
		}

		cpuMetrics := []string{
			"user", "nice", "system", "idle", "iowait", "irq", "softirq", "steal", "guest", "guestnice",
		}
		// metric prefix for all the CPU cores
		cores := make([]string, hostInfo.NumCores+1)
		for i := 0; i < hostInfo.NumCores; i++ {
			cores[i] = fmt.Sprintf("cpu.%d.", i)
		}
		cores[hostInfo.NumCores] = "cpu.total."
		if enableCPU {
			// add cpu to metric names
			for _, m := range cpuMetrics {
				for _, p := range cores {
					metricNames = append(metricNames, p+m)
				}
			}
		}

		log.Info(metricNames)

		// map of int series
		// TODO: support double series
		seriesMap := make(map[string]*common.IntSeries, len(metricNames))
		// init all the series
		for _, m := range metricNames {
			seriesMap[m] = common.NewIntSeries(m)
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
				// FIXME: IntPoint use UnixNano, but it is actually using millisecond
				currentTime := time.Now().Unix() * 1000
				log.Debugf("tick %d", currentTime)
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
					// For debugging https://github.com/xephonhq/xephon-k/issues/33 collector always have only 10 points on server
					log.Debug(string(serializer.Data()))

					// FIXME: sending request would block following collector
					res, err := client.Do(req)
					if err != nil {
						log.Warn(err)
					} else {
						io.Copy(ioutil.Discard, res.Body)
						res.Body.Close()
						log.Info("flushed")
					}
					serializer.Reset()
					currentBatchSize = 0
					// FIXME: should figure out a better way to just rest the points
					// reset all the series https://github.com/xephonhq/xephon-k/issues/33
					for _, m := range metricNames {
						seriesMap[m] = common.NewIntSeries(m)
					}
				} else {
					currentBatchSize++
				}
				// TODO: actually they should be updated concurrently
				cpuCollector.Update()
				memCollector.Update()
				var s *common.IntSeries
				if enableCPU {
					// FIXME: this should not be human work ....
					for i, p := range cores {
						var stat system.CPUStat
						if i != hostInfo.NumCores {
							stat = cpuCollector.CPUs[i]
						} else {
							stat = cpuCollector.CPUTotal
						}
						s = seriesMap[p+"user"]
						// FIXME: .... it's too hard to add points to metrics, need a lot of copy and paste code
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.User)})
						s = seriesMap[p+"nice"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.Nice)})
						s = seriesMap[p+"system"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.System)})
						s = seriesMap[p+"idle"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.Idle)})
						s = seriesMap[p+"iowait"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.IOWait)})
						s = seriesMap[p+"irq"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.Irq)})
						s = seriesMap[p+"softirq"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.SoftIrq)})
						s = seriesMap[p+"steal"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.Steal)})
						s = seriesMap[p+"guest"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.Guest)})
						s = seriesMap[p+"guestnice"]
						s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(stat.GuestNice)})
					}
				}
				s = seriesMap["mem.total"]
				s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(memCollector.MemTotal)})
				s = seriesMap["mem.free"]
				s.Points = append(s.Points, common.IntPoint{T: currentTime, V: int64(memCollector.MemFree)})
				log.Debugf("mem.free length %d", len(s.Points))
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
	// local flags
	// TODO: there is problem of sharing variables between global commands, though they would never get executed together
	CollectorCmd.Flags().StringVar(&db, "db", "xephonk", "target database: xephonk|influxdb|kairosdb")
}
