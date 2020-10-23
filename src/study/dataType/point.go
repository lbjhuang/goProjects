package main

import "fmt"

func main() {
	var c *int   //声明一个整型指针c，此时内存地址是nil， *c = 18  的写法是不对的，空指针没地址

	fmt.Println(c)  //打印c -> nil

	a := 123
	d := 108
	b :=&a   //类型推导直接声明一个指针b
	fmt.Println(b)  //打印a的内存地址
	m := *b
	fmt.Println(m)   //123

	*b++
	fmt.Println(a)

	c = &d
	fmt.Println(*c)   //108
	fmt.Println(c)

}