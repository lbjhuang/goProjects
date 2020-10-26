package main
//流出无流入，发生死锁
func main() {
	ch := make(chan int)
	<- ch // 阻塞main goroutine, 通道被锁
}