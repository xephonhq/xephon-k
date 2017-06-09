package bench

import (
	"github.com/spf13/cobra"
	common "github.com/xephonhq/xephon-k/pkg/cmd"
	cf "github.com/xephonhq/xephon-k/pkg/config"
	"github.com/xephonhq/xephon-k/pkg/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var log = util.Logger.NewEntryWithPkg("k.cmd.bench")

const (
	defaultConfigFile = "xkb.yml"
	defaultTarget     = "xephonk"
)

var (
	config     cf.BenchConfig
	configFile = defaultConfigFile
	debug      = false
	target     = defaultTarget
)

var RootCmd = &cobra.Command{
	Use:   "xkb",
	Short: "Xephon K Benchmark",
	Long:  "xkb is the bechmark tool for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Hi bench")
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
	RootCmd.PersistentFlags().StringVar(&target, "target", defaultTarget, "target: xephonk|kariosdb|influxdb")
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
	config = cf.NewBench()
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
}
