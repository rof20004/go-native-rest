package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rof20004/go-native-rest/routes"
)

func main() {
	log.Println("Servidor iniciado")
	log.Fatalln(http.ListenAndServe(":8080", routes.GetServeMux()))
}
