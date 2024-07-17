package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var templates = map[string]*template.Template{
	"index": parseTemplate("tmpl/index.tmpl"),
	"about": parseTemplate("tmpl/about.tmpl"),
	"posts": parseTemplate("tmpl/posts.tmpl"),
	"blog1": parseTemplate("tmpl/blog/blog1/blog1.tmpl"),
	"blog2": parseTemplate("tmpl/blog/blog2/blog2.tmpl"),
	"blog3": parseTemplate("tmpl/blog/blog3/blog3.tmpl"),
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", handler("index"))
	r.Get("/about", handler("about"))
	r.Get("/posts", handler("posts"))
	r.Get("/blog1", handler("blog1"))
	r.Get("/blog2", handler("blog2"))
	r.Get("/blog3", handler("blog3"))

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates[name].Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error rendering template:", err)
		}
	}
}

func parseTemplate(filePath string) *template.Template {
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Fatalf("Error parsing template %s: %v", filePath, err)
	}
	return tmpl
}
