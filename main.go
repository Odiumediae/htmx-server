package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
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
    time.Sleep(200 * time.Millisecond)
    title := r.PostFormValue("title")
    director := r.PostFormValue("director")
		t.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
  })

  log.Println("listening on", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
