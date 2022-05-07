package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Println(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
func main() {
	router := gin.New()
	router.Use(Logger())

	router.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
		time.Sleep(1 * time.Second)
	})

	router.Run(":8080")
}
