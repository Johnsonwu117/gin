package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "0")

		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	router.Run()
}
