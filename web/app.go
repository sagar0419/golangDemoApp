package web

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("html/front.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
