package user

import (
	"net/http"
)

const baseURL = "/api/v1/users/"

// SetRoutes - set user routes
func SetRoutes(mux *http.ServeMux) {
	mux.HandleFunc(baseURL, Resources)
}
