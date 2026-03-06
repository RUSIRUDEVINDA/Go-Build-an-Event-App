package main

// events.go — file that defines API handlers for event-related routes
import (
	"backend/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Event Handler
func (app *application) createEvent(c *gin.Context) {
	var event database.Event

	// ShouldBindJSON — Request body JSON parse කිරීම + validation
	// binding tags invalid නම් error return කරනවා
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 400 Bad Request — Client ගෙ data problem
		return
	}

	// Database insert
	// app.models.Events.Insert() — database method call කිරීම
	err := app.models.Events.Insert(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		// 500 Internal Server Error — Server side problem
		return
	}

	c.JSON(http.StatusCreated, event) // 201 Created — Success!
}

// Get All Events Handler
func (app *application) getAllEvents(c *gin.Context) {
	events, err := app.models.Events.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}
	c.JSON(http.StatusOK, events) // 200 OK
}

// Get Single Event Handler
func (app *application) getEvent(c *gin.Context) {
	// c.Param("id") = URL ඉදන් id extract කිරීම
	// /events/5 → "5" (string)
	id, err := strconv.Atoi(c.Param("id")) // String → Integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := app.models.Events.Get(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		// 404 Not Found
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// Update Event Handler
func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	// First check — event exists?
	existingEvent, err := app.models.Events.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}
	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Request body bind
	updateEvent := &database.Event{}
	if err := c.ShouldBindJSON(&updateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateEvent.Id = id // URL ඉදන් ගත් id set කිරීම

	if err := app.models.Events.Update(updateEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, updateEvent)
}

// Delete Event Handler
func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	// First check — event exists?
	existingEvent, err := app.models.Events.Get(id)
	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusNoContent, nil) // 204 — Success, no content return
}
