package main

import (
	"fmt"
	"time"
)

func mul(base int, c chan int) {

	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base

}

func main() {

	ch := make(chan int)
	fmt.Println("A")

	go mul(2, ch)

	a, b := <-ch, <-ch
	fmt.Println("B")

	fmt.Println(a, b)
	fmt.Println(<-ch)

}
