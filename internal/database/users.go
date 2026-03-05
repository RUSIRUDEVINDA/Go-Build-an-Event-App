package database

import "database/sql"

// UserModel — Database connection hold කරනවා
type UserModel struct {
	DB *sql.DB // Pointer to database connection
}

// User — Database users table ගේ Go representation
type User struct {
	Id       int    `json:"id"`    // JSON response: "id"
	Email    string `json:"email"` // JSON response: "email"
	Name     string `json:"name"`  // JSON response: "name"
	Password string `json:"-"`     // json:"-" = NEVER send in response!
	// Password ගෙ ඉදිරියේ "-" දෙව්වොත් JSON response
	// එකේ password field show වෙන්නේ නෑ — Security!
}
