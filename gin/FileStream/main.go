package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

// map for  Http Content-Type  Http 文件类型对应的content-Type
var HttpContentType = map[string]string{
	".avi":  "video/avi",
	".mp3":  "audio/mp3",
	".mp4":  "video/mp4",
	".wmv":  "video/x-ms-wmv",
	".asf":  "video/x-ms-asf",
	".rm":   "application/vnd.rn-realmedia",
	".rmvb": "application/vnd.rn-realmedia-vbr",
	".mov":  "video/quicktime",
	".m4v":  "video/mp4",
	".flv":  "video/x-flv",
	".jpg":  "image/jpeg",
	".png":  "image/png",
}

func main() {
	e := gin.Default()
	// 根据文件路径读取返回流文件 参数url
	e.GET("/stream", func(c *gin.Context) {
		filePath := c.Query("url")
		//获取文件名称带后缀
		fileNameWithSuffix := path.Base(filePath)
		//获取文件的后缀
		fileType := path.Ext(fileNameWithSuffix)
		//获取文件类型对应的http ContentType 类型
		fileContentType := HttpContentType[fileType]
		if fileContentType == "" {
			c.String(http.StatusNotFound, "file http contentType not found")
			return
		}
		c.Header("Content-Type", fileContentType)
		c.File(filePath)
	})
	e.Run()
}
