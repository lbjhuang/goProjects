package main

import "fmt"

type Person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	var p Person
	p.name = "james"
	p.age = 18
	p.hobby = []string{"basketball"}
	fmt.Println(p.name) //打印结构体属性

	fmt.Printf("%p \n", &p.name)
	fmt.Printf("%p \n", &p.age)
	fmt.Printf("%p \n", &p.hobby)
}
