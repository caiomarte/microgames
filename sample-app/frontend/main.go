package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Data struct {
	Title  string
	Player string
	Score  string
}

func main() {
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", Catalog)

	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func Render(w http.ResponseWriter, tmpl string, d Data) {
	tmpl = fmt.Sprintf("templates/%s", tmpl)

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, d)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

// -----------------------------------------
// REFERENCES
// https://github.com/Rosalita/GoViolin
// https://golang.org/doc/articles/wiki/
