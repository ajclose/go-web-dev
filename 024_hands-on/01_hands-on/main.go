package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/dog.jpg", dogPic)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		println(err)
	}
	io.WriteString(w, `<h1>This is from dog</h1>`)
	io.WriteString(w, `<img src="dog.jpg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}
