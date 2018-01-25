package routes

import (
	"net/http"

	"github.com/rof20004/go-native-rest/api/user"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()

	// User routes
	user.SetRoutes(mux)
}

// GetServeMux - get handler
func GetServeMux() *http.ServeMux {
	return mux
}
