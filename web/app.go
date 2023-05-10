package web

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("html/front.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hey this is a demo app")
	tpl.Execute(w, nil)
}
