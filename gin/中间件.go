package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func StatCost() gin.HandlerFunc  {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子")
		c.Next()
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	r := gin.Default()

	// 为全局注册中间件
	r.Use(StatCost())
	r.GET("/index", handle)
	shopGroup := r.Group("/shop", StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})}
}
	r.Run()
}
