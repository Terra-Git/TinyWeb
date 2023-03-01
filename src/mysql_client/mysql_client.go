package mysql_client

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Test(){
	DB, _ := sql.Open("mysql", "terra:password@tcp(127.0.0.1:3306)/test")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}