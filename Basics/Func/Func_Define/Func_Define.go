package main

import (
	"fmt"
)

//函数的定义
//定义了一个不需要参数也没有返回值的函数
func sayHello() {
	fmt.Println("Hello")
}

//定义一个带有参数没有返回值的函数
func sayHello2(name string) {
	fmt.Println("Hello", name)
}

//定义接收多个参数的函数并且有一个返回值
//第一种方法:
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

//第二种方法:
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

//函数接收可变参数,在参数后面加...表示可变参数
//可变参数在函数体中是切片类型
func intSumN(a ...int) int {
	ret := 0
	for _, i2 := range a {
		ret += i2
	}
	return ret
}

//固定参数和可变参数同时出现时，可变参数要放在最后
//go语言的函数中是没有默认参数的
func intSunN2(a int, b ...int) int {
	ret := a
	for _, i2 := range b {
		ret = i2 + ret
	}
	return ret
}

//go语言函数参数类型简写
func intSum3(a, b int) (ret int) {
	ret = a + b
	return
}

//定义具有多个返回值的函数,多返回值也支持类型简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	//函数调用
	//调用了一个不需要参数也没有返回值的函数
	sayHello() //Hello

	//调用了一个带有参数没有返回值的函数
	name := "李白"
	sayHello2(name) //Hello 李白

	//调用了接收两个参数的函数并且有一个返回值
	r := intSum(12, 22)   //第一种方法
	r0 := intSum2(12, 20) //第二种方法
	//也可以不接受值
	//intSum(12，33)
	fmt.Println(r, r0) //34 32

	r1 := intSumN(10, 20)
	r2 := intSumN(10, 555, 66)
	fmt.Println(r1) //30
	fmt.Println(r2) //631

	rr1 := intSunN2(10, 50) //a=10 b=[50]
	rr2 := intSunN2(10)     //a = 10 b = []
	fmt.Println(rr1)        //60
	fmt.Println(rr2)        //10

	x, y := calc(100, 200) //多个返回值的调用
	fmt.Println(x, y)      //300 -100
}
