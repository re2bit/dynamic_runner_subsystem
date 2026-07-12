package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

// ExampleFetch zeigt die Verwendung von Resty
func ExampleFetch() {
	client := resty.New()

	resp, err := client.R().
		Get("https://httpbin.org/get")

	if err != nil {
		log.Error().Err(err).Msg("Fehler beim HTTP-Request")
		return
	}

	log.Info().
		Int("status", resp.StatusCode()).
		Str("body", string(resp.Body()[:min(len(resp.Body()), 50)])).
		Msg("HTTP-Request erfolgreich")
}
