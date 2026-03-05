package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

// Attendee — user_id + event_id link
type Attendee struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}
