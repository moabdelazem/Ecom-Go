package types

import (
	"time"

	"github.com/google/uuid"
)

/*
The types package contains the data types that are used in the application.
*/

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	CreateUser(*User) error
}

// RegisterRequestUser represents the request body for the register endpoint.
// It contains the first name, last name, email, and password of the user.
type RegisterRequestUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// User represents the user data model.
// It contains the ID, first name, last name, email, password, and created at fields.
type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
