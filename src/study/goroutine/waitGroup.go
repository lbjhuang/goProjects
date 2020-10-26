package main

import (
	"fmt"
	"sync"
	"time"
)
//waitGroup 的计数器机制能保证所有协程都执行完毕
var wg sync.WaitGroup
func printOne()  {
	for i :=0; i<10; i++ {
		fmt.Println("print1: ",i)
		time.Sleep(1*time.Second)
	}
	wg.Done()   //计数器减1
	fmt.Println("print1计数器减1")
}

func printTwo()  {
	for i :=0; i<10; i++ {
		fmt.Println("print2: ",i)
		time.Sleep(1*time.Second)
	}
	wg.Done()   //计数器减1
	fmt.Println("print2计数器减1")

}

func main() {
	fmt.Println("计数器加2")
	wg.Add(2)   //计数器加2，可以用于2个协程的计数
	go printOne()
	go printTwo()
	wg.Wait()   //等待计数器为0，执行后面的
	fmt.Println("等待所有协程执行后，主线程退出")
}
