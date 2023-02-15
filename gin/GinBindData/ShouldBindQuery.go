package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	//http://localhost:8080/hello?name=liuxing&classes=算法
	engine.GET("/hello", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		var student Student

		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatalln("ERR:", err.Error())
			return
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)

		context.Writer.Write([]byte("hello " + student.Name))
	})

	//http://localhost:8080/hello?username=liuxing&phone=12345678&password=1234
	engine.POST("/register", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		var register Register

		if err := context.ShouldBind(&register); err != nil {
			log.Fatalln("ERR:", err.Error())
			return
		}

		fmt.Println(register.Username)
		fmt.Println(register.Phone)

		context.Writer.Write([]byte("hello " + register.Username))
	})

	//http://localhost:8080/hello?username=liuxing&phone=12345678&password=1234
	engine.POST("/addstudent", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		var newStu NewStudent

		if err := context.BindJSON(&newStu); err != nil {
			log.Fatalln("ERR:", err.Error())
			return
		}

		fmt.Println(newStu.Name)
		fmt.Println(newStu.Sex)
		fmt.Println(newStu.Age)

		context.Writer.Write([]byte("hello " + newStu.Name))
	})

	engine.Run()

}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

type Register struct {
	Username string `form:"username"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

type NewStudent struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}
