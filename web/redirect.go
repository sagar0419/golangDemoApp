package web

import (
	"html/template"
	"net/http"
)

var tplredirect = template.Must(template.ParseFiles("html/redirect.html"))

func Backend(w http.ResponseWriter, r *http.Request) {
	tplredirect.Execute(w, nil)
}
