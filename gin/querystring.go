package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/search", func (c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		addres := c.Query("address")
		c.JSON(http.StatusOK, gin.H {
			"message": "ok",
			"username": username,
			"addres": addres,
		})
	})
	r.Run(":8080")
}