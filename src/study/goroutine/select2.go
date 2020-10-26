package main

import (
	"fmt"
	"time"
)

var echo chan string
var receive chan string

func Echo(out chan<- string) {
	time.Sleep(1 * time.Second)
	echo <- "红色石头"
}

func Receive(out chan<- string, in <-chan string) {
	temp := <-in //阻塞等待echo返回
	out <- temp
}

func main() {
	echo = make(chan string)
	receive = make(chan string)
	go Echo(echo)
	go Receive(receive, echo)
	getStr := <-receive //接收goroutine 2的返回
	fmt.Println(getStr)

}
