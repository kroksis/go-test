package main

//Just added a comment

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("templates/index.gohtml")
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", index)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/css/", fs)
	http.Handle("/img/", fs)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
