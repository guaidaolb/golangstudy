package main

import (
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
	//fmt.Println(file)

	var byteSlice []byte
	var tempSlice = make([]byte, 128)

	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		byteSlice = append(byteSlice, tempSlice[:n]...)
	}
	fmt.Println(string(byteSlice))
}
