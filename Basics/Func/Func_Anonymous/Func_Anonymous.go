package main

import (
	"fmt"
)

func main() {
	//第一种方式 使用变量来接收匿名函数
	sayHello := func() {
		fmt.Println("你好")
	}
	sayHello()

	//第二种方式 直接执行匿名函数
	func() {
		fmt.Println("匿名函数")
	}()
}
