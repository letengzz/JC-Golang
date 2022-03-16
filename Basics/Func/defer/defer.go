package main

import "fmt"

func main() {
	//defer 延迟执行
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")
	/*
		start...
		end...
		3
		2
		1
	*/
}
