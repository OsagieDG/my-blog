package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var templates = map[string]*template.Template{}

func main() {
	templates = parseTemplates()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", handler("about"))
	r.Get("/about", handler("about"))
	r.Get("/tools", handler("tools"))
	r.Get("/posts", handler("posts"))
	r.Get("/blog1", handler("blog1"))

	// Start the server
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, ok := templates[name]
		if !ok {
			http.Error(w, "Template not found", http.StatusNotFound)
			return
		}
		err := tmpl.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error rendering template:", err)
		}
	}
}

func parseTemplates() map[string]*template.Template {
	layout := "tmpl/layout.tmpl"
	about := "tmpl/about.tmpl"
	tools := "tmpl/tools.tmpl"
	posts := "tmpl/posts.tmpl"
	blog1 := "tmpl/blog/blog1/blog1.tmpl"

	templates := map[string]*template.Template{
		"about": parseTemplateFiles(layout, about),
		"tools": parseTemplateFiles(layout, tools),
		"posts": parseTemplateFiles(layout, posts),
		"blog1": parseTemplateFiles(layout, blog1),
	}

	return templates
}

func parseTemplateFiles(files ...string) *template.Template {
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	return tmpl
}
