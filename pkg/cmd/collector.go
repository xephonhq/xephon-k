package cmd

import (
	"github.com/spf13/cobra"
	"os"
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
