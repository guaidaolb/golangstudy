package main

import (
	"fmt"
	"strings"
)

func main() {

	// var str1 string = "你好 golang"
	// var str2 = "你好 go"
	// str3 := "nihao"

	// fmt.Printf("%v---%T\n", str1, str1)
	// fmt.Printf("%v---%T\n", str2, str2)
	// fmt.Printf("%v---%T\n", str3, str3)

	// var str4 = `
	// this is str
	// this is str
	// this is str
	// this is str
	// this is str
	// `
	// fmt.Println(str4)

	var str1 = "123-456-789"
	arr := strings.Split(str1, "-")
	fmt.Println(arr)
	var str2 = strings.Join(arr, "*")
	fmt.Println(str2)

}
