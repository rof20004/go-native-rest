package user

import (
	"net/http"
)

// SetRoutes - set user routes
func SetRoutes(mux *http.ServeMux) {
	mux.HandleFunc(List, ListUser)
	mux.HandleFunc(Create, CreateUser)
	mux.HandleFunc(Get, GetUser)
	mux.HandleFunc(Update, UpdateUser)
}
