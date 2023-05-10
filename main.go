package main

import (
	"fmt"
	"net/http"

	web "alteryx/github.com/web"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", web.HomePage)
	mux.HandleFunc("/api", web.Backend)
	fmt.Println("Server is listening on localhost port 3000")

	http.ListenAndServe(":3000", mux)
}
