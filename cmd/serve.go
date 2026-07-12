package cmd

import (
	"forgejo_event_listener/internal/api"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Startet den API Server",
	Run: func(cmd *cobra.Command, args []string) {
		// Startet den Gin Server
		api.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
