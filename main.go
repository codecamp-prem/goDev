package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/codecamp-prem/goDev/controllers"
	"github.com/codecamp-prem/goDev/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))

	r.Get("/faq", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found!", http.StatusNotFound)
	})
	fmt.Println("Starting Server at 3000...")
	http.ListenAndServe(":3000", r)
}
