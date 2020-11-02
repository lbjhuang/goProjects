package config

//日志路径
const RUN_LOG_PATH = "../logs/library_run_log.log"

//数据库的配置信息
var MYSQL_CONFIG = map[string]string{"userName":"root", "password":"root", "ip":"localhost", "port": "3306", "dbName":"erp_database"}
var REDIS_URL = "localhost:6379"
var ES_CLUSTER_HOST_ONE   = "http://192.168.109.53:9200"
var ES_CLUSTER_HOST_TWO   = "http://192.168.109.54:9200"
var ES_CLUSTER_HOST_THREE = "http://192.168.109.55:9200"
