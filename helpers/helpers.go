package helpers

import (
	"encoding/json"
	"net/http"
)

// Response - formatt response message
func Response(w http.ResponseWriter, code int, message interface{}, data interface{}) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
