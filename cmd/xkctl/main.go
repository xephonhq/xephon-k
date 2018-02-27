package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"
	"google.golang.org/grpc"

	myrpc "github.com/xephonhq/xephon-k/xk/transport/grpc"
	"github.com/xephonhq/xephon-k/xk/util/logutil"
)

const (
	myname = "xkctl"
)

var log = logutil.Registry

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}
var client myrpc.XephonkClient
var addr = "localhost:2334"

func main() {
	cli := icli.New(
		icli.Name(myname),
		icli.Description("Xephonk ctrl"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	root.AddCommand(pingCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mustCreateClient() {
	if client != nil {
		return
	}
	if conn, err := grpc.Dial(addr, grpc.WithInsecure()); err != nil {
		log.Fatalf("can't dial %v", err)
	} else {
		client = myrpc.NewClient(conn)
	}
}
