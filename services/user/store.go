package user

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/moabdelazem/ecom/types"
)

// Store is a struct that handles the database operations for the user.
type Store struct {
	db *sql.DB
}

// NewStore creates a new instance of the user store.
func NewStore(db *sql.DB) *Store {
	return &Store{}
}

/*
The CreateUser function creates a new user in the database. It takes a user struct as an argument and returns an error if the operation fails.
*/
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// Query the database to get the user by email
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	// Scan the row into a user struct
	u := new(types.User)
	for rows.Next() {
		// Scan the row into the user struct
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	// Check if the user ID is nil
	if u.ID == uuid.Nil {
		return nil, nil
	}

	return u, nil
}

/*
The GetUserByID function retrieves a user from the database by ID. It takes a UUID as an argument and returns a user struct and an error if the operation fails.
*/
func (s *Store) GetUserByID(id uuid.UUID) (*types.User, error) {
	// Query the database to get the user by ID
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	// Scan the row into a user struct
	u := new(types.User)
	for rows.Next() {
		// Scan the row into the user struct
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	// Check if the user ID is nil
	if u.ID == uuid.Nil {
		return nil, nil
	}

	return u, nil
}

/*
The CreateUser function creates a new user in the database. It takes a user struct as an argument and returns an error if the operation fails.
*/
func (s *Store) CreateUser(u *types.User) error {
	// Insert the user into the database
	_, err := s.db.Exec("INSERT INTO users (id, first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		u.ID, u.FirstName, u.LastName, u.Email, u.Password, u.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

/*
The scanRowIntoUser function is a utility function that scans the rows into a user struct.
*/
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	// Create a new user struct
	user := new(types.User)

	// Scan the row into the user struct
	err := rows.Scan(&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	// Check if there is an error
	if err != nil {
		return nil, err
	}
	return user, nil
}
