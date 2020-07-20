package main

import (
	"fmt"
	"time"
)

func main() {

	// time.Now()
	// ticker := time.NewTicker(time.Second)
	// n := 5
	// for t := range ticker.C {
	// 	n--
	// 	fmt.Println(t)
	// 	if n == 0 {
	// 		ticker.Stop()
	// 		break
	// 	}

	// }

	fmt.Println("aaa1")
	time.Sleep(time.Second)
	fmt.Println("aaa2")
	time.Sleep(time.Second)
	fmt.Println("aaa3")
	time.Sleep(time.Second)
	fmt.Println("aaa4")
	time.Sleep(time.Second * 3)
	fmt.Println("aaa5")
}
