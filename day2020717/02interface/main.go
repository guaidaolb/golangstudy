package main

import (
	"fmt"
)

//MyPrint 传入空接口类型
func MyPrint(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int 类型")
	case string:
		fmt.Println("string 类型")
	case bool:
		fmt.Println("bool 类型")
	default:
		fmt.Println("输入错误")

	}
}

func main() {
	var a interface{}
	//a = "hello world!"
	a = 20
	v, ok := a.(string)
	if ok {
		fmt.Println("mocuo tashi :", v)
	} else {
		fmt.Println("cuolou")
	}

	//x.(type)可以判断类型 只能在switch语句中使用

	MyPrint("hello world!")
}
