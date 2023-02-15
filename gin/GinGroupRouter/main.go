package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//POST:http://localhost:8080/user/register
	userGroup := engine.Group("/user")
	userGroup.POST("/register", func(context *gin.Context) {
		fullPath := "用户注册：" + context.FullPath()
		fmt.Println(fullPath)

		context.Writer.Write([]byte(fullPath))

	})


	//DELETE:http://localhost:8080/user/1001
	userGroup.DELETE("/:id", func(context *gin.Context) {

		fullPath := "用户删除：" + context.FullPath()
		fmt.Println(fullPath)

		userId := context.Param("id")

		context.Writer.Write([]byte(fullPath + ":" + userId))

	})

	engine.Run()
}
