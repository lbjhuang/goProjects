package main

import "fmt"

//多返回值，两种写法
func sumandsem(a, b int) (int, int) {
	sum := a + b
	sem := a - b
	return sum, sem
}

func sumandsem1(a, b int) (sum int, sem int) {
	sum = a + b
	sem = a - b
	return
}

//可变参数，会用数组接收
func sumandsem5(b ... int) (int)  {
	sum := 0
	//遍历数组计算相加
	for _, value := range b {
		sum += value
	}
	return sum
}

func main() {
	a := 5
	b := 2
	//多返回值 下划线忽略结果
	//c,d := sumandsem(a, b)
	c,_ := sumandsem1(a, b)
	fmt.Println("main", c)

	g:= sumandsem5(a, b,5,6)
	fmt.Println("main", g)
}
