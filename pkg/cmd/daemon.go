package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/server"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
)

var (
	configFile        = ""
	defaultConfigFile = "xephon-k.yml"
	port              = server.DefaultPort
	backend           = "memory"
	cassandraHost     = "localhost"
)

var DaemonCmd = &cobra.Command{
	Use:   "xkd",
	Short: "Xephon K Daemon",
	Long:  "xkd is the server daemon for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		// a dirty way to create schema
		cassandra.CassandraHost = cassandraHost
		srv := server.Server{Port: port, Backend: backend, CassandraHost: cassandraHost}
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
	// NOTE: schema command also need this
	DaemonCmd.PersistentFlags().StringVar(&cassandraHost, "cassandra-host", "localhost", "cassandra host address")
	// local flags
	DaemonCmd.Flags().IntVarP(&port, "port", "p", server.DefaultPort, "port to listen on")
	DaemonCmd.Flags().StringVarP(&backend, "backend", "b", "memory", "memory|cassandra")
}
