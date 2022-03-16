package main

import "fmt"

func main() {
	var s []int
	//通过var声明的零值切片可以在`append()`函数直接使用，无需初始化。
	s = append(s, 1)       // [1]
	s = append(s, 2, 3, 4) // [1 2 3 4]
	s2 := []int{5, 6, 7}
	s = append(s, s2...) // [1 2 3 4 5 6 7]

	for i, i2 := range s {
		fmt.Printf("%d的值的%d\n", i, i2)
	}
}
