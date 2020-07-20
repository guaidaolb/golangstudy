package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var ch1 = make(chan int, 10)

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {
		ch1 <- rand.Intn(1000)
	}

	close(ch1)

	for val := range ch1 {
		fmt.Println(val)
	}

}
