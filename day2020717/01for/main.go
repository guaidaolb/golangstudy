package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().Unix()

	for x := 2; x < 120000; x++ {
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

	end := time.Now().Unix()
	fmt.Println(end - start)

}
