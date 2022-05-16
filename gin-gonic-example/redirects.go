package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Issuing a HTTP redirect is easy. Both internal and external locations are supported.
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com.hk/")
	})

	// Issuing a Router redirect, use HandleContext like below.
	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test22"
		r.HandleContext(c)
	})
	r.GET("/test22", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	r.Run(":8080")
}
