package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/index", func(c *gin.Context) {
			c.String(http.StatusOK, "GET")
		})
		user.POST("/index", func(c *gin.Context) {
			c.String(http.StatusOK, "POST")
		})

		// 嵌套路由組
		boy := user.Group("/boy")
		{
			boy.GET("/index", func(c *gin.Context) {
				c.String(http.StatusOK, "boy")
			})
		}
	}

	r.Run(":7877")
}
