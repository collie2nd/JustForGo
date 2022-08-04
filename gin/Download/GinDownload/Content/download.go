package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/download", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")
		// filename 中文乱码问题 filename*=utf-8''xxx
		//disposition := fmt.Sprintf("attachment; filename*=utf-8''%s", url.QueryEscape("工作簿1.xlsx"))
		disposition := "attachment; filename=下载测试.xlsx"
		c.Header("Content-Disposition", disposition)
		//file, err := ioutil.ReadFile("../../PostgreSQL数据库内核分析.pdf")
		file, err := ioutil.ReadFile("../../工作簿1.xlsx")
		if err != nil {
			log.Println(err)
		}
		c.Writer.Write(file)
	})

	router.Run(":8080")
}
