package main

import "github.com/gin-gonic/gin"

//创建 RESTful API 端点
//要求：使⽤gin 框架创建一个 HTTP 服务器，处理 请求，返回 JSON 响应
//{"status": "ok", "code": 200}

//func main() {
//	router := gin.Default()
//	router.GET("/api/status", status)
//
//	router.Run(":8080")
//}
//
//type Response struct {
//	Status string
//	Code   int
//}
//
//func status(c *gin.Context) {
//	re := Response{Status: "ok", Code: 200}
//	c.JSON(200, re)
//}

// 优化
func main() {
	r := gin.Default()
	r.GET("/api/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"code":   200,
		})
	})
	r.Run(":8080")
}
