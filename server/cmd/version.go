package cmd

import (
	"fmt"
	"goweibo/core"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\napp version = %s\n", core.GetConfig().String("APP.VERSION"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
