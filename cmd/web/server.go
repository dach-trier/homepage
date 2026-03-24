package main

import (
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(chi_middleware.Logger)

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./favicon.ico")
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("404 not found"))
	})

	log.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
