package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"library_pro/database"
	"regexp"
	"strconv"
)

var my_redis redis.Conn

//接口限制频繁请求中间件
func NotOftenVisitUrl(c *gin.Context) {
	id := c.Param("id")
	url := c.Request.RequestURI
	reg, _ := regexp.Compile("\\/\\d");
	reg2 := reg.ReplaceAllString(url, "");
	//fmt.Printf("%T,%v",reg2,reg2)
	//路由键值及其失效时间
	var url_key_map = make(map[string]map[string]string)
	url_key_map["/employee/get"] = make(map[string]string)
	url_key_map["/employee/list"] = make(map[string]string)
	url_key_map["/employee/get"]["key_name"] = "GET_ONE_EMPLOYEE"
	url_key_map["/employee/get"]["expire"] = "5"
	url_key_map["/employee/list"]["key_name"] = "GET_EMPLOYEE_LIST"
	url_key_map["/employee/list"]["expire"] = "10"

	//每个请求设置的key
	//fmt.Printf(c.Request.RequestURI)
	key := url_key_map[reg2]["key_name"] + ":" + id
	fmt.Println(key)
	my_redis = database.RedisPool.Get()
	forbbiden_visit, err := redis.Bool(my_redis.Do("exists", key))
	if err != nil {
		fmt.Println("exists error : ", err)
		return
	}
	if forbbiden_visit {
		ttl_second,_ := redis.Int(my_redis.Do("TTL", key))
		ttl_string := strconv.Itoa(ttl_second)
		c.JSON(200, gin.H{
			"result": "请勿频繁请求,"+ttl_string+"秒之后再试，谢谢！",
		})
		c.Abort()
		return
	} else {
		//设置键的值和失效时间
		_, err := my_redis.Do("SET", key, id, "EX", url_key_map[reg2]["expire"])
		if err != nil {
			fmt.Println("set error : ", err.Error())
		}
		c.Next()
	}

}

//func NotOftenUrl(url string, c *gin.Context)  {
//t := time.Now()
//fmt.Println("我是自定义中间件第1种定义方式---请求之前")
////在gin上下文中定义一个变量
//c.Set("example", "CustomRouterMiddle1")
////请求之前
//c.Next()
//fmt.Println("我是自定义中间件第1种定义方式---请求之后")
////请求之后
////计算整个请求过程耗时
//t2 := time.Since(t)
//log.Println(t2)
//}
