package main

import "fmt"

/*
`iota`是go语言的常量计数器，只能在常量的表达式中使用。

`iota`在`const`关键字出现时将被重置为0。
const中每新增一行常量声明将使`iota`计数一次(`iota`可理解为`const`语句块中的行索引)。
使用`iota`能简化定义，在定义枚举时很有用。
*/
func main() {
	const (
		n1 = iota //0
		n2        //1
		n3        //2
	)
	fmt.Println(n1, n2, n3) //0 1 2

	//	使用_跳过某些值
	const (
		m1 = iota //0
		m2        //1
		_
		m4 //3
	)
	fmt.Println(m1, m2, m4) //0 1 3

	//	iota声明中间插队
	const (
		k1 = iota //0
		k2 = 100  //100
		k3 = iota //2
		k4        //3
	)
	const k5 = iota                 //0
	fmt.Println(k1, k2, k3, k4, k5) //0 100 2 3 0

	const (
		_  = iota             //0
		KB = 1 << (10 * iota) //1<< 10*1
		MB = 1 << (10 * iota) //1<< 10*2
		GB = 1 << (10 * iota) //1<< 10*3
		TB = 1 << (10 * iota) //1<< 10*4
		PB = 1 << (10 * iota) //1<< 10*5
	)
	fmt.Println(KB, MB, GB, TB, PB) //1024 1048576 1073741824 1099511627776 1125899906842624

	//多个iota定义在一行
	const (
		a, b = iota + 1, iota + 2 //0 + 1,0 + 2
		c, d                      //2,3
		e, f                      //3,4
	)
	fmt.Println(a, b, c, d, e, f) //1 2 2 3 3 4
}
