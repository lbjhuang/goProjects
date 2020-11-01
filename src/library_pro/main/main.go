package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/gomodule/redigo/redis"
	"io"
	"library_pro/controller"
	"os"
	"library_pro/config"
	"library_pro/middlerware"
)


func main() {
	fmt.Println("运行开始...")
	f,err :=os.Create(config.RUN_LOG_PATH)
	if err != nil{
		fmt.Println(err.Error())
	}


	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	// Engine
	router := gin.Default()
	router.Use(gin.Logger(),gin.Recovery())
	// 路由组
	employee := router.Group("/employee")
	{ // 请求参数在请求路径上
		employee.GET("/get/:id", middlerware.NotOftenVisitUrl, controller.GetById)
		employee.GET("/list", middlerware.NotOftenVisitUrl, controller.GetEmployeeList)

		employee.POST("/update", controller.UpdateOne)
		employee.POST("/insert", controller.InsertOne)
		employee.GET("/salary", controller.GetSalary)
		//可以自己添加其他，一个请求的路径对应一个函数

	}

	// 指定地址和端口号
	router.Run(":9090")
}
