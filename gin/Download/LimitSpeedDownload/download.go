package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

//http 文件下载 限速下载
func main() {
	http.HandleFunc("/", download)
	//启动
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	fmt.Println(err)
}

func download(res http.ResponseWriter, req *http.Request) {

	filename := "../PostgreSQL数据库内核分析.pdf"

	//读取服务器端文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//读取文件头部信息
	fileHeader := make([]byte, 512)
	file.Read(fileHeader) //取出文件头部信息

	fileStat, _ := file.Stat()
	//定义响应头
	res.Header().Set("Content-Disposition", "attachment; filename="+filename)  //返回文件名
	res.Header().Set("Content-Type", http.DetectContentType(fileHeader))       //返回检测到的文件类型
	res.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10)) //返回文件大小

	//把文件读取指针归零
	file.Seek(0, 0)
	//io.Copy(res, file)
	for {
		tmp := make([]byte, 1000) //通过切片长度控制流速
		n, err := file.Read(tmp)
		if err == io.EOF {
			return
		}
		res.Write(tmp[:n])
		time.Sleep(time.Second * 1) //通过sleep时间控制流速
	}

}
