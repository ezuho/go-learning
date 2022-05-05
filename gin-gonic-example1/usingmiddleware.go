package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// gin中间件(middleware)提供了类似于面向切面编程或路由拦截器的功能，可以在请求前和请求之后添加一些自定义逻辑。实际开发中有很多场景会用到中间件，例如：权限验证，缓存，错误处理，日志，事务等。
// 使用中间件
// gin的中间件分为三类：全局中间件、路由中间件、分组路由中间件。
// 全局中间件：注册全局中间件之后注册的路由才会生效，如果有一些不希望使用全局中间件的路由规则，注册路由代码要放在注册全局中间件之前。
// 路由中间件：在注册路由时传入的中间件，只对当前路由规则生效。
// 分组路由中间件：在分组路由中注册，对当前组下的全部路由生效。
func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", func(c *gin.Context) {
		log.Println("benchmark log")
		c.Next()
	}, func(c *gin.Context) {
		c.String(http.StatusOK, "benchmark")
	})

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(func(c *gin.Context) {
		log.Println("auth")
		c.Next()
	})
	{
		authorized.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "loginEndpoint")
		})
		authorized.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submitEndpoint")
		})
		authorized.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "readEndpoint")
		})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			c.String(http.StatusOK, "analyticsEndpoint")
		})
	}
	r.Run(":8080")
}
