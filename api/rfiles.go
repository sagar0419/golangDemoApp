package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func Ufiles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "After the URL add your pdf file name to read. for eg:- http://localhost:3000/read?file=xyz.pdf")
	fmt.Fprintln(w, "Below is the list of the uploaded files")

	list, err := exec.Command("ls", "-a", "uploads").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(list))
}

func Rfiles(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	fileName := url.Get("file")

	w.Header().Set("Content-Type", "application/pdf")
	resPDF, _ := ioutil.ReadFile("uploads/" + fileName)
	w.Write(resPDF)
	fmt.Println("PDF served successfully")
}
