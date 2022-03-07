package main

import (
	"math"
	"fmt"
)

//定义几何接口，抽象2个方法: 面积 和 周长
type geometry interface {
	area() float64
	girth() float64
}

//定义2个结构体 长方形和圆形 分别取实现几何接口的方法
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//两个结构体 分别实现接口的面积和周长计算方法
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) girth() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) girth() float64 {
	return 2 * math.Pi * c.radius
}

//调用两个计算方法
func measure(g geometry)  {
	fmt.Println(g.area())
	fmt.Println(g.girth())

}

func main() {
	//初始化两个结构体，分别输出面积周长计算结果
	r := rect{6, 6}
	c := circle{8}

	measure(r)
	measure(c)
}
