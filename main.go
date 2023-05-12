package main

import (
	"fmt"
	"net/http"

	api "demo/github.com/api"
	db "demo/github.com/pkg/db"
	web "demo/github.com/web"
)

func main() {
	db.CreateDb()
	db.CreateTable()

	mux := http.NewServeMux()
	mux.HandleFunc("/", web.HomePage)
	mux.HandleFunc("/api", web.Backend)
	mux.HandleFunc("/cinema", api.Cinema)
	mux.HandleFunc("/record", api.MovieDetail)

	fmt.Println("Server is listening on localhost port 3000")
	http.ListenAndServe(":3000", mux)
}
