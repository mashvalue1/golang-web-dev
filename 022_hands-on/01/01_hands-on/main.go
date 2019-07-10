package main

import (
	"io"
	"net/http"
)

// funcはmainのあとに書いたほうがいい？

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "index")
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "my name")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/foo/", foo)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
