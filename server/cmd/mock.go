package cmd

import (
	"goweibo/database/factory"

	"github.com/spf13/cobra"
)

var mockCmd = &cobra.Command{
	Use:   "mock",
	Short: "mock data",
	Run: func(cmd *cobra.Command, args []string) {
		factory.Run()
	},
}

func init() {
	rootCmd.AddCommand(mockCmd)
}
