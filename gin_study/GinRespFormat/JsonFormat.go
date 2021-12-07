package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//http://localhost:8080/hellojson
	engine.GET("hellojson", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())

		context.JSON(200, map[string]interface{}{
			"name": "liuxing",
			"age":  16,
		})
	})

	engine.GET("jsonstruct", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())

		var resp = RespJson{Code: 6, Message: "JSON IS HERE", Data: context.FullPath()}

		context.JSON(200, &resp)
	})

	engine.Run()
}

type RespJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
