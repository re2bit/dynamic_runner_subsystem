package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initialisiert das globale Zerolog-Logger-Setup
func Init() {
	// Konsolenausgabe für bessere Lesbarkeit während der Entwicklung
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()

	// Standardmäßig auf Info-Level setzen, falls nicht anders konfiguriert
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
