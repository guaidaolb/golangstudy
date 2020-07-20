package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("./hello.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	var fileStr string
	reader := bufio.NewReader(file)
	for {

		str, err := reader.ReadString('\n')

		if err == io.EOF {
			fileStr += str
			fmt.Println("文件读取完毕")
			break
		}

		if err != nil {
			fmt.Println("读取文件失败")
			return
		}
		fileStr += str
	}

	fmt.Println(fileStr)

}
