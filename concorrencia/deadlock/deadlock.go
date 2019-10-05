package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1

	fmt.Println("After read")
}

func main() {

	c := make(chan int) //channel without buffer

	go routine(c)

	fmt.Println(<-c) //block
	fmt.Println("Read")
	fmt.Println(<-c) //deadlock

}
