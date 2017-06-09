package bench

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/bench2"
	common "github.com/xephonhq/xephon-k/pkg/cmd"
	cf "github.com/xephonhq/xephon-k/pkg/config"
	"github.com/xephonhq/xephon-k/pkg/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
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
		config.Loader.Target = target
		if !promptConfig() {
			return
		}
		scheduler, err := bench2.NewScheduler(config)
		if err != nil {
			log.Fatal(err)
			return
		}
		if err := scheduler.Run(); err != nil {
			log.Fatal(err)
			return
		}
		log.Info("bench finished")
	},
}

func Execute() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func promptConfig() bool {
	b, _ := yaml.Marshal(config)
	fmt.Print(string(b))
	fmt.Print("Do you want to proceed? [Y/N]")
	var choice string
	// TODO: we should only wait for a limit amount of time
	fmt.Scanf("%s", &choice)
	if strings.ToLower(choice) == "n" {
		fmt.Print("you said no, bye~\n")
		return false
	}
	return true
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
