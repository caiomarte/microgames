package main

import (
	"net/http"
)

func Catalog(w http.ResponseWriter, r *http.Request) {
	d := Data{
		Title:  "MicroGames",
		Player: "CaioMarte",
		Score:  "1000pt",
	}

	Render(w, "catalog.html", d)
}
