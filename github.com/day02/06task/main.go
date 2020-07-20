package main

import "fmt"

func main() {

	//练习1，有两个变量，a，b要求将其交换
	// var a = 34
	// var b = 10
	// var temp = a
	// a = b
	// b = temp
	// fmt.Println(a, b)

	//练习2，有两个变量，a，b要求将其交换（不能使用中间变量）
	var a = 34
	var b = 10
	a, b = b, a
	fmt.Println(a, b)
}
