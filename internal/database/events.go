package database

import "database/sql"

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id      int `json:"id"`
	OwnerId int `json:"ownerId"`

	// binding:"required" = field empty ලා submit කරන්න බෑ
	// binding:"min=3" = අවම characters 3ක්
	Name string `json:"name" binding:"required,min=3"`

	Description string `json:"description" binding:"required,min=10"`

	// datetime=2006-01-02 = date format validate කිරීම
	// Go වල special date: January 2, 2006 (Mon Jan 2 15:04:05 MST 2006)
	Date string `json:"date" binding:"required,datetime=2006-01-02"`

	Location string `json:"location" binding:"required,min=3"`
}
