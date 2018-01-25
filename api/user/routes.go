package user

import (
	"net/http"
)

// SetRoutes - set user routes
func SetRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/users/list", ListUser)
	mux.HandleFunc("/api/v1/users/create", CreateUser)
	mux.HandleFunc("/api/v1/users/get/", GetUser)
}
