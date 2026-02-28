package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IhsanAlhakim/go-auth-api/internal/config"
	"github.com/IhsanAlhakim/go-auth-api/internal/database"
	"github.com/IhsanAlhakim/go-auth-api/internal/handlers"
	"github.com/IhsanAlhakim/go-auth-api/internal/middlewares"
	"github.com/IhsanAlhakim/go-auth-api/internal/mux"
	"github.com/IhsanAlhakim/go-auth-api/internal/routes"
)

var db *sql.DB

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Panicf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	mux := mux.New()

	m := middlewares.New(cfg)

	h := handlers.New(db, cfg)

	routes.Register(mux, m, h)

	server := new(http.Server)
	server.Addr = ":" + cfg.Port
	server.Handler = mux

	go func() {
		log.Println("Server started at localhost:" + cfg.Port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
		log.Println("Stopped serving new connection.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete")
}
