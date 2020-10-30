package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"library_pro/config"
	"log"
	"strings"
)

/*
封装一些基本的Mysql数据库操作方法
 */
//获取数据库实例
func GetDataBase() *sql.DB {
	//mysql 数据库
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{config.MYSQL_CONFIG["userName"], ":", config.MYSQL_CONFIG["password"], "@tcp(", config.MYSQL_CONFIG["ip"], ":", config.MYSQL_CONFIG["port"], ")/", config.MYSQL_CONFIG["dbName"], "?charset=utf8"}, "")
	fmt.Println(path)
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ := sql.Open("mysql", path)
	if DB == nil {
		log.Fatal("连接失败！")
		return nil
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(5)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Fatal("open database fail")
		return nil
	}
	return DB
}

//取一条或多条数据
func SelectSome(db *sql.DB, sql_str string, args ...interface{}) (*[]map[string]string, error) {
	stmt_out, err := db.Prepare(sql_str)
	if err != nil {
		panic(err.Error())
	}
	defer stmt_out.Close()

	rows, err := stmt_out.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := []map[string]string{}
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		v_map := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			v_map[columns[i]] = value
		}
		ret = append(ret, v_map)
	}
	return &ret, nil
}

//更新或插入，删除等
func Exec(db *sql.DB, sql_str string, values ...interface{}) (affect map[string]int64) {
	stmt, err := db.Prepare(sql_str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(values)
	res, err := stmt.Exec(values...)
	if err != nil {
		log.Fatal(err)
	}
	last_id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	row_affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	affect = map[string]int64{"last_id" : last_id, "row_affect": row_affect}

	return affect

}