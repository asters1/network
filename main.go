package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asters1/tools"
)

func main() {

	method := os.Args[1]
	url := os.Args[2]

	content, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println("读取文件失败, 错误:", err)
		return
	}
	header := string(content)
	data := ""
	if method == "post" {
		content1, err := ioutil.ReadFile(os.Args[4])
		if err != nil {
			fmt.Println("读取文件失败, 错误:", err)
			return
		}
		data = string(content1)

	}

	a := tools.RequestClient(url, method, header, data)
	fmt.Println(a)
}
