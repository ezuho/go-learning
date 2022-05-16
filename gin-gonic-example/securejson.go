package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Using SecureJSON to prevent json hijacking. Default prepends "while(1)," to response body if the given struct is array values.
func main() {
	r := gin.Default()

	// You can also use your own secure json prefix
	// r.SecureJsonPrefix(")]}',\n")
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// Will output : while(1);["lena", "austin", "foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	r.Run(":8080")
}
