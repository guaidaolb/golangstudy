package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("./hello.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}

	writer := bufio.NewWriter(file)
	writer.WriteString("hello golang!\r\n")
	writer.Flush()

}
