package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var users = []User{}

// ListUser endpoint
func ListUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// CreateUser endpoint
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user := User{}
	json.NewDecoder(r.Body).Decode(&user)

	if user.ID != 0 && strings.TrimSpace(user.Name) != "" {
		users = append(users, user)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Usuário inválido"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

// GetUser endpoint
func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/users/get/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao realizar a consulta"))
		return
	}

	fmt.Println(id)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func sayHello() {

}
