package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adrg/xdg"
	"github.com/envelope-zero/backend/v4/pkg/controllers"
	"github.com/envelope-zero/backend/v4/pkg/database"
	"github.com/envelope-zero/backend/v4/pkg/models"
	"github.com/envelope-zero/backend/v4/pkg/router"
	"github.com/gin-contrib/static"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skratchdot/open-golang/open"
)

// Embedding code taken from https://github.com/gin-contrib/static/issues/19#issue-830589998
//
// Embed the "public" directory as filesystem
//
//go:embed public
var server embed.FS

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(_, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	dbPath, err := xdg.DataFile("envelope-zero/envelope-zero.db")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Str("database file", dbPath).Msg("Init")

	dbConnectionOptions := "_pragma=foreign_keys(1)"

	db, err := database.Connect(fmt.Sprintf("%s?%s", dbPath, dbConnectionOptions))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = models.Migrate(db)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	url, _ := url.Parse("http://localhost:3200/api")

	// Set the DB context and add it to the controller
	ctx := context.Background()
	ctx = context.WithValue(ctx, database.ContextURL, url)
	controller := controllers.Controller{DB: db.WithContext(ctx)}

	r, teardown, err := router.Config(url)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer teardown()

	// Attach all backend routes to /api
	router.AttachRoutes(controller, r.Group("/api/"))

	// Serve the frontend on /
	r.Use(static.Serve("/", EmbedFolder(server, "public")))

	// The following code is taken from https://github.com/gin-gonic/examples/blob/91fb0d925b3935d2c6e4d9336d78cf828925789d/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	// and has been modified to not wait for the timeout to expire before returning
	srv := &http.Server{
		Addr:    ":3200",
		Handler: r,
	}

	_ = open.Run("http://localhost:3200")

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
