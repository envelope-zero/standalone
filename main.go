package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/envelope-zero/backend/pkg/database"
	"github.com/envelope-zero/backend/pkg/models"
	"github.com/envelope-zero/backend/pkg/router"
	"github.com/gin-contrib/static"
	"github.com/rs/zerolog/log"
)

func main() {
	err := database.Database()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = models.MigrateDatabase()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	// Set the API_URL
	os.Setenv("API_URL", "http://localhost:3200/api")

	r, err := router.Config()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	router.AttachRoutes(r.Group("/api/"))

	// Serve the static content
	r.Use(static.Serve("/", static.LocalFile("public", true)))

	// The following code is taken from https://github.com/gin-gonic/examples/blob/91fb0d925b3935d2c6e4d9336d78cf828925789d/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	// and has been modified to not wait for the timeout to expire before returning
	srv := &http.Server{
		Addr:    ":3200",
		Handler: r,
	}

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Str("event", "Error during startup").Err(err).Msg("Router")
		}
	}()
	log.Info().Str("event", "Startup complete").Msg("Router")

	<-quit
	log.Info().Str("event", "Received SIGINT or SIGTERM, stopping gracefully with 25 seconds timeout").Msg("Router")

	// Create a context with a 25 second timeout for the server to shut down in
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Str("event", "Graceful shutdown failed, terminating").Err(err).Msg("Router")
	}
	log.Info().Str("event", "Backend stopped").Msg("Router")
}
