package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	router.POST("/somePost", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	router.PUT("/somePut", func(c *gin.Context) {
		c.String(http.StatusOK, "put")
	})
	router.DELETE("/someDelete", func(c *gin.Context) {
		c.String(http.StatusOK, "delete")
	})
	router.PATCH("/somePatch", func(c *gin.Context) {
		c.String(http.StatusOK, "patch")
	})
	router.HEAD("/someHead", func(c *gin.Context) {
		c.Header("foo", "bar")
		c.String(http.StatusOK, "head")
	})
	router.OPTIONS("/someOptions", func(c *gin.Context) {
		c.String(http.StatusOK, "option")
	})

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run()
	// router.Run(":3000") hardcode 端口号
}
