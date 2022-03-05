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

func main() {
	a := 5
	b := 2
	c, d := sumandsem1(a, b)
	fmt.Println("main", c, d)
}
