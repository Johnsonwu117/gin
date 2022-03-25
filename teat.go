package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 注册路由和Handler
	// url为 /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		// 获取参数内容
		// 获取的所有参数内容的类型都是 string
		// 如果不存在，使用第二个当做默认内容
		firstname := c.DefaultQuery("firstname", "Guest")
		// 获取参数内容，没有则返回空字符串
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}
