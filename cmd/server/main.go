package main

import (
	"log"
	"net/http"

	"github.com/dach-trier/homepage/internal/router"
)

func main() {
	server := http.Server{Addr: ":8080", Handler: router.NewRouter()}
	log.Println("Server is running at http://localhost:8080")
	err := server.ListenAndServe()
	log.Fatal(err)
}
