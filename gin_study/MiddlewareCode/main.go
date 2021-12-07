package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()


	//放下面的话就仅为单个接口服务
	//engine.Use(RequestInfo())

	//GET:http://localhost:8080/query
	engine.GET("/query", RequestInfo(), func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})

	engine.Run()

}

func RequestInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("请求地址：" + context.FullPath())
		fmt.Println("请求方法：" + context.Request.Method)

		//context.Next函数可以将中间件代码的执行顺序一分为二
		//Next函数调用之前的代码在请求处理之前执行
		//当程序执行到context.Next时，会中断向下执行，转而先去执行具体的业务逻辑
		//执行完业务逻辑处理函数之后，程序会再次回到context.Next处，继续执行中间件后续的代码
		context.Next()
		fmt.Println(context.Writer.Status())

	}

}
