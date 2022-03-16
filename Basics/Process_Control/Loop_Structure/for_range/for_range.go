package main

import "fmt"

func main() {
	var a = [4]int{1, 2, 3, 5}
	for index, value := range a {
		//index索引 value数组的值
		fmt.Println("第", index, "的数是", value)
	}

	//只打印索引
	for index := range a {
		//index索引 value数组的值
		fmt.Println(index)
	}

	//只打印值
	for _, value := range a {
		//index索引 value数组的值
		fmt.Println("每一个的数是", value)
	}
}
