package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moabdelazem/ecom/services/user"
)

/*
APIServer struct contains the address of the server and a pointer to the database
*/
type APIServer struct {
	addr string
	db   *sql.DB
}

/*
NewServer is a constructor for APIServer struct
*/
func NewServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

/*
Run method of APIServer struct
*/
func (s *APIServer) Run() error {
	// Initialize a new router
	r := mux.NewRouter()
	sr := r.PathPrefix("/api/v1").Subrouter()
	// Initialize the user store
	userStore := user.NewStore(s.db)
	// Initialize the user handler
	userHandler := user.NewHandler(userStore)
	// Register the routes
	userHandler.RegisterRoutes(sr)

	// Log the server is running
	log.Println("Server is running on port", s.addr)

	return http.ListenAndServe(s.addr, r)
}
