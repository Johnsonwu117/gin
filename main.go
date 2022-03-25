package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", gin.H{
			"title": "johnson welcome",
		})
	})

	router.Run(":8000")
}
