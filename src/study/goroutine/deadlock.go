package main

import "fmt"
//流入无流出，死锁
func main() {
	ch1, ch2 :=make(chan int),make(chan int)    //ch1无缓冲，则死锁等待
	//ch1, ch2 :=make(chan int,2),make(chan int)    //ch1有缓冲，不发生死锁
	go func() {
		ch1 <-1
		ch2 <-0
	}()
    //<-ch1     //ch1管道无缓冲，注释了本行则ch1数据没有被其他go程读走，发生死锁
	var b =<-ch2
	fmt.Println(b)
}
