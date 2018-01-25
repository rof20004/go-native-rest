package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/rof20004/go-native-rest/helpers"
)

// ListUser endpoint
func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		helpers.Response(w, http.StatusMethodNotAllowed, "Método não permitido", nil)
		return
	}

	users, err := list()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		helpers.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.Response(w, http.StatusOK, nil, users)
}

// CreateUser endpoint
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		helpers.Response(w, http.StatusMethodNotAllowed, "Método não permitido", nil)
		return
	}

	user := new(User)
	json.NewDecoder(r.Body).Decode(user)

	if user.validate("insert") {
		err := create(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			helpers.Response(w, http.StatusBadRequest, err.Error(), nil)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		helpers.Response(w, http.StatusBadRequest, "Dados inválidos", nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.Response(w, http.StatusOK, nil, user)
}

// GetUser endpoint
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		helpers.Response(w, http.StatusMethodNotAllowed, "Método não permitido", nil)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, Get))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		helpers.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := read(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		helpers.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.Response(w, http.StatusOK, nil, user)
}

// UpdateUser endpoint
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		helpers.Response(w, http.StatusMethodNotAllowed, "Método não permitido", nil)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, Update))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		helpers.Response(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.ID = int64(id)

	if user.validate("update") {
		err := update(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			helpers.Response(w, http.StatusBadRequest, err.Error(), nil)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		helpers.Response(w, http.StatusBadRequest, "Dados inválidos", nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.Response(w, http.StatusOK, nil, user)
}
