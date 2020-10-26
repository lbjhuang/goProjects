package main

import "fmt"

//管道是引用类型，先进先出
func main() {
	var ch = make(chan int, 5)
	var ch2 = make(chan int, 5)

	ch <- 5
	ch <- 6
	ch <- 8
	fmt.Printf("值：%v，长度%v，容量%v \n", ch, len(ch), cap(ch)) //值：0xc00007c000 长度3，容量5   //值是一个地址
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//引用类型，ch ch2同一片内存
	ch2 = ch

	ch2 <-99
	fmt.Println(<-ch)   //99
	fmt.Printf("值：%v，长度%v，容量%v \n", ch2, len(ch2), cap(ch2)) //值：0xc00007c000，长度0，容量5


}
