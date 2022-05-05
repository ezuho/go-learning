package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()
	r.Any("/testing", func(c *gin.Context) {
		var person Person
		// ShouldBindQuery function only binds the query params and not the post data.
		if c.ShouldBindQuery(&person) == nil {
			log.Println("====== Only Bind By Query String ======")
			log.Println(person.Name)
			log.Println(person.Address)
		}
		c.String(http.StatusOK, "success")
	})
	r.Run(":8080")
}
