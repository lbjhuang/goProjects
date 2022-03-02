package main

import "fmt"

//定义一个接口，里面提供一个方法，约束后来者实现这个方法，接口只有方法没有方法体，不能有值
type Speaker interface {
	speak()
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

type human struct {
	name string
}

func (c Cat) speak() {        //三个结构体都实现了接口的speak方法
	fmt.Println(c.name,"is 喵喵")
}

func (d Dog) speak() {
	fmt.Println(d.name,"is 旺旺")
}

func (h human) speak() {
	fmt.Println(h.name,"is 喔喔")
}

func da(x Speaker)  {
	x.speak()
}

func main() {
	var c1 = Cat{name:"lucky"}
	var d1 = Dog{name:"lamb"}
	var h1 = human{name:"jiagu"}

	da(c1)
	da(d1)
	da(h1)
}

