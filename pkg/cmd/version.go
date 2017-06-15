package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Xephon-K version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pkg.Version)
	},
}
