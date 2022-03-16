package main

import "fmt"

/*
定义方式:
	var 数组变量名 [元素数量]T
*/
func main() {
	//数组的定义
	var a [5]int  //数组的长度必须是常量，并且长度是数组类型的一部分
	var b [10]int //`[5]int`和`[10]int`是不同的类型
	fmt.Printf("a:%#v\nb:%#v\n", a, b)
	//a:[5]int{0, 0, 0, 0, 0}
	//b:[10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	//数组会初始化为int类型的零值
	for i := 0; i < len(a); i++ {
		a[i] = i //赋值
	}
	fmt.Println(a) //[0 1 2 3 4]

	//字符数组定义
	var str [2]string
	str[0] = "world"
	str[1] = "hello"
	fmt.Println(str) //[world hello]

	//第二种方法
	str1 := [2]string{"12", "safas"}
	fmt.Println(str1) //[12 safas]

	//数组可以通过下标进行访问，下标是从`0`开始，最后一个元素下标是：`len-1`，
	//访问越界（下标在合法范围之外），则触发访问越界，会`panic`。
	s1 := string(str1[1][1])
	fmt.Print(s1) //a
}
