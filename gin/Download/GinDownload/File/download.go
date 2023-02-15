package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/download", func(c *gin.Context) {
		c.File("../../PostgreSQL数据库内核分析.pdf")
	})
	router.Run(":8080")
}
