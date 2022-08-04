package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//设置HTML目录
	engine.LoadHTMLGlob("./html/*")

	//设置静态文件地址
	engine.Static("img","./img")

	//http://localhost:8080/hellohtml
	engine.GET("hellohtml", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())

		//context.HTML(200, "index.html", nil)

		//定义html模板变量
		context.HTML(200, "index.html", gin.H{
			"fullpath": "请求路径：" + context.FullPath(),
		})
	})

	engine.Run()
}
