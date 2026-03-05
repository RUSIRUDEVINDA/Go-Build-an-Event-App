package main

// server.go — HTTP server configuration and startup
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *application) serve() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.port), // ":8080"
		Handler: app.routes(),                 // Routes handler

		// Timeouts — terminating slow requests
		IdleTimeout:  time.Minute,      // Idle connection limit
		ReadTimeout:  10 * time.Second, // Request read limit
		WriteTimeout: 30 * time.Second, // Response write limit
	}

	log.Printf("Starting server on port %d", app.port)
	return server.ListenAndServe() // Server start!
}
