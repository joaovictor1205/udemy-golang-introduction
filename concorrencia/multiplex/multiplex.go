package main

import (
	"fmt"

	"github.com/joaovictor1205/html"
)

func sending(origin <-chan string, destiny chan string) {

	for {
		destiny <- <-origin
	}
}

func multiplex(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go sending(input1, c)
	go sending(input2, c)

	return c
}

func main() {

	c := multiplex(
		html.Title("https://www.google.com.br", "https://www.google.com.br"),
		html.Title("https://www.facebook.com.br", "https://www.youtube.com.br"),
	)

	fmt.Println(<-c, <-c)

}
