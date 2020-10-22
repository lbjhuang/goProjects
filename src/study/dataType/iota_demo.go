package main

import "fmt"

const (
	n3 = iota    //iota   从0开始的枚举排序  0 1 2
	n4
	n5
	n6 = 102   //iota在这里中断
	n7   //102
)

func main() {
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)

}
