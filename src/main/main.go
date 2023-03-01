package main

import "web_server"

func main() {
	server := web_server.WebServer{}
	server.Init()
	server.Run(":9000")
}