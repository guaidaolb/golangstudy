package main

import (
	"fmt"
	"sort"
)

func main() {

	var numSlice = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	//选择排序
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
	}
	fmt.Println(numSlice)
	//冒泡排序
	for i := 0; i < len(numSlice); i++ {
		for j := 0; j < len(numSlice)-1-i; j++ {
			if numSlice[j] > numSlice[j+1] {
				temp := numSlice[j]
				numSlice[j] = numSlice[j+1]
				numSlice[j+1] = temp
			}
		}
	}
	fmt.Println(numSlice)
	//sort包 降序排序
	var intSlice = []int{9, 22, 7, 89, 36, 15, 34, 23, 11}
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice)

}
