package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("./hello1.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	file.WriteString("直接写入字符串\r\n")

}
