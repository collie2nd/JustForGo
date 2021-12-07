package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		context.Writer.Write([]byte("hello,gin\n"))

	})

	if err := engine.Run(":8090"); err != nil {
		log.Fatalln(err.Error())
	}
}
