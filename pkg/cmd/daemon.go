package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/server"
)

var (
	configFile        = ""
	defaultConfigFile = "xephon-k.yml"
	port              = 8080
	backend           = "memory"
)

var DaemonCmd = &cobra.Command{
	Use:   "xkd",
	Short: "Xephon K Daemon",
	Long:  "xkd is the server daemon for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.Server{Port: port, Backend: backend}
		srv.Start()
	},
}

func ExecuteDaemon() {
	if DaemonCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	// global flags
	DaemonCmd.PersistentFlags().StringVar(&configFile, "config", defaultConfigFile, "config file")
	DaemonCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
	// local flags
	DaemonCmd.Flags().IntVarP(&port, "port", "p", 8080, "port to listen on")
	DaemonCmd.Flags().StringVarP(&backend, "backend", "b", "memory", "memory|cassandra")
}