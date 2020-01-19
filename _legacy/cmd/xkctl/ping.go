package main

import (
	"context"

	"github.com/spf13/cobra"

	pb "github.com/xephonhq/xephon-k/xk/transport/grpc"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping xephonk",
	Long:  "Ping Xephonk server using gRPC",
	Run: func(cmd *cobra.Command, args []string) {
		mustCreateClient()
		if res, err := client.Ping(context.Background(), &pb.PingReq{Message: "ping from xkctl"}); err != nil {
			log.Fatal(err)
		} else {
			log.Infof("ping finished central response is %s", res.Message)
		}
	},
}
