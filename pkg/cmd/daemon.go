package cmd

import (
	"os"

	"fmt"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/server"
)

var (
	configFile        = ""
	defaultConfigFile = "xephon-k.yml"
	port              = server.DefaultPort
	backend           = "memory"
	cassandraHost     = "localhost"
	diskFolder        = ""
)

// xkd -b disk --folder /home/at15/workspace/tmp
var DaemonCmd = &cobra.Command{
	Use:   "xkd",
	Short: "Xephon K Daemon",
	Long:  "xkd is the server daemon for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(Banner)
		srv := server.HTTPServer{Port: port, Backend: backend, CassandraHost: cassandraHost, DiskFolder: diskFolder}
		srv.Start()
		// TODO: capture ctrl + c and call shutdown of store
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
	// NOTE: schema command also need this
	DaemonCmd.PersistentFlags().StringVar(&cassandraHost, "cassandra-host", "localhost", "cassandra host address")
	// local flags
	DaemonCmd.Flags().IntVarP(&port, "port", "p", server.DefaultPort, "port to listen on")
	DaemonCmd.Flags().StringVarP(&backend, "backend", "b", "memory", "memory|cassandra|disk")
	DaemonCmd.Flags().StringVar(&diskFolder, "folder", "", "root folder of disk ")
}
