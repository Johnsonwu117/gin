package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 建立一個預設的路由引擎
	r := gin.Default()

	// GET:  請求方式: /hello:  請求的路徑
	// 當客戶端以GET的方法請求/hello路徑時,會執行後面的匿名函式
	r.GET("/hello", func(c *gin.Context) {

		// c.JSON:  返回JSON格式的資料
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/form", func(c *gin.Context) {
		// 表單引數設定預設值
		type1 := c.DefaultPostForm("type", "alert")

		// 接受其他的
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 多選框
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s, habbys is %v",
			type1, username, password, hobbys))
	})
	// 啟動HTTP服務,預設在0.0.0.0:8080啟動服務
	r.Run()
}
