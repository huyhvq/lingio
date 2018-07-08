package cmd

import (
	"github.com/spf13/cobra"
	"github.com/huyhvq/lingio/server"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs API service",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer()
		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
