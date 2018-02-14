package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("pic"))))
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
