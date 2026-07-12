package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Load lädt die Umgebungsvariablen aus einer .env Datei
func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("Keine .env Datei gefunden oder Fehler beim Laden")
	} else {
		log.Info().Msg(".env Datei erfolgreich geladen")
	}
}

// GetEnv ist ein Helper, um Umgebungsvariablen mit einem Default-Wert abzurufen
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
