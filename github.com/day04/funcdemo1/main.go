package main

import (
	"fmt"
	"sort"
)

func mapSort(map1 map[string]string) string {
	var keySlice []string
	for k := range map1 {
		keySlice = append(keySlice, k)
	}
	sort.Strings(keySlice)
	var str string
	for _, v := range keySlice {
		str += fmt.Sprintf("%v==>%v\n", v, map1[v])
	}
	return str
}
func main() {

	m1 := map[string]string{
		"username": "zhangsna",
		"age":      "88",
		"sex":      "man",
		"heght":    "18889cm",
		"hobby":    "eat",
	}
	str1 := mapSort(m1)
	fmt.Println(str1)
}
