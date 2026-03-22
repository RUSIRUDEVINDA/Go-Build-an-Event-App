package database

import (
	"context"
	"database/sql"
	"time"
)

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

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id`
	err := m.DB.QueryRowContext(ctx, stmt, user.Email, user.Password, user.Name).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

// Add get user by id method
func (m *UserModel) Get(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE id = $1`

	var user User
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

//get user by email method
func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `SELECT id, email, password, name FROM users WHERE email = $1`

	var user User
	err := m.DB.QueryRow(query, email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Name,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
