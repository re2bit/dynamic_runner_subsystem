package cmd

import (
	"os"

	"forgejo_event_listener/internal/config"
	"forgejo_event_listener/internal/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forgejo_event_listener",
	Short: "Ein Listener für Forgejo Events",
	Long:  `Eine Applikation, die auf Webhooks von Forgejo reagiert und diese verarbeitet.`,
}

// Execute führt das Root-Kommando aus
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialisierungen, die vor jedem Kommando laufen sollen
	cobra.OnInitialize(func() {
		logger.Init()
		config.Load()
	})
}
