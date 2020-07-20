package main

import "fmt"

func main() {
	fmt.Println(fn1(100))
	fmt.Println(fn2(5))
}

//递归实现1到100的和
func fn1(i int) int {
	if i == 1 {
		return 1
	}
	return i + fn1(i-1)
}

//递归实现5的阶乘
func fn2(i int) int {
	if i == 1 {
		return 1
	}
	return i * fn2(i-1)

}
