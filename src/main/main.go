package main

import (
	// "web_server"
	"mysql_client"
)

func main() {
	client  := mysql_client.MysqlClient{}
	client.Connet()
	client.Test()
}