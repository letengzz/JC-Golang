package main

import "fmt"

//返回一个值
//第一种方式:
func intSum1(x, y int) int {
	return x + y
}

//第二种方式:
func intSum2(x, y int) (a int) {
	a = x + y
	return
}

//函数支持多返回值，函数如果有多个返回值时必须用`()`将所有返回值包裹起来。
func calc1(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过`return`关键字返回。
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//函数返回值类型为slice时
//func someFunc(x string) []int {
//
//	if x == "" {
//		return nil // 没必要返回[]int{}
//	}
//	...
//}
func main() {
	//Go语言中通过`return`关键字向外输出返回值。
	a := intSum1(10, 20)
	fmt.Println(a) //30
	b := intSum2(20, 30)
	fmt.Println(b) //50
	x, y := calc1(40, 30)
	fmt.Println(x, y) //70 10
	m, n := calc2(80, 60)
	fmt.Println(m, n) //140 20
}
