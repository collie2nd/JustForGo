package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"unicode"
)

const ApiStr = "https://gw.tvs.qq.com/echo?msg=12%E4%B8%AD%E6%96%87%E6%B6%88%E6%81%AFabc&hello=world"

type EchoResp struct {
	EchoMsg string     `json:"echoMsg"`
	Result  EchoResult `json:"result"`
}

type EchoResult struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	SessionID  string `json:"sessionId"`
	RandomData string `json:"randomData"`
}

func main() {
	getApiResp()
	transMoney(123456789.123)
}

func getApiResp() {
	c := http.Client{}
	r, err := c.Get(ApiStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	rBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(rBytes))

	var newRes EchoResp
	err = json.Unmarshal(rBytes, &newRes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", newRes) // a i answer

	for k, v := range rBytes {
		if unicode.IsDigit(rune(v)) {
			rBytes[k] = '-'
		}
	}
	fmt.Println(string(rBytes)) // a ii answer

	f, err := os.Create("./text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = io.WriteString(f, string(rBytes))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func transMoney(num float64) string {
	numStr := strconv.FormatFloat(num, 'f', -1, 64)
	numBytes := []rune(numStr)
	fmt.Println(numBytes)
	numMap := map[int]string{
		0: "零",
		1: "壹",
		2: "贰",
		3: "叁",
		4: "肆",
		5: "伍",
		6: "陆",
		7: "柒",
		8: "捌",
		9: "玖",
	}
	//indexSlice := []string{"厘", "分", "角", "圆", "十", "佰", "仟", "万", "亿"}
	var res []string
	for _, v := range numBytes {
		if unicode.IsDigit(v) {
			res = append(res, numMap[int(v)])
		}
	}
	fmt.Println(res)
	return numStr
}
