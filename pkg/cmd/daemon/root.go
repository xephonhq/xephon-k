package daemon

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	common "github.com/xephonhq/xephon-k/pkg/cmd"
	cf "github.com/xephonhq/xephon-k/pkg/config"
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"

	//"github.com/xephonhq/xephon-k/pkg/server/grpc"
	"github.com/xephonhq/xephon-k/pkg/server/http"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"github.com/xephonhq/xephon-k/pkg/util"
	"gopkg.in/yaml.v2"
)

const (
	defaultConfigFile = "xkd.yml"
	defaultStorage    = "memory"
)

var (
	config     cf.DaemonConfig
	configFile = defaultConfigFile
	debug      = false
	cfgStorage = defaultStorage
)

var log = util.Logger.NewEntryWithPkg("k.cmd.daemon")

var RootCmd = &cobra.Command{
	Use:   "xkd",
	Short: "Xephon K Daemon",
	Long:  "xkd is the server daemon for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Print(c.Banner)
		// start http/grpc server
		// trap ctrl + c for shutdown

		// TODO: actually this logic should be handled in the storage package
		log.Debugf("use %s storage", cfgStorage)
		var (
			store      storage.Store
			err        error
			httpServer *http.Server
			// grpcServer *grpc.Server
		)
		if cfgStorage == "memory" {
			if err = memory.CreateStore(config.Storage.Memory); err != nil {
				log.Fatalf("can't create mem store %v", err)
				return
			}
			if store, err = memory.GetStore(); err != nil {
				log.Fatalf("can't get mem store %v", err)
				return
			}
		}
		// TODO: disk and cassandra
		// TODO: do we still need service? yes
		writeService := service.NewWriteService(store)
		if config.Server.Http.Enabled {
			httpServer = http.NewServer(config.Server.Http, writeService)
		}
		if config.Server.Grpc.Enabled {
			// grpcServer = grpc.NewServer(config.Server.Grpc, writeService)
		}
		// TODO: start another go routine, graceful shutdown etc.
		// TODO: start will block all the other things
		httpServer.Start()
	},
}

func Execute() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(common.VersionCmd)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", defaultConfigFile, "specify config file location")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
	RootCmd.PersistentFlags().StringVar(&cfgStorage, "storage", defaultStorage, "storage backend: memory|disk|cassandra")
}

func configFileError(err error) {
	if configFile != defaultConfigFile {
		log.Fatalf("can't read specified config file %s, got %v", configFile, err)
	}
	log.Warnf("use default config because can't read specified config file %s, got %v", configFile, err)
}

func initConfig() {
	if debug {
		util.UseVerboseLog()
	}
	config = cf.NewDaemon()
	if configFile == "" {
		return
	}
	// load the config when file is specified
	log.Debugf("load config file %s", configFile)
	f, err := os.OpenFile(configFile, os.O_RDONLY, 0666)
	if err != nil {
		configFileError(err)
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		configFileError(err)
		return
	}
	if err := yaml.Unmarshal(b, &config); err != nil {
		configFileError(err)
		return
	}
	if err := config.Apply(); err != nil {
		configFileError(err)
		return
	}
	// TODO: apply storage
	// TODO: debug flag should override config file
	// FIXME: what if trace is specified, calling use debug would result in trace log hidden
}
