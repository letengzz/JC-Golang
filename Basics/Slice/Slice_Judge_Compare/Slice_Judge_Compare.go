package main

import "fmt"

func main() {
	a := make([]int, 3, 10)
	var b []int
	c := []int{}

	//判断为空使用len
	if len(a) == 0 {
		fmt.Println("该切片为空")
	}

	//比较切片 切片唯一合法的比较操作是和`nil`比较
	//一个`nil`值的切片并没有底层数组

	//比较a是否为nil:a不为nil
	if a == nil {
		fmt.Println("a为nil")
	} else {
		fmt.Println("a不为nil")
	}
	fmt.Println(a, len(a), cap(a)) //[0 0 0] 3 10

	//比较b是否为nil:b为nil
	if b == nil {
		fmt.Println("b为nil")
	} else {
		fmt.Println("b不为nil")
	}
	fmt.Println(b, len(b), cap(b)) //[] 0 0

	////比较c是否为nil:c不为nil
	if c == nil {
		fmt.Println("c为nil")
	} else {
		fmt.Println("c不为nil")
	}
	fmt.Println(c, len(c), cap(c)) //[] 0 0

}
