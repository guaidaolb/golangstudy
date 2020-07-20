package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(n int) {
	defer wg.Done()
	for x := (n - 1) * 30000; x < n*30000; x++ {
		if x > 1 {
			var flag = true
			for y := 2; y < x; y++ {
				if x%y == 0 {
					flag = false
					break
				}
			}
			if flag {
				fmt.Printf("%v是素数。\n", x)
			}
		}
	}
}

func main() {
	var start = time.Now().Unix()
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	var end = time.Now().Unix()
	fmt.Printf("打印完毕。。总耗时%v秒", end-start)
}
