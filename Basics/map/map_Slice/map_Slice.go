package main

import "fmt"

func main() {
	var sliceMap = make(map[string][]int, 8) //只完成了map的初始化
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		sliceMap["中国"] = make([]int, 8)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 100

	}
	//遍历sliceMap
	for s, ints := range sliceMap {
		fmt.Println(s, ints)
	}
}
