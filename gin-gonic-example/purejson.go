package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	h := gin.H{
		"html": "<b>hello world!</b>",
	}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, h)
	})

	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, h)
	})

	r.Run(":8080")
}
