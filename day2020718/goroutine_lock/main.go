package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var count int

func test() {
	mutex.Lock()
	count++
	fmt.Printf("the count is : %v \n", count)
	time.Sleep(time.Millisecond * 50)
	mutex.Unlock()
	wg.Done()
}

func main() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go test()
	}

	wg.Wait()

}
