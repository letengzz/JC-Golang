package main

import "fmt"

/*
Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

变量声明: Go语言中的变量需要声明后才能使用，
同一作用域内不支持重复声明。并且**Go语言的变量声明后必须使用。
*/
/*
1. 整型和浮点型变量的默认值为`0`。
2. 字符串变量的默认值为`空字符串`。
3. 布尔型变量默认为`false`。
4. 切片、函数、指针变量的默认为`nil`。
*/
/*
Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。
并且Go语言的变量声明后必须使用
*/
// 全局变量m
var m = 100

func foo() (int, string) {
	return 10, "HH"
}

func main() {
	//Go变量声明格式
	//var 变量名 变量类型
	var a int
	fmt.Println(a) //0

	//批量声明
	var (
		b string
		c int8
	)
	fmt.Println(b, c) // 0

	//在声明变量的时候为其指定初始值。变量初始化的标准格式如下：
	//	var 变量名 类型 = 表达式
	var name string = "nihao"
	var age int = 50
	fmt.Printf("%s %d\n", name, age) //nihao 50

	//类型推导:编译器会根据等号右边的值来推导变量的类型完成初始化。
	var n = "ff"
	var age1 = 21
	fmt.Printf("%s %d\n", n, age1) //ff 21

	//一次初始化多个变量
	var d, e = "aad", 20
	fmt.Printf("%s,%d\n", d, e) //aad,20

	//在函数内部，可以使用更简略的 `:=` 方式声明并初始化变量。
	//注意:`:=`不能使用在函数外。
	f := 50
	fmt.Println(f) //50

	//在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量
	//匿名变量用一个下划线`_`表示
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x) //x=10
	fmt.Println("y=", y) //y=HH
	//注意：
	//1. 函数外的每个语句都必须以关键字开始（`var`、`const`、`func`等）
	//2. `_`多用于占位，表示忽略值。
}
