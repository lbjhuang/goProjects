package main

import (
	"fmt"
	"strconv"
)

var name string
var age int
var salary float32

const country = "China"  //常量定义
const (n1 = 100
       n2    //如果没定义则跟上一个变量一致
)

func main() {
	name = "huangwei"
	age = 18
	salary = 18888.89
	fmt.Println(name)
	fmt.Printf("%T, %s",strconv.Itoa(age),strconv.Itoa(age))  //转字符串
	fmt.Printf("%T, %f",float64(age),float64(age))  //转浮点型
	fmt.Println(salary)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Printf("%s is %d, and he is from %s", name, age, country) //%s字符串占位符  %d整型占位符
}
