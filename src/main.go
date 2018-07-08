package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

type Config struct {
	Host  string
	Theme string
	Title string
	Port  uint16
}

var (
	tpl = template.Must(template.New("index.tmpl").ParseFiles(path.Join("src", "index.tmpl")))
)

func main() {
	cfg, err := config()
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir(path.Join("src", "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "index.tmpl", cfg); err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("Server has started at port %d\n", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))

}

func config() (*Config, error) {
	if port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16); err != nil {
		return nil, err
	} else {
		return &Config{
			Host:  os.Getenv("HOST"),
			Theme: os.Getenv("THEME"),
			Title: os.Getenv("TITLE"),
			Port:  uint16(port),
		}, nil
	}
}
