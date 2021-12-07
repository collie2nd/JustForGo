package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("hellobyte", func(context *gin.Context) {

		context.Writer.Write([]byte("请求路径：" + context.FullPath()))
	})
	engine.GET("hellostring", func(context *gin.Context) {

		context.Writer.WriteString("请求路径：" + context.FullPath())
	})
	engine.Run()
}
