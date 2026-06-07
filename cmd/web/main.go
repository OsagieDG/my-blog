package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/OsagieDG/mlog/service/middleware"
	embedStatic "github.com/OsagieDG/osagiedg.me/static"
	tmplEmbed "github.com/OsagieDG/osagiedg.me/tmpl"
	"github.com/go-chi/chi/v5"
)

func main() {
	tmpl := parseTemplates()

	router := chi.NewRouter()

	router.Handle("/static/*", http.StripPrefix("/static/",
		http.FileServerFS(embedStatic.Static)))

	router.Get("/", handler("about", tmpl))
	router.Get("/about", handler("about", tmpl))
	router.Get("/hobbies", handler("hobbies", tmpl))
	router.Get("/posts", handler("posts", tmpl))
	router.Get("/post1", handler("post1", tmpl))
	router.Get("/post2", handler("post2", tmpl))
	router.Get("/post3", handler("post3", tmpl))
	router.Get("/post4", handler("post4", tmpl))
	router.Get("/post5", handler("post5", tmpl))
	router.Get("/post6", handler("post6", tmpl))
	router.Get("/projects", handler("projects", tmpl))

	mlog := middleware.MLog(
		middleware.LogResponse,
		middleware.RecoverPanic,
	)

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", mlog(router))
	if err != nil {
		slog.Error("Error starting server", "err", err)
		os.Exit(1)
	}
}

func handler(name string, tmpl map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, ok := tmpl[name]
		if !ok {
			http.Error(w, "Template not found", http.StatusNotFound)
			return
		}
		err := t.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			http.Error(w, "Internal Server Error",
				http.StatusInternalServerError)
			slog.Error("Error rendering template", "err", err)
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
	post5 := "post/post5/post5.tmpl"
	post6 := "post/post6/post6.tmpl"
	projects := "projects/projects.tmpl"

	tmpl := map[string]*template.Template{
		"about":    parseTemplateFiles(layout, about),
		"hobbies":  parseTemplateFiles(layout, hobbies),
		"posts":    parseTemplateFiles(layout, posts),
		"post1":    parseTemplateFiles(layout, post1),
		"post2":    parseTemplateFiles(layout, post2),
		"post3":    parseTemplateFiles(layout, post3),
		"post4":    parseTemplateFiles(layout, post4),
		"post5":    parseTemplateFiles(layout, post5),
		"post6":    parseTemplateFiles(layout, post6),
		"projects": parseTemplateFiles(layout, projects),
	}

	return tmpl
}

func parseTemplateFiles(layout, content string) *template.Template {
	tmpl, err := template.New("layout.html").ParseFS(tmplEmbed.Files,
		layout, content)
	if err != nil {
		slog.Error("Error parsing templates", "err", err)
		os.Exit(1)
	}
	return tmpl
}
