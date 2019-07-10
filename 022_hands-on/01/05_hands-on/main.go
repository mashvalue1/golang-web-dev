package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/foo/", http.HandlerFunc(foo))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.gohtml"))
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("error raised.", err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("foo.gohtml"))
	err := tpl.ExecuteTemplate(w, "foo.gohtml", nil)
	if err != nil {
		log.Fatalln("error raised.", err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("me.gohtml"))
	err := tpl.ExecuteTemplate(w, "me.gohtml", "Ryuichi")
	if err != nil {
		log.Fatalln("error raised.", err)
	}
}
