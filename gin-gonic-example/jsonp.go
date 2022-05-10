package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// callback is x
		// will output : x{\"foo\": \"bar\"}
		c.JSONP(http.StatusOK, data)
	})

	r.Run(":8080")
}
