package main

import (
	"fmt"

	"sync"

	"time"
)

var wg sync.WaitGroup

//go putNum 存放数字
func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	wg.Done()
	close(intChan)
}

//go primeNum 存放素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for val := range intChan {
		var flag = true
		for i := 2; i < val; i++ {
			if val%i == 0 {
				flag = false
			}
		}
		if flag {
			primeChan <- val
		}
	}
	//每完成一个放素数的协程就在exitChan中放入一个true
	exitChan <- true
	wg.Done()
}

//存放16个存放素数协程的通道
func exitFn(primeChan chan int, exitChan chan bool) {
	for i := 0; i < 16; i++ {
		<-exitChan
	}
	close(primeChan)
	wg.Done()
}

//printPrime 打印素数
func printPrime(primeChan chan int) {
	for val := range primeChan {

		fmt.Println(val)

	}
	wg.Done()
}

func main() {

	start := time.Now().Unix()

	//声明存放数字的通道
	var intChan = make(chan int, 1000)

	//声明存放素数的通道
	var primeChan = make(chan int, 5000)

	//声明存放16个存放素数协程完成的通道
	var exitChan = make(chan bool, 16)

	//开启存放数字的协程
	wg.Add(1)
	go putNum(intChan)

	//开启16和存放素数的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	//开启打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)

	//开启判断协程
	wg.Add(1)
	go exitFn(primeChan, exitChan)

	wg.Wait()
	fmt.Print("执行完毕...总耗时")

	end := time.Now().Unix()

	fmt.Println(end-start, "秒")
}
