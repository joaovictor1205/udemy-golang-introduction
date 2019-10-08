package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func correctHour(w http.ResponseWriter, r *http.Request) {

	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1> Correct Hour: %s</h1>", s)

}

func main() {

	http.HandleFunc("/", correctHour)

	log.Println("SERVER ON")

	log.Fatal(http.ListenAndServe(":8000", nil))

}
