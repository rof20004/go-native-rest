package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/rof20004/go-native-rest/helpers"
)

// Resources API
func Resources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch r.Method {
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, baseURL)
		if id != "" {
			GetUser(w, r)
		} else {
			ListUser(w, r)
		}
	case "POST":
		CreateUser(w, r)
	case "PUT":
		UpdateUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		helpers.Response(w, http.StatusNotFound, helpers.ResourceNotFound, nil)
	}
}

// ListUser endpoint
func ListUser(w http.ResponseWriter, r *http.Request) {
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
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, baseURL))
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
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, baseURL))
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
