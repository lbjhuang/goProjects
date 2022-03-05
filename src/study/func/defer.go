package main

import "fmt"

func test()  {
	fmt.Println("aa")
	fmt.Println("bb")
	defer fmt.Println("this is defer1")
	defer fmt.Println("this is defer2")
	defer fmt.Println("this is defer3")
}

func defer1(){
	for i :=0;i<10;i++{
		defer fmt.Println("this is defer ",i)
	}
}

func defer2(){
	var i = 100
	defer fmt.Println("defer i =",i)  //定义defer 的时候已经传入初始值，不会随着程序的变化而变化 defer i = 100
	i = 120
	fmt.Println("i =", i)
}


func main()  {
	test()
	defer1()
	defer2()
	//执行顺序逆序 出栈的行为类似
	//aa
	//bb
	//this is defer3
	//this is defer2
	//this is defer1
}