package main

import (
	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/xk/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start Xephonk daemon",
	Long:  "Start Xephonk daemon with gRPC and HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		mustLoadConfig()
		mgr, err := server.NewManager(cfg)
		if err != nil {
			log.Fatal(err)
		}
		if err := mgr.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
