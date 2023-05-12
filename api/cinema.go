package api

import (
	db "demo/github.com/pkg/db"
	"fmt"
	"net/http"
)

var films = make(map[string]string)

func Cinema(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	movieName := url.Get("movie")
	heroName := url.Get("hero")
	result := db.CheckDb(movieName)
	if result != "" {
		fmt.Fprintln(w, "Movie details are already exsist in the database")
		return
	}
	db.InsertMovie(movieName, heroName)
	fmt.Fprintln(w, "Your Movie details are saved")
}
