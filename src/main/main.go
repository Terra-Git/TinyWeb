package main

import (
	// "web_server"
	"mysql_client"
)

func main() {
	// server := web_server.WebServer{}
	// server.Init()
	// server.Run(":9000")
	print("1")
	mysql_client.Test();
	print("2")
}