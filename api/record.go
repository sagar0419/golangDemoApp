package api

import (
	db "demo/github.com/pkg/db"
	"fmt"
	"net/http"
)

func MovieDetail(w http.ResponseWriter, r *http.Request) {
	list, _ := db.QueryDb()
	for _, details := range list {
		BookInfo := fmt.Sprintf("Movie name is %v and Hero is %v \n", details.Movie, details.Hero)
		fmt.Fprintln(w, BookInfo)
	}
}
