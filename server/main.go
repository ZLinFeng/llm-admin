package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Hello")
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"res": "pong",
		})
	})
	r.Run()
}
