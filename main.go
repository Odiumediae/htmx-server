package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

type Film struct {
  Title string
  Director string
}

func main () {
  port := os.Getenv("PORT")
  if port == "" {
      port = "8080"
  }

  http.HandleFunc("/", func (w http.ResponseWriter, _ *http.Request) {
    data := map[string][]Film {
      "Films": {
        {Title: "The Godfather", Director: "Francis Ford Coppola"},
        {Title: "Blade Runner", Director: "Ridley Scott"},
        {Title: "The Thing", Director: "John Carpenter"},
      },
    }
    t.ExecuteTemplate(w, "index.html", data)
  })

  http.HandleFunc("/add-film/", func (w http.ResponseWriter, r *http.Request) {
    title := r.PostFormValue("title")
    director := r.PostFormValue("director")
    listItem := fmt.Sprintf("<li class=\"list-group-item bg-primary text-white\">%s: %s</li>", title, director)
    tmpl, _ := template.New("t").Parse(listItem)
    tmpl.Execute(w, nil)
  })

  log.Println("listening on", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
