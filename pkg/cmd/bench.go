package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg"
	"github.com/xephonhq/xephon-k/pkg/bench"
	"github.com/xephonhq/xephon-k/pkg/bench/loader"
	"github.com/xephonhq/xephon-k/pkg/bench/reporter"
)

var (
	yes         = false
	db          = "xephonk"
	concurrency = 10
	batchSize   = 100
	timeout     = 30
)

var BenchCmd = &cobra.Command{
	Use:   "xkb",
	Short: "Xephon K Benchmark",
	Long:  "xkb is the bechmark tool for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Xephon K Bench %s \n", pkg.Version)
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
		config := loader.NewConfig(targetDB)
		fmt.Print(config)
		if !yes {
			fmt.Print("Do you want to proceed? [Y/N]")
			var choice string
			// TODO: we should only wait for a limit amount of time
			fmt.Scanf("%s", &choice)
			if strings.ToLower(choice) == "n" {
				fmt.Print("you said no, bye~\n")
				return
			}
		}
		loader := loader.NewHTTPLoader(config, &reporter.BasicReporter{})
		loader.Run()
		// print config again, after the report
		fmt.Print(config)
	},
}

func ExecuteBench() {
	if BenchCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	// global flags
	BenchCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
	BenchCmd.PersistentFlags().BoolVarP(&yes, "yes", "y", false, "yes to all, no more prompt")
	// local flags
	BenchCmd.Flags().StringVarP(&db, "db", "d", "xephonk", "target database: xephonk|influxdb|kairosdb")
	BenchCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "concurrency (worker gorutine number)")
	BenchCmd.Flags().IntVarP(&batchSize, "batch", "b", 100, "batch size")
	BenchCmd.Flags().IntVarP(&timeout, "timeout", "t", 30, "time out in seconds")
}
