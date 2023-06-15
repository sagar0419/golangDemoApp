package web

import (
	"html/template"
	"net/http"
)

var tplform = template.Must(template.ParseFiles("html/form.html"))

func FormHandler(w http.ResponseWriter, r *http.Request) {
	tplform.Execute(w, nil)
}
