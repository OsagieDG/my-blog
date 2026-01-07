package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/OsagieDG/mlog/service/middleware"
	embedStatic "github.com/OsagieDG/osagiedg.me/static"
	tmplEmbed "github.com/OsagieDG/osagiedg.me/tmpl"
	"github.com/go-chi/chi/v5"
)

var tmpl = map[string]*template.Template{}

func main() {
	tmpl = parseTemplates()

	router := chi.NewRouter()

	router.Handle("/static/*", http.StripPrefix("/static/",
		http.FileServerFS(embedStatic.Static)))

	router.Get("/", handler("about"))
	router.Get("/about", handler("about"))
	router.Get("/hobbies", handler("hobbies"))
	router.Get("/posts", handler("posts"))
	router.Get("/post1", handler("post1"))
	router.Get("/post2", handler("post2"))
	router.Get("/post3", handler("post3"))
	router.Get("/post4", handler("post4"))
	router.Get("/projects", handler("projects"))

	mlog := middleware.MLog(
		middleware.LogResponse,
		middleware.RecoverPanic,
	)

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", mlog(router))
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, ok := tmpl[name]
		if !ok {
			http.Error(w, "Template not found", http.StatusNotFound)
			return
		}
		err := tmpl.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			http.Error(w, "Internal Server Error",
				http.StatusInternalServerError)
			log.Println("Error rendering template:", err)
		}
	}
}

func parseTemplates() map[string]*template.Template {
	layout := "layout.html"
	about := "about.tmpl"
	hobbies := "hobbies.tmpl"
	posts := "posts.tmpl"
	post1 := "post/post1/post1.tmpl"
	post2 := "post/post2/post2.tmpl"
	post3 := "post/post3/post3.tmpl"
	post4 := "post/post4/post4.tmpl"
	projects := "projects/projects.tmpl"

	tmpl := map[string]*template.Template{
		"about":    parseTemplateFiles(layout, about),
		"hobbies":  parseTemplateFiles(layout, hobbies),
		"posts":    parseTemplateFiles(layout, posts),
		"post1":    parseTemplateFiles(layout, post1),
		"post2":    parseTemplateFiles(layout, post2),
		"post3":    parseTemplateFiles(layout, post3),
		"post4":    parseTemplateFiles(layout, post4),
		"projects": parseTemplateFiles(layout, projects),
	}

	return tmpl
}

func parseTemplateFiles(layout, content string) *template.Template {
	tmpl, err := template.New("layout.html").ParseFS(tmplEmbed.Files,
		layout, content)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	return tmpl
}
