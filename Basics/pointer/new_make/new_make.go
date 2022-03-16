package main

import "fmt"

func main() {
	//new函数
	//new函数不太常用，使用new函数得到的是一个类型的指针，
	//并且该指针对应的值为该类型的零值
	a := new(int)
	b := new(bool)
	var c *int
	a = new(int)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Printf("%T\n", c) // *int

	fmt.Println(*a) // 0
	fmt.Println(*b) // false
	fmt.Println(*c) // 0
	//make
	//它只用于slice、map以及chan的内存创建
	var d map[string]int
	d = make(map[string]int, 10)
	d["哈哈哈"] = 100
	fmt.Println(d)
}
