package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/download", func(c *gin.Context) {
		file, err := os.Open("../../工作簿1.xlsx")
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		//读取文件头部信息
		fileHeader := make([]byte, 512)
		file.Read(fileHeader) //取出文件头部信息

		fileStat, _ := file.Stat()

		//c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Type", http.DetectContentType(fileHeader))       //返回检测到的文件类型
		c.Header("Content-Length", strconv.FormatInt(fileStat.Size(), 10)) //返回文件大小
		// filename 中文乱码问题 filename*=utf-8''xxx
		//disposition := fmt.Sprintf("attachment; filename*=utf-8''%s", url.QueryEscape("工作簿1.xlsx"))
		disposition := "attachment; filename=下载测试.xlsx"
		c.Header("Content-Disposition", disposition)

		//file.Seek(0, 0)
		//io.Copy(c.Writer, file)
		c.Writer.Write(fileHeader)
		for {
			tmp := make([]byte, 1000) //通过切片长度控制流速
			n, err := file.Read(tmp)
			if err == io.EOF {
				return
			}
			c.Writer.Write(tmp[:n])
			time.Sleep(time.Second * 1) //通过sleep时间控制流速
		}
	})

	router.Run(":8080")
}
