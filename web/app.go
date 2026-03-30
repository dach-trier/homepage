package web

import (
	"log"
	"net/http"

	"html/template"

	chi "github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

// Global Application State
type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Router() http.Handler {
	router := chi.NewRouter()
	router.Use(chi_middleware.Logger)

	// assets

	fs := http.FileServer(http.Dir("web/assets/"))
	router.Get("/assets/*", http.StripPrefix("/assets/", fs).ServeHTTP)

	// pages

	router.Get("/", app.serveIndexPage)
	router.NotFound(app.serveNotFoundPage)

	return router
}

func (app *App) serveIndexPage(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template

	tmpl = template.New("")
	tmpl = template.Must(tmpl.ParseFiles("web/templates/index.html"))
	tmpl = template.Must(tmpl.ParseGlob("web/icons/*.svg"))

	tmpl.ExecuteTemplate(w, "page", map[string]any{})
}

func (app *App) serveNotFoundPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("404 not found"))

	if err != nil {
		log.Fatalf("failed to write 404 response: %v", err)
	}
}
