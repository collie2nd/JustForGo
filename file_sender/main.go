package main

import (
	"flag"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"os"
)

var filePath string

var rootCmd = &cobra.Command{}

var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"s"},
	Short:   "发送文件",
	Long:    "发送文件",
	Run: func(cmd *cobra.Command, args []string) {
		err := sendHandler(filePath)
		if err != nil {
			log.Fatalf("sendHandler err:%v", err)
			return
		}
	},
}

var recvCmd = &cobra.Command{
	Use:     "recv",
	Aliases: []string{"r"},
	Short:   "接收文件",
	Long:    "接收文件",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			err := recvHandler()
			if err != nil {
				log.Fatalf("recvHandler err:%v", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(recvCmd)
	sendCmd.Flags().StringVarP(&filePath, "file", "f", "", "传输文件的路径")
	recvCmd.Flags().StringVarP(&filePath, "file", "f", "./", "文件的保存路径")
}

func sendHandler(filePath string) (err error) {
	//获取文件名info.Name()
	info, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("os.Stat err:%v", err)
		return
	}
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("net.Dial err:%v", err)
		return
	}
	//程序结束之后要关闭connect
	defer conn.Close()

	//给接收方。先发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		log.Fatalf("conn.Write err:%v", err)
		return
	}
	//接收对方的回复，如果回复OK，则开始发送内容
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatalf("conn.Read err:%v", err)
		return
	}
	if "ok" != string(buf[:n]) {
		log.Fatalln("have not receive \"ok\"")
		return
	}
	//发送文件内容
	return sendFile(filePath, conn)
}

func sendFile(filePath string, conn net.Conn) (err error) {
	//以只读方式打开文件
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("os.Open err:%v", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*4)
	//读文件，读多少发多少
	for {
		n, err := f.Read(buf) //从文件读取内容
		if err != nil {
			if err == io.EOF {
				log.Println("文件发送完毕")
				break
			} else {
				log.Fatalf("f.Read err:%v", err)
				return err
			}
		}
		//发送内容
		conn.Write(buf[:n])
	}
	return nil
}

func recvHandler() (err error) {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("net.Listen err:%v", err)
		return
	}
	defer listener.Close()

	//阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("listener.Accept err:%v", err)
		return
	}

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatalf("conn.Read() err:%v", err)
		return
	}
	defer conn.Close()

	fileName := string(buf[:n])
	//回复"OK"
	_, err = conn.Write([]byte("ok"))
	if err != nil {
		log.Fatalf("conn.Write err:%v", err)
		return
	}
	//接收内容
	return recvFile(fileName, conn)
}

func recvFile(fileName string, conn net.Conn) (err error) {
	//新建文件
	f, err := os.Create(filePath + fileName)
	if err != nil {
		log.Fatalf("os.Create err:%v", err)
		return
	}

	buf := make([]byte, 1024*4)
	//接受多少，写多少
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("文件接收完毕")
				break
			} else {
				log.Fatalf("conn.Read() err:%v", err)
				return err
			}
		}
		if n == 0 {
			log.Println("文件接收完毕")
			break
		}
		f.Write(buf[:n]) //往文件写入内容
	}
	return nil
}

func main() {
	flag.Parse()
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("rootCmd.Execute err:%v", err)
		return
	}
}
