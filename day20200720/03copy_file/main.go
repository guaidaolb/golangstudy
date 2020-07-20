package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(srcFileName string, dstFileName string) (err error) {

	file, err1 := os.Open(srcFileName)
	defer file.Close()
	fileWrite, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_RDWR, 0666)
	defer fileWrite.Close()

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	var bytestr = make([]byte, 1024)
	for {
		n1, e1 := file.Read(bytestr)
		if e1 == io.EOF {
			break
		}
		if e1 != nil {
			return e1
		}

		if _, e2 := fileWrite.Write(bytestr[:n1]); e2 != nil {
			return e2
		}

	}
	return nil
}

func main() {

	srcFileName := "D:/[Star_Trek][Voyager][1x01-02][Caretaker][GB][DVDrip][Repack].rmvb"
	dstFileName := "E:/星际迷航SE01-01-02.rmvb"

	err := CopyFile(srcFileName, dstFileName)
	if err == nil {
		fmt.Println("success！")
	} else {
		fmt.Println("fail！")
	}

}
