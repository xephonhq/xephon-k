package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var BenchCmd = &cobra.Command{
	Use:   "xkb",
	Short: "Xephon K Benchmark",
	Long:  "xkb is the bechmark tool for Xephon K",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Xephon K Bench")
	},
}

func ExecuteBench() {
	if BenchCmd.Execute() != nil {
		os.Exit(-1)
	}
}
