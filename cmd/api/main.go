package main

// this is the main.go file for the API server. It sets up the application struct, connects to the database, and starts the server.
// The application struct holds the port number, JWT secret, and database models as dependencies for the server.
//  The main function initializes these dependencies and calls the serve method to start the server.
import (
	"backend/internal/database"
	"backend/internal/env"
	"database/sql"
	"log"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// application struct — API server's dependencies holding struct
type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

// main — API server startup function
func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),                           // PORT environment variable provides port number, default 8080
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-default-secret"), // JWT_SECRET environment variable provides secret key, default "some-default-secret"
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
