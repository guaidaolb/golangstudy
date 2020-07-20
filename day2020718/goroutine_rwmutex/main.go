package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

// Write 执行写操作
func Write() {
	mutex.Lock()
	fmt.Println("--执行写操作--")
	time.Sleep(time.Second)
	mutex.Unlock()
	wg.Done()
}

// Read 执行读操作
func Read() {
	mutex.RLock()
	fmt.Println("执行读操作")
	time.Sleep(time.Second)
	mutex.RUnlock()
	wg.Done()
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Write()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Read()
	}

	wg.Wait()
}
