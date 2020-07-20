package main

import (
	"fmt"
	"time"
)

func main() {

	var intChan = make(chan int, 20)
	var stringChan = make(chan string, 20)

	for i := 1; i <= 20; i++ {
		intChan <- i
		stringChan <- fmt.Sprintf("hello %v world!", i)
	}

	for {
		//使用select多路复用的时候不需要关闭channel
		select {
		case v := <-intChan:
			fmt.Printf("从intChan中读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 50)

		case v := <-stringChan:
			fmt.Printf("从stringChan中读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 50)

		default:
			fmt.Println("数据读取完毕...")
			return
		}

	}

}
