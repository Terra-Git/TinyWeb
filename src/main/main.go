package main

import (
	. "mysql_client"
)

func main() {
	client := MysqlClient{}
	mysql_user := MysqlConnectIfo{}
	
    mysql_user.Username_         = "terra"
	mysql_user.Password_         = "kun123456"
	mysql_user.Dp_ip_            = "127.0.0.1"
	mysql_user.Dp_port_          = "3306"
	mysql_user.Dbname_           = "test"
	mysql_user.Max_connect_       = 10

	client.Connect(mysql_user)

	client.Test_insert()

	test1 := MysqlColData{"name","gyk",MYSQL_STRING}
	test2 := MysqlColData{"phone","123",MYSQL_INT}
	test3 := MysqlColData{"address","hunan",MYSQL_STRING}

	var data []MysqlColData

	data = append(data , test1)
	data = append(data , test2)
	data = append(data , test3)

	var datas [][]MysqlColData

	datas = append(datas,data)
	datas = append(datas,data)

	client.Insert_record("t_note_user",data)
}