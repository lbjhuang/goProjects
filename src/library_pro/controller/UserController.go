package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"library_pro/database"
	"library_pro/model"
	_ "library_pro/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

var db *sql.DB

func init() {
	db = database.GetDataBase()
	log.Println(">>>> get database connection start <<<<")
}

//插入数据到mysql 和 es
func InsertOne(context *gin.Context) {
	//接受参数
	var employee model.Employee
	if err := context.ShouldBind(&employee); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": 500, "error": err.Error()})
		return
	}
	//参数校验
	if employee.Chinese_name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"status": 500, "error": "请传入中文名"})
		return
	}
	if employee.English_name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"status": 500, "error": "请传入英文名"})
		return
	}

	//插入数据到Mysql
	insert_sql := "insert into erp_employee_info (chinese_name,english_name,position_name,birthday,create_time) values (?,?,?,?,?);"
	affect_map := database.Exec(db, insert_sql, employee.Chinese_name, employee.English_name, employee.Position_name, employee.Birthday, time.Now().Format("2006-01-02 15:04:05"))
	//插入数据到ES
	jsonMap := make(map[string]interface{})
	jsonMap["chinese_name"] = employee.Chinese_name
	jsonMap["english_name"] = employee.English_name
	jsonMap["position_name"] = employee.Position_name
	jsonMap["birthday"] = employee.Birthday
	jsonMap["status"] = "1"
	jsonMap["create_time"] = time.Now().Format("2006-01-02 15:04:05")
	client := database.ConnectES();
	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}

	fmt.Println(jsonData) //打印json数据 []bytes 类型
	index := "erp_employee_info"
	id := affect_map["last_id"]
	esId := database.CreateESDoc(client, string(jsonData), index, strconv.FormatInt(id, 10))
	es_insert_id, _ := strconv.ParseInt(esId, 10, 64)
	affect_map["es_insert_id"] = es_insert_id

	//返回处理结果
	context.JSON(200, gin.H{
		"result": affect_map,
	})
}

func GetById(context *gin.Context) {

	println(">>>> get user by id and name action start <<<<")
	// 获取请求参数
	id := context.Param("id")
	sendTopicMessage()
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
	if err := context.ShouldBind(&employee); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": 500, "error": err.Error()})
		return
	}
	if employee.Id == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"status": 500, "error": "请传入ID"})
	}

	//更新数据
	insert_sql := "update erp_employee_info set chinese_name = ?, english_name = ?, position_name = ?, status = ?, birthday = ? where id = ?"
	affect_map := database.Exec(db, insert_sql, employee.Chinese_name, employee.English_name, employee.Position_name, employee.Status, employee.Birthday, employee.Id)
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

func TestESSearch(ginContext *gin.Context) {
	var res *elastic.SearchResult
	var err error
	client := database.ConnectES();
	mustMatchQuery := elastic.NewBoolQuery()

	//高亮显示字段
	hl := elastic.NewHighlight()
	hl = hl.Fields(elastic.NewHighlighterField("goodsname"))
	hl.HighlightFilter(true)
	hl.RequireFieldMatch(true)
	//高亮显示格式
	hl = hl.PreTags("<span style='color:red'>").PostTags("</span>")

	mustMatchQuery.Must(elastic.NewMatchQuery("goodsname", "无线蓝牙手机直连"))
	res, err = client.Search("erp_lazada_item_list").Query(mustMatchQuery).Highlight(hl).Size(150).Pretty(true).Do(context.Background())
	checkError(err)
	ginContext.JSON(200, gin.H{
		"result": res,
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
