package main

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int, 8)
	stringCh := make(chan string, 6)

	for i := 0; i < 8; i++ {
		intCh <- i
	}

	for i := 0; i < 5; i++ {
		stringCh <- fmt.Sprintf("%v", i) + " is insert into stringCh"
	}

	for {
		select {
		case v := <-intCh: //ch管道内有数据，可以从ch里面读取到，所以走这里的分支
			fmt.Printf("intCh pop one element %v \n",v)
			time.Sleep(time.Microsecond*50)
		case v :=<-stringCh:
			fmt.Printf("intCh pop one element %v \n",v)
			time.Sleep(time.Microsecond*50)
		default:
			fmt.Println("default")
			return
		}
	}
}
