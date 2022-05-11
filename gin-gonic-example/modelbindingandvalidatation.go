package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//To bind a request body into a type, use model binding. We currently support binding of JSON, XML, YAML and standard form values (foo=bar&boo=baz).
//
//Gin uses go-playground/validator/v10 for validation. Check the full docs on tags usage here.
//
//Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set json:"fieldname".
//
//Also, Gin provides two sets of methods for binding:
//
//Type - Must bind
//Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML
//Behavior - These methods use MustBindWith under the hood. If there is a binding error, the request is aborted with c.AbortWithError(400, err).SetType(ErrorTypeBind). This sets the response status code to 400 and the Content-Type header is set to text/plain; charset=utf-8. Note that if you try to set the response code after this, it will result in a warning [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422. If you wish to have greater control over the behavior, consider using the ShouldBind equivalent method.
//Type - Should bind
//Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML
//Behavior - These methods use ShouldBindWith under the hood. If there is a binding error, the error is returned and it is the developer’s responsibility to handle the request and error appropriately.
//When using the Bind-method, Gin tries to infer the binder depending on the Content-Type header. If you are sure what you are binding, you can use MustBindWith or ShouldBindWith.
//
//You can also specify that specific fields are required. If a field is decorated with binding:"required" and has a empty value when binding, an error will be returned.
func main() {
	r := gin.Default()
	// Example for binding JSON ({"user": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>manu</user>
	//		<password>123</password>
	//	</root>)
	r.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for bingding a HTML form (user=manu&password=123)
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.Run(":8080")
}

type Login struct {
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	//Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"-"`
}
