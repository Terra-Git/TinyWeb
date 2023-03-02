package main

import (
	"mysql_client"
)

func main() {
	client  := mysql_client.MysqlClient{}
	client.Connet()
	client.Test()
}