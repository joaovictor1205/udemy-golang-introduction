package main

import "runtime/debug"

func funct() {
	debug.PrintStack()
}

func secondfunc() {
	funct()
}

func firstfunc() {
	secondfunc()
}

func main() {

	firstfunc()

}
