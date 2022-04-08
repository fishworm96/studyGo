package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "ok",
			})
		})
		xx := userGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "ok",
			})
		})
	}
	r.Run()
}