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
	mux.HandleFunc("/api", web.Backend)
	mux.HandleFunc("/form", web.FormHandler)
	mux.HandleFunc("/upload", api.Submit)
	mux.HandleFunc("/cinema", api.Cinema)
	mux.HandleFunc("/record", api.MovieDetail)
	mux.HandleFunc("/metrics", api.Metrics)
	mux.HandleFunc("/list", api.Read)
	mux.HandleFunc("/files", api.Ufiles)
	mux.HandleFunc("/read", api.Rfiles)
	mux.HandleFunc("/", web.HomePage)
	fmt.Println("Server is listening on localhost port 3000")
	http.ListenAndServe(":3000", mux)
}
