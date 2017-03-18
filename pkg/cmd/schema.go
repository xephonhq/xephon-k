package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Create schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: create schema!")
	},
}

func init() {
	RootCmd.AddCommand(SchemaCmd)
}
