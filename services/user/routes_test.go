package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/moabdelazem/ecom/types"
)

func TestUserServiceHanlders(t *testing.T) {
	// Create a new instance of the mock user store
	userStore := &mockUserStore{}
	// Create a new instance of the user handler
	handler := NewHandler(userStore)

	// Test the Register endpoint
	// Invalid email
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		// Create a new instance of the register request payload
		payload := types.RegisterRequestUser{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "hosdacom", // Invalid email
			Password:  "password",
		}
		// Marshal the payload into JSON
		JSONData, _ := json.Marshal(payload)

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(JSONData))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP response recorder
		rr := httptest.NewRecorder()
		// Create a new router
		router := mux.NewRouter()

		// Register the routes
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		// Serve the HTTP request
		router.ServeHTTP(rr, req)

		// Check if the status code is 400
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	// Test the Register endpoint
	// Correct email
	t.Run("should correctly register the user", func(t *testing.T) {
		// Create a new instance of the register request payload
		payload := types.RegisterRequestUser{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndow@gmail.com", // Correctn email
			Password:  "password",
		}
		// Marshal the payload into JSON
		JSONData, _ := json.Marshal(payload)

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(JSONData))
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP response recorder
		rr := httptest.NewRecorder()
		// Create a new router
		router := mux.NewRouter()

		// Register the routes
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		// Serve the HTTP request
		router.ServeHTTP(rr, req)

		// Check if the status code is 400
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) GetUserByID(id uuid.UUID) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u *types.User) error {
	return nil
}
