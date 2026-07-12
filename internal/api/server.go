package api

import (
	//	"forgejo_event_listener/internal/client"
	//	"net/http"

	"forgejo_event_listener/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// StartServer initialisiert und startet den Gin-Server
func StartServer() {
	// Gin in den Release-Modus versetzen, wenn gewünscht
	if config.GetEnv("GIN_MODE", "debug") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Standard Middlewares
	r.Use(gin.Recovery())

	// Logger Middleware (optional, wir können auch Zerolog direkt nutzen)
	r.Use(func(c *gin.Context) {
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Msg("API Request")
		c.Next()
	})

	// Beispiel Route
	/*
		   		r.GET("/health", func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{
						"status": "ok",
					})
				})

				r.GET("/fetch", func(c *gin.Context) {
				client.ExampleFetch()
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
			})

	*/

	r.GET("/*param", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	port := config.GetEnv("PORT", "8080")
	log.Info().Msgf("Server startet auf Port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal().Err(err).Msg("Fehler beim Starten des Servers")
	}
}
