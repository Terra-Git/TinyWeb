package web_server

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// ping 回调
func Ping_func(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 用户表，后续引入mysql？
var user_map = map[string]string {"gyk":"gyk123",
                                  "test":"test123"}

// 登录回调
func Login_func(c *gin.Context) {
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