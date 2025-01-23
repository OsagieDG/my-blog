package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/osag1e/logstack/service/middleware"
)

var templates = map[string]*template.Template{}

func main() {
	templates = parseTemplates()

	router := chi.NewRouter()

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.Get("/", handler("about"))
	router.Get("/about", handler("about"))
	router.Get("/tools", handler("tools"))
	router.Get("/posts", handler("posts"))
	router.Get("/blog", handler("blog"))
	router.Get("/blog1", handler("blog1"))

	logstack := middleware.LogStack(
		middleware.LogRequest,
		middleware.LogResponse,
		middleware.RecoverPanic,
	)

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", logstack(router))
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
	blog := "tmpl/blog/og-blog.tmpl"
	blog1 := "tmpl/blog/blog1/blog1.tmpl"

	templates := map[string]*template.Template{
		"about": parseTemplateFiles(layout, about),
		"tools": parseTemplateFiles(layout, tools),
		"posts": parseTemplateFiles(layout, posts),
		"blog":  parseTemplateFiles(layout, blog),
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
