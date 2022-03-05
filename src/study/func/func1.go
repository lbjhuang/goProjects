package main

import "fmt"

func changeNum(num int) {
	num = 10000
	fmt.Println("changNum", num, &num)
}

func changeNumP(num *int, fac int)  {
	 *num = *num * fac
}

func main()  {
	var num = 100
	changeNum(num)
	changeNumP(&num,22)
	fmt.Println("main", num,&num)
}
