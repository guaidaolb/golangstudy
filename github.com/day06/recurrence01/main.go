package main

import "fmt"

func main() {

	fn(3)

}

func fn(i int) {
	if i > 0 {
		fmt.Println(i)
		i--
		fn(i)
	}
}
