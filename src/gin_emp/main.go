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