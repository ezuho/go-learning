package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates_advanced")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Welcome!"})
	})
	r.GET("/article", func(c *gin.Context) {
		c.HTML(http.StatusOK, "article.html", gin.H{"title": "Html5 Article Engine"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		log.Println(files)
		log.Println(filepath.Base(include))
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
