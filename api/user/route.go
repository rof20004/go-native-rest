package user

import (
	"net/http"
)

// SetRoutes - set user routes
func SetRoutes(mux *http.ServeMux) {
	const baseURL = "/api/v1/users/"
	mux.HandleFunc(baseURL, ListUser)
	mux.HandleFunc(baseURL, CreateUser)
	mux.HandleFunc(baseURL, GetUser)
	mux.HandleFunc(baseURL, UpdateUser)
}
