package main

import "fmt"

func main() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}

	b := a         //a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]

	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
	c[0] = 1
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
}
