package main

import (
	"github.com/gin-gonic/gin"
	"library_pro/controller"
)

func main() {

	// Engin
	router := gin.Default()
	//router := gin.New()

	// 路由组
	employee := router.Group("/employee")
	{ // 请求参数在请求路径上
		employee.GET("/get/:id/", controller.GetById)
		employee.GET("/list", controller.GetEmployeeList)
		//employee.POST("/update", controller.UpdateById)
		employee.POST("/insert", controller.InsertOne)
		//可以自己添加其他，一个请求的路径对应一个函数

	}

	// 指定地址和端口号
	router.Run(":9090")
}
