package web_server

import (
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

// 服务初始化
func (server *WebServer) Init() *gin.Engine{
	server.engine_ = gin.Default()
	server.init_method()
	server.register()
	return server.engine_
}

func (server *WebServer) Run(port string) {
	server.engine_.Run(port)
}

func (server *WebServer) init_method() {
	server.methods_ = append(server.methods_, HttpMethodNode{"GET", "/func/ping", Ping_func})
	server.methods_ = append(server.methods_, HttpMethodNode{"GET", "/login", Login_func})
}

func (server *WebServer) register() {
	for _,node := range  server.methods_{
		if( node.http_method_ == "GET" ){
			server.get_method_router(node.http_url_,node.http_method_func_)
		}else if( node.http_method_ == "POST" ){
			server.post_method_router(node.http_url_,node.http_method_func_)
		}else{
			return 
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
