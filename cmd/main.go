package main

import (
	"go-svelte/cache"
	"go-svelte/config"
	"go-svelte/web"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	version = "source"
)

func setLogger() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	if os.Getenv("DEBUG") != "" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		return
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func main() {
	setLogger()
	config.InitConfig()
	config.InitSecrets()
	cache.InitCache()
	log.Info().Msgf("Starting TradeMaster (Version: %s)", version)
	if os.Getenv("DEBUG") == "" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.Info().
			Str("path", c.Request.URL.Path).
			Str("action", c.Request.Method).
			Int("status", c.Writer.Status()).
			Dur("latency", latency).
			Msg("Request")
	})
	staticDir := "static"
	if os.Getenv("STATIC_DIR") != "" {
		staticDir = os.Getenv("STATIC_DIR")
	}

	web.HandleRoutes(router, staticDir)
	port := os.Getenv("LISTEN_PORT")
	listenPort := ":8888"
	if port != "" {
		listenPort = ":" + port
	}
	log.Info().Msgf("Listening on %s", listenPort)
	log.Fatal().Err(router.Run(listenPort))
}
