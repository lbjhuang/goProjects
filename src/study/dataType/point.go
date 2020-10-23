package main

import "fmt"

func main() {
	var c *int   //声明一个整型指针c


	a := 123
	d := 108
	b :=&a   //类型推导直接声明一个指针b
	fmt.Println(b)
	m := *b
	fmt.Println(m)   //123

	*b++
	fmt.Println(a)

	c = &d
	fmt.Println(*c)   //108
	fmt.Println(c)

}