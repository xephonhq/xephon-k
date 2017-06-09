package daemon

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	common "github.com/xephonhq/xephon-k/pkg/cmd"
	cf "github.com/xephonhq/xephon-k/pkg/config"
	"github.com/xephonhq/xephon-k/pkg/server/grpc"
	"github.com/xephonhq/xephon-k/pkg/server/http"
	"github.com/xephonhq/xephon-k/pkg/server/service"
	"github.com/xephonhq/xephon-k/pkg/storage"
	"github.com/xephonhq/xephon-k/pkg/util"
	"gopkg.in/yaml.v2"
	"os/signal"
	"syscall"
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

		var (
			store      storage.Store
			err        error
			httpServer *http.Server
			grpcServer *grpc.Server
		)

		log.Infof("use %s storage", cfgStorage)
		store, err = storage.CreateStore(cfgStorage, config.Storage)
		if err != nil {
			log.Fatalf("can't create store %v", err)
			return
		}
		writeService := service.NewWriteService(store)
		readService := service.NewReadService(store)

		sigInt := make(chan os.Signal)
		sigTerm := make(chan os.Signal)
		signal.Notify(sigInt, os.Interrupt)
		signal.Notify(sigTerm, syscall.SIGTERM)
		// TODO: #56 when one server has error, another server will still try to start, but the process (main goroutine) exit
		// so it does not cause much trouble, but we should be able control that. this also why there is nil check in
		// http and grpc's stop function
		serverErr := make(chan error)

		if config.Server.Http.Enabled {
			httpServer = http.NewServer(config.Server.Http, writeService, readService)
			go func() {
				if err := httpServer.Start(); err != nil {
					serverErr <- err
				}
			}()
		}
		if config.Server.Grpc.Enabled {
			grpcServer = grpc.NewServer(config.Server.Grpc, writeService)
			go func() {
				if err := grpcServer.Start(); err != nil {
					serverErr <- err
				}
			}()
		}

		select {
		case <-sigInt:
			log.Info("received SIGINT, exiting gracefully")
		case <-sigTerm:
			log.Info("received SIGTERM, exiting gracefully")
		case err := <-serverErr:
			log.Errorf("server error %v", err)
		}

		if config.Server.Http.Enabled {
			httpServer.Stop()
		}
		if config.Server.Grpc.Enabled {
			grpcServer.Stop()
		}

		store.Shutdown()

		log.Info("See you!")
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
	RootCmd.AddCommand(SchemaCmd)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", defaultConfigFile, "specify config file location")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
	RootCmd.PersistentFlags().StringVar(&cfgStorage, "storage", defaultStorage, "storage backend: null|memory|disk|cassandra")
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
