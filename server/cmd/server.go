package cmd

import (
	"goweibo/bootstrap"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run app server",
	Run: func(cmd *cobra.Command, args []string) {
		// init db
		bootstrap.SetupDB()
		// init server
		bootstrap.SetupServer()
		bootstrap.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
