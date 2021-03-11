package main

import (
	"net/http"
)

func Catalog(w http.ResponseWriter, r *http.Request) {
	d := Data{
		Title: "Catalog",
	}
	Render(w, "catalog.html", d)
}
