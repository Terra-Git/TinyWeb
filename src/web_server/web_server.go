package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "html/template"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/login", func(c *gin.Context) {
		// 当使用DefaultQuery时，如果没有获取到浏览器输入的username，则返回设置defaultValue给username
		username := c.Query("username")
		if( username == ""){
			c.String(http.StatusOK, "invilid param username")
			return 
		}
		// 当使用Query时，如果没有获取到浏览器输入的password，则默认返回""空串
		password := c.Query("password")
		c.
		// 返回json给浏览器
		c.JSON(http.StatusOK, gin.Header{
			"message": "success",
			"username": username,
			"password": password,
		})
	})

    r.LoadHTMLFiles("load.html")
    r.GET("/html", func(c *gin.Context) {
        c.HTML(200, "load.html", "flysnow_org")
    })

	r.MaxMultipartMemory = 8 << 20
    r.POST("/upload", func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.String(500, "上传图片出错")
        }
        // c.JSON(200, gin.H{"message": file.Header.Context})
        c.SaveUploadedFile(file, file.Filename)
        c.String(http.StatusOK, file.Filename)
    })

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:9000
	r.Run(":9000")
}