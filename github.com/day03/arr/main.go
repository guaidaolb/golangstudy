package main

import "fmt"

func main() {

	// var arr = [...]string{"java", "php", "js", "golang", "python"}

	// for k, v := range arr {

	// 	fmt.Printf("key:%v value:%v\n", k, v)

	// }

	//求出一个数组里面的元素的和以及这些元素的平均值分别用for和for-range实现
	var arr1 = [...]int{12, 23, 45, 67, 2, 5}
	var sum = 0
	var ave = 0
	for _, v := range arr1 {
		sum += v
	}
	ave = sum / len(arr1)
	fmt.Printf("和为%v--平均值为%v", sum, ave)
	fmt.Println("")

	//请求出一个数组的最大值，并获得对应坐标
	var max = 0
	var index = 0
	for k, v := range arr1 {
		if v > max {
			max = v
			index = k
		}
	}
	fmt.Println(max, index)

}
