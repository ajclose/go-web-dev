package main

import (
	"io"
	"net/http"
)

type index int

func (i index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the root page")
}

type dog int

func (d dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "THIS IS DOG")
}

type me int

func (m me) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "My name is Alex")
}

func main() {
	var i index
	var d dog
	var m me

	http.Handle("/", i)
	http.Handle("/dog/", d)
	http.Handle("/me/", m)

	http.ListenAndServe(":8080", nil)
}
