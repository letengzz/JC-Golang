package main

import "fmt"

const MAX int = 3

//ptr 为整型指针数组。因此每个元素都指向了一个值。
func main() {
	arr := [MAX]int{1, 20, 300}
	var ptr [MAX]*int

	for i := 0; i < MAX; i++ {
		ptr[i] = &arr[i] //整数地址赋值给指针数组
	}
	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
	/*
	   a[0] = 1
	   a[1] = 20
	   a[2] = 300
	*/
}
