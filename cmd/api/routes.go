package main

// routes.go — file that defines API server routes
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()
	// gin.Default() — Logger + Recovery middleware included
	// Logger = logs requests
	// Recovery = catches panics and returns 500 instead of crashing

	return g
}
