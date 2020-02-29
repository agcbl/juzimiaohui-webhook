package connection

import (
	"database/sql"
	"fmt"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",
		configs.DefaultConfig.Database.User, configs.DefaultConfig.Database.Password,
		configs.DefaultConfig.Database.Host, configs.DefaultConfig.Database.Name)
	fmt.Printf("%s\n", dsn)
	DB, err = sql.Open("mysql", dsn, )
	if err != nil {
		panic(err)
	}
}
