package mysql_client

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlClient struct{
	connect_  *sql.DB
}

func (client *MysqlClient) Connet(){
	client.connect_,_ = sql.Open("mysql", "terra:kun123456@tcp(127.0.0.1:3306)/test")
	//设置数据库最大连接数
	client.connect_.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	client.connect_.SetMaxIdleConns(10)

}

func (client *MysqlClient) Test(){

	//验证连接
	if err := client.connect_.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}