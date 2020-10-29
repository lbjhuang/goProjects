package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"library_pro/database"
	_ "library_pro/model"
	"log"
	"net/http"
	"time"
)

var db *sql.DB

func init() {
	log.Println(">>>> get database connection start <<<<")
	log.Println(time.Now().Format("2006-01-02 15:04:05")) //奇葩的是必须是这个时间点:2006-05-02 18:04:05, 是go的诞生日，否则无法格式化这个效果
	db = database.GetDataBase()
}

func InsertOne(context *gin.Context) {
	//定义一个结构体接受参数
	type Employee struct {
		Chinese_name  string `form:"chinese_name" json:"chinese_name" binding:"required"`
		English_name  string `form:"english_name" json:"english_name" binding:"required"`
		Position_name string `form:"position_name" json:"position_name" binding:"required"`
		Birthday      string `form:"birthday" json:"birthday" binding:"required"`
	}
	var employee Employee
	if context.BindJSON(&employee) != nil {
		context.JSON(http.StatusOK, gin.H{"status": "参数格式有误"})
	}

	//更新数据
	insert_sql := "insert into erp_employee_info (chinese_name,english_name,position_name,birthday,create_time) values (?,?,?,?,?)"
	affect_map := database.Exec(db, insert_sql, "黄的伟","james","php","1991-08-27","2020-09-08 15:15:18")
	context.JSON(200, gin.H{
		"result": affect_map,
	})
}

func GetById(context *gin.Context) {
	println(">>>> get user by id and name action start <<<<")
	// 获取请求参数
	id := context.Param("id")
	// 查询数据库
	data_map, err := database.SelectSome(db, "select * from  erp_employee_info where id = ?", id)
	checkError(err)
	context.JSON(200, gin.H{
		"result": data_map,
	})
}

func UpdateById(context *gin.Context) {
	// 获取请求参数
	id := context.Param("id")
	chines_name := context.Param("chines_name")
	english_name := context.Param("english_name")
	position_name := context.Param("position_name")
	status := context.Param("status")
	//更新数据
	update_sql := "update erp_employee_info set chines_name = ?,english_name=?,position_name=?,status=? where id = ?"
	affect_map := database.Exec(db, update_sql, chines_name, english_name, position_name, status, id)
	context.JSON(200, gin.H{
		"result": affect_map,
	})
}

func GetEmployeeList(context *gin.Context) {
	println(">>>> get user by id and name action start <<<<")
	// 获取请求参数
	// 查询数据库
	select_sql := "select * from erp_employee_info";
	data_map, err := database.SelectSome(db, select_sql)
	checkError(err)
	context.JSON(200, gin.H{
		"result": data_map,
	})
}

// 跳转html
func RenderForm(context *gin.Context) {
	println(">>>> render to html action start <<<<")

	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200, "insertUser.html", gin.H{})
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
