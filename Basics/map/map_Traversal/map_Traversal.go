package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	//遍历map
	a := make(map[string]int, 5)
	a["dd"] = 10
	a["ff"] = 20

	//for range
	//map是无序的,键值对和添加的顺序无关
	for s, i := range a {
		fmt.Println(s, i)
	}

	//只遍历map中的key
	for k := range a {
		fmt.Print(k)
	}

	fmt.Println()

	//只遍历map中的value
	for _, i := range a {
		fmt.Print(i)
	}

	//按照某个固定顺序遍历map
	scoreMap := make(map[string]int, 100)

	//添加键值对
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	for s, i := range scoreMap {
		fmt.Println(s, i)
	}

	//按照key从小到大的顺序去遍历scoreMap
	//先取出所有的key存放到切片中
	keys := make([]string, 0, 100) //默认为0 不写否则为string的默认值
	for i := range scoreMap {
		keys = append(keys, i)
	}

	//对key做排序
	sort.Strings(keys)
	fmt.Println("----")

	//按照排序后的key对scoreMap排序
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
