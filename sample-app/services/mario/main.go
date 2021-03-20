package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "minTest.html")
}

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/code/", http.StripPrefix("/code/", http.FileServer(http.Dir("code"))))
	http.Handle("/Enjine/", http.StripPrefix("/Enjine/", http.FileServer(http.Dir("Enjine"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/midi/", http.StripPrefix("/midi/", http.FileServer(http.Dir("midi"))))
	http.Handle("/sounds/", http.StripPrefix("/sounds/", http.FileServer(http.Dir("sounds"))))
	http.HandleFunc("/", handler)
    http.ListenAndServe(":8086", nil)
}