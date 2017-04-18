package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/collector"
	"os"
	"os/signal"
	"time"
)

var CollectorCmd = &cobra.Command{
	Use:   "xkc",
	Short: "Xephon K Collector",
	Long:  "xkc is the metrics collector for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		// let's just assume we only report to Xephon-K
		// TODO:
		// - set the timer
		// - serialize the payload
		// need to have a look at other client
		config := collector.NewConfig()
		tickChan := time.NewTicker(config.Interval).C

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
