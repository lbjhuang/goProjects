package main

import (
	"fmt"
)
//通过通信共享内存而不是通过共享内存来通信，给内存加锁开销大切换频繁，非常好资源，channel管道则是一种更优秀的协程间通信方式


func main() {
	ch := make(chan int, 1) // 注意这里给的容量是1
	ch <- 1
	<-ch   //这里把管道的东西移走则下面的第一个case 有机会执行

	select {
	case ch <- 1:
		fmt.Println("ch可以插入数据")
	default:
		fmt.Println("通道channel已经满啦，塞不下东西了!")
	}
}
