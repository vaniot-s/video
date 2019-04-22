package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	dbConn *sql.DB
	err    error
)

//连接数据库,内部方法 root:123!@#@tcp(localhost:3306)/video_server?charset=utf8
func init() {
	dbConn, err = sql.Open(os.Getenv("DRIVENAME"), os.Getenv("USERNAME")+":"+os.Getenv("PASSWORD")+"@tcp("+os.Getenv("DBHOST")+":"+os.Getenv("PORT")+")/"+os.Getenv("DATABASE")+"?charset=utf8") //tcp
	if err != nil {
		panic(err.Error()) //
	}
}
