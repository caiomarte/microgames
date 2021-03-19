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

	g := Games{
		Flappy: "localhost:8081",
		Blocks: "localhost:8082",
		G2048:  "localhost:8083",
	}

	Render(w, "catalog.html", d, g)
}
