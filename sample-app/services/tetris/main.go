package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
}

func main() {
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.HandleFunc("/", handler)
    http.ListenAndServe(":8085", nil)
}