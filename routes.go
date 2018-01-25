package main

import "net/http"

var mux = http.NewServeMux()

func init() {
	user := &User{}
	mux.HandleFunc("/api/v1/users/list", user.ListUser)
	mux.HandleFunc("/api/v1/users/create", user.CreateUser)
	mux.HandleFunc("/api/v1/users/get/", user.GetUser)
}
