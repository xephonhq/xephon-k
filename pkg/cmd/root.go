package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.cmd")

var (
	configFile        = ""
	defaultConfigFile = "xephon-k.yml"
	debug             = false
)

var RootCmd = &cobra.Command{
	Use:   "xkd",
	Short: "Xephon K Daemon",
	Long:  "xkd is the server daemon for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Xephon K Daemon:" + pkg.Version + " Use `xkd -h` for more information")
	},
}

func Execute() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", defaultConfigFile, "config file")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
}
func initConfig() {
	if debug {
		util.UseVerboseLog()
	}
	// TODO: configuration is not supported yet
	// viper.AutomaticEnv()
	// // TODO: check file existence
	// viper.SetConfigFile(configFile)
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Warn(err)
	// } else {
	// 	log.Debugf("config file %s is loaded", configFile)
	// }
}
