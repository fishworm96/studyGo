package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func (c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"hello": "world",
		})
	})
	r.Run(":8080")
}