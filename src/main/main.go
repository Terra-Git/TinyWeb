package main

import (
	"mysql_client"
)


func main() {
	client := mysql_client.MysqlClient{}
	mysql_user := mysql_client.MysqlConnectIfo{}
	
    mysql_user.Username_         = "terra"
	mysql_user.Password_         = "kun123456"
	mysql_user.Dp_ip_            = "127.0.0.1"
	mysql_user.Dp_port_          = "3306"
	mysql_user.Dbname_           = "test"
	mysql_user.Max_connect_       = 10

	client.Connect(mysql_user)

	client.Test()
}