package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		data := map[string]string{}
		//c.JSON(200, data)
		//data["name"] = c.Query("name") //获取参数值
		data["name"] = c.DefaultQuery("name","wade") //获取参数值，没有则指定默认值wade
		c.JSON(200, data)              //返回json
	})
	r.Run(":9000") // listen and serve on 0.0.0.0:8080

}
