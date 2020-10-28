package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

type UserInfo struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`            //User  结构体首字母一定要大些
	Password string `form:"password" json:"password" xml:"password" binding:"required"` //Password  结构体首字母一定要大些
}
var my_db *sql.DB
type ERPPerson struct {
	Id         int
	ChineseName string
	EnglishName  string
	PostName  string
	IsON  string
}


func main() {
	fmt.Println(78912)
	db, err := sql.Open("mysql", "root:!QAZxsw2@tcp(192.168.109.21:3306)/erp_database")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	r := gin.Default()
	//r.Use(loginCheck());   //这里注册全局的中间件
	r.GET("/person/:id", func(c *gin.Context) {
		var (
			person ERPPerson
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id,chineseName,englishName,postName from erp_employeeInfo where id = ?;", id)
		err = row.Scan(&person.Id, &person.ChineseName, &person.EnglishName, &person.PostName)
		if err != nil {
			// If no results send null
			result = gin.H{
				"error_result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})



	r.GET("/persons", func(c *gin.Context) {
		var (
			person  ERPPerson
			persons []ERPPerson
		)
		rows, err := db.Query("select id,chineseName,englishName,postName from erp_employeeInfo where staffState < ? order by id desc ;",3)
		if err != nil {
			fmt.Print("mysql 报错了：",err.Error())
		}

		for rows.Next() {
			err = rows.Scan(&person.Id, &person.ChineseName, &person.EnglishName, &person.PostName)
			persons = append(persons, person)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		var is_ok = map[int]string{1611:"nice",1006:"nice"}
		//循环里面进行赋值
		for key, val := range persons {
			persons[key].IsON = is_ok[val.Id]
		}

		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})

	r.GET("/persons", func(c *gin.Context) {
		var (
			person  ERPPerson
			persons []ERPPerson
		)
		rows, err := db.Query("select id,chineseName,englishName,postName from erp_employeeInfo where staffState < ? order by id desc ;",3)
		if err != nil {
			fmt.Print("mysql 报错了：",err.Error())
		}

		for rows.Next() {
			err = rows.Scan(&person.Id, &person.ChineseName, &person.EnglishName, &person.PostName)
			persons = append(persons, person)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		var is_ok = map[int]string{1611:"nice",1006:"nice"}
		//循环里面进行赋值
		for key, val := range persons {
			persons[key].IsON = is_ok[val.Id]
		}

		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})












	r.GET("/ping", FirstMiddleWare(), func(c *gin.Context) {   //某个路由单独局部中间件
		data := map[string]string{}
		//c.JSON(200, data)
		//data["name"] = c.Query("name") //获取参数值
		//data["name"] = c.DefaultQuery("name", "wade") //获取参数值，没有则指定默认值wade
		data["name"] = c.DefaultQuery("name", "wade") //获取参数值，没有则指定默认值wade
		c.JSON(200, data)                             //返回json
	})

	r.POST("/login", func(c *gin.Context) {
		var user UserInfo
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
			return
		}

		if user.User != "james" || user.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are login in"})
	})

	//内部重定向
	r.GET("/logout", func(c *gin.Context) {
		c.Request.URL.Path = "/index"
		r.HandleContext(c)
	})

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "this is index page "})
	})

	//外部重定向
	r.GET("/2345site", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.2345.com")
	})

	r.Run(":9000") // listen and serve on 0.0.0.0:8080
}

//中间件
func FirstMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") //可以通过c.Set设置一个值放在请求上下文中，后续的函数可用到
		c.Next()
		time.Sleep(1 * time.Second)
		//c.Abort()
		cost := time.Since(start) //睡眠1s
		log.Println("cost time is ", cost)
	}
}


func loginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//redis取token，判断是否token过期，过期则跳转到登录页
	}
}