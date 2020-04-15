package connection

import (
	"database/sql"
	"fmt"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

var DB *sql.DB

func InitDB() {
	var err error
	loc := url.QueryEscape("Asia/Shanghai")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		configs.DefaultConfig.Database.User, configs.DefaultConfig.Database.Password,
		configs.DefaultConfig.Database.Host, configs.DefaultConfig.Database.Name, loc)
	DB, err = sql.Open("mysql", dsn, )
	if err != nil {
		panic(err)
	}
}
