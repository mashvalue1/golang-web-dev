package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./starting-files/public/"))))
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
