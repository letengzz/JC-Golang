package main

import "fmt"

func main() {
	a := 500
	b := 600
	if a == b {
		fmt.Println("a与b相等")
	} else {
		fmt.Println("不相等")
	}
	fmt.Println("两个数是否相等:", a == b)    //两个数是否相等
	fmt.Println("两个数是否不相等:", a != b)   //两个数是否不相等
	fmt.Println("左边是否大于右边:", a > b)    //左边是否大于右边
	fmt.Println("左边是否大于等于右边:", a >= b) //左边是否大于等于右边
	fmt.Println("右边是否大于左边:", a < b)    //右边是否大于左边
	fmt.Println("右边是否大于等于左边:", a <= b) //右边是否大于等于左边
}
