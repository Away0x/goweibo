package cmd

import (
  "goweibo/bootstrap"
  "goweibo/database/factory"

	"github.com/spf13/cobra"
)

var mockCmd = &cobra.Command{
	Use:   "mock",
	Short: "mock data",
	Run: func(cmd *cobra.Command, args []string) {
    bootstrap.SetupDB()
		factory.Run()
	},
}

func init() {
	rootCmd.AddCommand(mockCmd)
}
