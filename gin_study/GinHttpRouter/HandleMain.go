package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	engine := gin.Default()
	engine.Handle("POST", "/hello", func(context *gin.Context) {
		log.Println("POST:", context.FullPath())

		name := context.PostForm("name")
		fmt.Println("name is ", name)

		context.Writer.Write([]byte("hello ," + name))
	})
	engine.Run()
}
