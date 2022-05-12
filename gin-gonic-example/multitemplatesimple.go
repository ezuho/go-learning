package main

import "C"
import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/index.tmpl")
	r.AddFromFiles("article", "templates/index.tmpl")
	return r
}
func main() {
	r := gin.Default()
	r.HTMLRender = createMyRender()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{"title": "Html5 Template Engine"})
	})
	r.GET("/article", func(c *gin.Context) {
		c.HTML(http.StatusOK, "article", gin.H{"title": "Html5 Article Engine"})
	})

	r.Run(":8080")
}
