package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "names[first]=thinkerou&names[second]=tianou" -g "localhost:8080/post?ids[a]=1234&ids[b]=hello"
func main() {
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v\n", ids, names)
	})
	r.Run(":8080")
}
