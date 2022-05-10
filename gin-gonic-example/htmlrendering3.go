package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	html := template.Must(template.ParseFiles("templates/index.tmpl"))
	r.SetHTMLTemplate(html)
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Custom template render",
		})
	})
	r.Run(":8080")
}
