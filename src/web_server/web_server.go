package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "html/template"
)

// http方法结构体
type HttpMethodNode struct{
	http_method_       string
	http_url_          string
	http_method_func_  func(c *gin.Context)
}

// 服务结构体
type WebServer struct{
	engine_   *gin.Engine
	methods_   []HttpMethodNode
}

// ping 回调
func ping_func(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 登录回调
func login_func(c *gin.Context) {
	// 当使用DefaultQuery时，如果没有获取到浏览器输入的username，则返回设置defaultValue给username
	username := c.Query("username")
	if( username == ""){
		c.String(http.StatusOK, "invilid param username")
		return 
	}
	// 当使用Query时，如果没有获取到浏览器输入的password，则默认返回""空串
	password := c.Query("password")
	
	// 返回json给浏览器
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"username": username,
		"password": password,
	})
}

// 服务初始化
func (server *WebServer) init() *gin.Engine{
	server.engine_ = gin.Default()
	server.init_method()
	server.register()
	return server.engine_
}

func (server *WebServer) run(port string) {
	server.engine_.Run(port)
}

func (server *WebServer) init_method() {
	server.methods_ = append(server.methods_, HttpMethodNode{"GET", "/ping", ping_func})
	server.methods_ = append(server.methods_, HttpMethodNode{"GET", "/login", login_func})

}

func (server *WebServer) register() {
	for _,node := range  server.methods_{
		var http_method = node.http_method_
		if( http_method == "GET" ){
			var http_url = node.http_url_
			var method_func = node.http_method_func_
			server.get_method_router(http_url,method_func)
		}else if( http_method == "POST" ){
			var http_url = node.http_url_
			var method_func = node.http_method_func_
			server.post_method_router(http_url,method_func)
		}else{

		}
	}
}

// 注册get方法
func (server *WebServer) get_method_router(method string, method_func func(c *gin.Context)) {
	server.engine_.GET(method,method_func)
}

// 注册post方法
func (server *WebServer) post_method_router(method string, method_func func(c *gin.Context)) {
	server.engine_.POST(method,method_func)
}


func main() {
	server := WebServer{}
	server.init()
	server.run(":9000")
}