package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/:title", func(c *gin.Context) {
		name := c.Param("name")
		title := c.Param("title")
		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"title": title,
		})
	})
	r.Run(":8000")
}
