package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestStruct StructA
	FieldB     string `form:"field_b"`
}

type StructC struct {
	NestStructPointer *StructA
	FieldC            string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func getDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	log.Println(b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestStruct,
		"b": b.FieldB,
	})
}

func getDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	log.Println(b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestStructPointer,
		"c": b.FieldC,
	})
}

func getDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	log.Println(b)
	c.JSON(http.StatusOK, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", getDataB)
	r.GET("/getc", getDataC)
	r.GET("/getd", getDataD)

	r.Run()
}
