package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("第(%v)个协程的第(%v)条数据\n", num, i)
		time.Sleep(time.Millisecond * 500)
	}

}

func main() {

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	fmt.Println("打印完毕。。。")
}
