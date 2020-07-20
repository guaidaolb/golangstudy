package main

import "fmt"

func main() {

	//打印0-50中所有的偶数
	// for i := 0; i <= 50; i++ {

	// 	if i%2 == 0 {

	// 		fmt.Println(i)
	// 	}
	// }

	//求1+2+3+...+100的和
	// var sum int
	// for i := 1; i <= 100; i++ {
	// 	sum = sum + i
	// }
	// fmt.Println(sum)

	//打印0-100之间所有是9的倍数的整数的个数及总和
	// var cont = 0
	// var sum = 0
	// for i := 1; i <= 100; i++ {
	// 	if i%9 == 0 {
	// 		sum += i
	// 		cont++
	// 	}
	// }
	// fmt.Printf("个数为%v-和为%v", cont, sum)

	//打印一个矩形
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Print("\n")
	// }
	//打印一个三角形
	// var row = 5
	// for i := 1; i <= row; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	//打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%vX%v=%v \t", j, i, i*j)
		}
		fmt.Println("")
	}

}
