package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ObjectData struct {
	Objects []ObjectItem `json:"objects" form:"objects"`
}

type ObjectItem struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func main() {
	engine := gin.Default()

	engine.GET("/test_param/:name/*grade", test_param)
	engine.GET("/test_get", test_get)
	engine.POST("/test_post", test_post)

	engine.Run(":8080")
}

func test_param(c *gin.Context) {
	fmt.Println(c.Params.ByName("name"))
	fmt.Println(c.Params)
	fmt.Println(c.Param("name"))
	fmt.Println(c.Param("grade"))
}

func test_get(c *gin.Context) {
	fmt.Println(c.Query("c"))
	fmt.Println(c.DefaultQuery("d", "null"))
}

func test_post(c *gin.Context) {
	//var d ObjectData
	//err := c.ShouldBindJSON(&d)
	////err := c.ShouldBind(&d)
	//if err != nil {
	//	fmt.Errorf("BIND JSON Failed.\n")
	//	return
	//}
	//fmt.Printf("%#v\n", d)

	fmt.Printf("%#v\n", c.PostFormMap("test"))
	fmt.Printf("%#v\n", c.PostForm("form1"))
	fmt.Printf("%#v\n", c.PostForm("form3"))
	fmt.Printf("%#v\n", c.DefaultPostForm("form0", "null_form"))
}
