package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users/", UserHandler)
	http.HandleFunc("/usersAll/", userAll)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
