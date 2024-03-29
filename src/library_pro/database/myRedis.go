package database

import (
	"github.com/gomodule/redigo/redis"
	"library_pro/config"
)


var RedisPool *redis.Pool  //创建redis连接池


func init(){
	RedisPool = &redis.Pool{     //实例化一个连接池
		MaxIdle:16,    //最初的连接数量
		MaxActive:0,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout:300,    //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn ,error){     //要连接的redis数据库
			c,error := redis.Dial("tcp",config.REDIS_URL)
			return c,error
		},
	}

}