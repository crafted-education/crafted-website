package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("web/static"))))
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
