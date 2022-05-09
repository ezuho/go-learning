package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "loginV1")
		})

		v1.GET("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submitV1")
		})

		v1.GET("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "readV1")
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "loginV2")
		})

		v2.GET("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submitV2")
		})

		v2.GET("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "readV2")
		})
	}

	router.Run(":8080")
}
