package main

import (
	"errors"
	"fmt"
)

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

//函数作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
func main() {
	//定义函数类型
	var a, b calculation // 声明一个calculation类型的变量
	a = add              // 把add赋值给a
	b = sub              // 把sub赋值给b

	fmt.Printf("type of :%T\n", a) // type of c:main.calculation
	fmt.Println(a(10, 220))        //230

	fmt.Printf("type of :%T\n", b) // type of c:main.calculation
	fmt.Println(b(100, 20))        //80

	//函数作为参数
	ret2 := calc(10, 20, add)
	fmt.Println(ret2) //30

	//函数作为返回值
	c, d := do("+")
	e := c(50, 60)
	fmt.Println(e, d) //110 <nil>
}
