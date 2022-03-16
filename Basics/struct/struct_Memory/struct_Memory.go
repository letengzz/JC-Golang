package main

import (
	"fmt"
	"unsafe"
)

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a) //n.a 0xc000016098
	fmt.Printf("n.b %p\n", &n.b) //n.b 0xc000016099
	fmt.Printf("n.c %p\n", &n.c) //n.c 0xc00001609a
	fmt.Printf("n.d %p\n", &n.d) //n.d 0xc00001609b

	//空结构体
	//空结构体是不占用空间的
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0
}
