package main

import "fmt"

//常量定义了之后不可以修改
//在程序运行期间不会改变的量
//批量声明常量时，如果某一行声明后没有赋值，就默认和上一行一样
const pi = 3.1415926
const (
	statusOK = 200
	notFound = 404
)

//iota
const (
	a1 = iota
	a2
	a3
)

var (
	name string
	age  int
	isOk bool
)

func main() {
	fmt.Println("Hello World！")

	name = "lixiang"
	age = 16
	isOk = true
	//非全局变量声明后必须使用，否则编译不通过

	// fmt.Printf("name:%s", name) //%s占位符，用后面的变量去填充
	fmt.Printf("name=%s name的类型是%T", name, name)
	fmt.Print("  ")
	fmt.Println(age) //println打印完后自动换行
	fmt.Print(isOk)
	fmt.Print("  ")
	//简短变量声明 只能在函数内部使用
	s1 := "haha"
	fmt.Println(s1)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
}
