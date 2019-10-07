package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("SERVER ON")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
