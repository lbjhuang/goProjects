package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 28

	select {
	case <-ch:    //ch管道内有数据，可以从ch里面读取到，所以走这里的分支
		fmt.Println("ch1 pop one element")
	case <-ch:
		fmt.Println("ch2 pop one element")
	default:
		fmt.Println("default")
	}

}
