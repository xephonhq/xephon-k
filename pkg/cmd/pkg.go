package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.cmd")

func init() {
	cobra.OnInitialize(initConfig)

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
