package mysql_client

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlUser struct{
	Username_          string
	Password_          string
	Dp_ip_             string
	Dp_port_           string
	Dbname_            string
	Max_connect_       int 
}

// 实现增删改查
type MysqlClient struct{
	connect_     *sql.DB
	mysql_user_  MysqlUser
}

func (client *MysqlClient) Connet(username string,password string,ip string,port string,dbName string,max_connect int){
	client.connect_,_ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, ip, port, dbName))
	//设置上数据库最大闲置连接数
	client.connect_.SetMaxIdleConns(max_connect)
}

func (client *MysqlClient) Connect( user MysqlUser){
	client.mysql_user_ = user
	client.connect_,_= sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",user.Username_,user.Password_,user.Dp_ip_,user.Dp_port_,user.Dbname_))
	//设置数据库最大连接数
	client.connect_.SetMaxIdleConns(user.Max_connect_)
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