package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"unicode"
)

//const ApiString = "https://gw.tvs.qq.com/echo?msg=abc123xyz"
//
//type AutoGenerated struct {
//	EchoMsg string `json:"echoMsg"`
//	Result  Result `json:"result"`
//}
//type Result struct {
//	Code       int    `json:"code"`
//	Message    string `json:"message"`
//	SessionID  string `json:"sessionId"`
//	RandomData string `json:"randomData"`
//}

func main() {
	s, err := getRespString()
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonStruct AutoGenerated
	err = json.Unmarshal([]byte(s), &jsonStruct)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", jsonStruct)
	sr := []rune(s)
	for k, v := range sr {
		if unicode.IsDigit(v) {
			sr[k] = '-'
		}
	}
	fmt.Println(s)
	f, err := os.Create("tx.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(string(sr))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getRespString() (string, error) {
	c := http.Client{}
	resp, err := c.Get(ApiString)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	jsonByte := make([]byte, 10)
	var jsonString string
	for {
		n, err := resp.Body.Read(jsonByte)
		if err == io.EOF {
			break
		}
		//fmt.Println(n, string(jsonByte[:n]))
		jsonString += string(jsonByte[:n])
	}
	return jsonString, nil
}

func getRespString2() (string, error) {
	url := ApiString
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("接口报错:", err)
		return "", err
	}
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("接口报错:", err)
		return "", err
	}
	return string(b), nil
}

// curl "https://gw.tvs.qq.com/echo?msg=abc123xyz" -o test.txt
// sed -i "" "s/[0-9]/-/g" test.txt
