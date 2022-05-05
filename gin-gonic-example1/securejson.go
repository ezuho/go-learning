package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")
	r.GET("/somejson", func(c *gin.Context) {
		name := []string{"a", "b", "c"}
		c.SecureJSON(http.StatusOK, name)
	})

	r.Run(":8080")
}
