package main

import "fmt"

var ch = make(chan int)

func hello() {
	fmt.Println("hello goroutine")
	ch <- 2
}

func main() {

	go hello()
	var a = <-ch
	fmt.Println("this is main func a", a)
}
