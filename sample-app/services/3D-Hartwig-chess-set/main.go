package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/mckps/", http.StripPrefix("/mckps/", http.FileServer(http.Dir("mckps"))))
	http.HandleFunc("/", handler)
    http.ListenAndServe(":8085", nil)
}