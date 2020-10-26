package main

import (
	"fmt"
	"time"
)

func DelayPrint() {
	for i := 1; i <= 4; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
	}
}

func HelloWorld() {
	fmt.Println("Hello world goroutine")
}

func main() {
	go DelayPrint()    // 开启第一个goroutine
	go HelloWorld()    // 开启第二个goroutine
	time.Sleep(2*time.Second)
	fmt.Println("main function")
}