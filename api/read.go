package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Read(w http.ResponseWriter, r *http.Request) {
	list, err := os.ReadDir("./uploads")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range list {
		fmt.Fprintln(w, e.Name())
	}
}
