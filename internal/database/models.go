package database

import "database/sql"

// Models struct — සියලු database models එකතු කරලා
type Models struct {
    Users     UserModel     // User related operations
    Events    EventModel    // Event related operations
    Attendees AttendeeModel // Attendee related operations
}

// NewModels — Database connection models වලට pass කිරීම
func NewModels(db *sql.DB) Models {
    return Models{
        Users:     UserModel{DB: db},
        Events:    EventModel{DB: db},
        Attendees: AttendeeModel{DB: db},
    }
}