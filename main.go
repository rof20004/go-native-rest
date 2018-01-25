package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Servidor iniciado")
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
