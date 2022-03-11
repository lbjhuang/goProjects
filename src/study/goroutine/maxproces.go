package main

import (
	"fmt"
	"runtime"
	"time"
)

//设置程序执行的CPU核数设置
var i int

//死循环执行
func cal()  {
	for {
		i++
	}
}

func main()  {
	cpu := runtime.NumCPU()  //本电脑8核
	fmt.Println(cpu)
	//runtime.GOMAXPROCS(1)  //设置1核
	//runtime.GOMAXPROCS(2)  //设置2核
	//runtime.GOMAXPROCS(5)  //设置5核
	//runtime.GOMAXPROCS(8)  //设置8或者大于8的核数，则CPU会占满
	for i:=0;i<10;i++{
		go cal()
	}
	//设置睡眠时间
	time.Sleep(time.Hour)
}