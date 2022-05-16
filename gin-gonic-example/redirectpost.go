package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	r.GET("/login", func(c *gin.Context) {
		form := "<form method='POST' action='/login/do'><input type='submit' /></form>"
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, form)
	})

	r.POST("/login/do", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/welcome")
	})

	r.GET("/welcome", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome")
	})

	r.Run(":8888")
}
