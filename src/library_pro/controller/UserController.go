package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"library_pro/database"
	_ "library_pro/model"
	"log"
	"net/http"
	"time"
	"library_pro/model"
)

var db *sql.DB

func init() {
	log.Println(">>>> get database connection start <<<<")
	log.Println(time.Now().Format("2006-01-02 15:04:05")) //奇葩的是必须是这个时间点:2006-05-02 18:04:05, 是go的诞生日，否则无法格式化这个效果
	db = database.GetDataBase()
}

func InsertOne(context *gin.Context) {
	//接受参数
	var employee model.Employee
	if context.BindJSON(&employee) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "参数格式有误"})
		return
	}
	fmt.Println(employee)

	//插入数据
	insert_sql := "insert into erp_employee_info (chinese_name,english_name,position_name,birthday,create_time) values (?,?,?,?,?);"
	affect_map := database.Exec(db, insert_sql, employee.Chinese_name, employee.English_name, employee.Position_name, employee.Birthday, time.Now().Format("2006-01-02 15:04:05"))
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

func UpdateOne(context *gin.Context) {
	var employee model.Employee
	if context.BindJSON(&employee) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "参数格式有误"})
		return
	}

	//更新数据
	insert_sql := "update erp_employee_info set chinese_name = ?, english_name = ?, position_name = ?, status = ?, birthday = ? where id = ?"
	affect_map := database.Exec(db, insert_sql, employee.Chinese_name, employee.English_name, employee.Position_name, employee.Status, employee.Birthday,employee.Id)
	context.JSON(200, gin.H{
		"result": affect_map,
	})
}

func GetSalary(context *gin.Context) {
	// 获取请求参数
	// 查询数据库
	select_sql := "select a.id,a.chinese_name,b.salary from erp_employee_info a left join erp_employee_salary b on a.id = b.user_id where a.id = 1";
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
