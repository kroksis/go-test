package main

import (
	"bufio"
	// "html/template"
	"net/http"
	"os"
)

// var tpl *template.Template

// func init() {
// 	var err error
// 	tpl, err = template.ParseFiles("templates/page.gohtml")
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	http.HandleFunc("/", index)

	fs := http.FileServer(http.Dir("public"))

	http.Handle("/css/", fs)
	http.Handle("/img/", fs)
	http.Handle("/js/", fs)
	http.Handle("/html/", fs)

	http.ListenAndServe(":80", nil)
}

// func indexTemplate(w http.ResponseWriter, r *http.Request) {
// 	tpl.Execute(w, nil)
// }

func index(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("templates/page.html")
	if err == nil {
		bufferedReader := bufio.NewReader(f)
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
