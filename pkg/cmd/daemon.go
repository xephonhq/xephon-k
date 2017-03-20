package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/server"
)

var (
	configFile        = ""
	defaultConfigFile = "xephon-k.yml"
	debug             = false
)

var DaemonCmd = &cobra.Command{
	Use:   "xkd",
	Short: "Xephon K Daemon",
	Long:  "xkd is the server daemon for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Xephon K Daemon:" + pkg.Version + " Use `xkd -h` for more information")
		srv := server.Server{}
		srv.Start()
	},
}

func ExecuteDaemon() {
	if DaemonCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	DaemonCmd.PersistentFlags().StringVar(&configFile, "config", defaultConfigFile, "config file")
	DaemonCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
}
