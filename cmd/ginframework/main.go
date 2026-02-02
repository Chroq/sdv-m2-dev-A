package main

import (
	"fmt"
	"time"

	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Log Middleware
	router.Use(func(ctx *gin.Context) {
		fmt.Println(time.Now().Format(time.RFC3339), ctx.Request.Method, ctx.Request.URL.Path)
		ctx.Next()
	})

	// Header Middleware
	router.Use(ginhelmet.Default())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	router.POST("/hello", func(c *gin.Context) {
		c.String(200, "Hello World "+c.PostForm("name"))
	})

	router.Run(":8080")
}
