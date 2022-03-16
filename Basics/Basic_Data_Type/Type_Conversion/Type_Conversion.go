package main

import (
	"fmt"
	"math"
)

/*
Go语言中只有强制类型转换，没有隐式类型转换。
该语法只能在两个类型之间支持相互转换的时候使用。
基本语法: T(表达式)
*/
func main() {
	var i1, i2 = 50, 60
	var f float64
	f = float64(i1 + i2)
	fmt.Println(f) //110

	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c) //5
}
