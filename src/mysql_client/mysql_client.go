package mysql_client

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type MysqlUser struct{
	username_        string
	password_        string
	dp_ip_           string
	dp_port_         string
	dbname_          string
	max_connet       int
	max_free_connect int 
}
// 实现增删改查
type MysqlClient struct{
	connect_  *sql.DB
}

func (client *MysqlClient) Connet(username string,password string,ip string,port string,dbName string,max_connet int,max_free_connect int){
	client.connect_,err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, ip, port, dbName))
	if err != nil {
		return nil, err
	}
	// client.connect_,_ = sql.Open("mysql", "terra:kun123456@tcp(127.0.0.1:3306)/test")
	//设置数据库最大连接数
	client.connect_.SetConnMaxLifetime(max_connet)
	//设置上数据库最大闲置连接数
	client.connect_.SetMaxIdleConns(max_free_connect)
}

func (client *MysqlClient) Connect( user MysqlUser){
	client.connect_,err= sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",user.username_,user.password_,user.dp_ip_,user.dp_port_,user.dbname_))
	if err != nil {
		return nil, err
	}
	// client.connect_,_ = sql.Open("mysql", "terra:kun123456@tcp(127.0.0.1:3306)/test")
	//设置数据库最大连接数
	client.connect_.SetConnMaxLifetime(max_connet)
	//设置上数据库最大闲置连接数
	client.connect_.SetMaxIdleConns(max_free_connect)
}

func (client *MysqlClient) query(){
    
}

func (client *MysqlClient) Test(){

	//验证连接
	if err := client.connect_.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}