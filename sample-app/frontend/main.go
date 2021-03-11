package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<div>%s</div>", p.Body)
}

func loadPage(title string) (*Page, error) {
	filepath := title + "/" + title + ".html"
	body, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// REFERENCES
// https://golang.org/doc/articles/wiki/
// https://github.com/Rosalita/GoViolin
