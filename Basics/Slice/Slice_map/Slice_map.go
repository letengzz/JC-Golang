package main

import "fmt"

func main() {
	//元素类型为map的切片
	mapSlice := make([]map[string]int, 8, 8) //完成了切片的初始化
	//还需要完成内部的map初始化
	mapSlice[0] = make(map[string]int, 8) //完成了map的初始化

	mapSlice[0]["你好"] = 100
	fmt.Print(mapSlice) //[map[你好:100] map[] map[] map[] map[] map[] map[] map[]]
}
