package config

//日志路径
const RUN_LOG_PATH = "../logs/library_run_log.log"

//数据库的配置信息
var MYSQL_CONFIG = map[string]string{"userName":"root", "password":"root", "ip":"localhost", "port": "3306", "dbName":"erp_database"}
var REDIS_URL = "localhost:6379"