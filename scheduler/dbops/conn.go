package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

//连接数据库,内部方法 root:123!@#@tcp(localhost:3306)/video_server?charset=utf8
func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/video?charset=utf8") //tcp
	if err != nil {
		panic(err.Error()) //
	}
}
