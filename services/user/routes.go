package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/moabdelazem/ecom/services/auth"
	"github.com/moabdelazem/ecom/types"
	"github.com/moabdelazem/ecom/utils"
)

// Handler is a struct that handles the user routes.
type Handler struct {
	store types.UserStore
}

// NewHandler creates a new instance of the user handler.
func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

// RegisterRoutes registers the user routes with the router.
// It maps the handler functions to the corresponding HTTP methods and paths.
func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
	r.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
}

// handleLogin is the handler function for the login endpoint.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Implement the login logic here
}

// handleRegister is the handler function for the register endpoint.
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Check JSON Payload
	var payload types.RegisterRequestUser
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate The Payload
	if err := utils.Validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Check If The User Already Exists
	if _, err := h.store.GetUserByEmail(payload.Email); err != nil {
		utils.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	// Hash The Password Before Saving It
	hashedPassword, err := auth.HashPassword(payload.Password)

	// Check If The Password Hashing Failed
	if err != nil {
		utils.ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	// Create The User if the user does not exist
	u := &types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	}
	if err := h.store.CreateUser(u); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}
