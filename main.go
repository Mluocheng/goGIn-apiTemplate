package main

import (
	"goGIn-apiTemplate/core"

	"github.com/gin-gonic/gin"
)

func main() {
	core.InitializeViper()

	// 初始化gin对象
	g := gin.Default()

	// 设置一个get请求，其URL为/hello，并实现简单的响应
	g.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})

	// 启动服务
	g.Run()
}
