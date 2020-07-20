package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := map[string]string{
		"username": "zhangsan",
		"hobby":    "playgame",
		"height":   "188cm",
		"sex":      "man",
		"id":       "123456789",
	}
	str := mapSort(map1)
	fmt.Println(str)
}

func mapSort(map1 map[string]string) string {

	var keySlice []string
	for k := range map1 {
		keySlice = append(keySlice, k)
	}
	sort.Strings(keySlice)

	var str string
	for _, v := range keySlice {
		str += fmt.Sprintf("%v=>%v\n", v, map1[v])
	}
	return str
}
