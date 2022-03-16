package main

import "fmt"

//将MyInt定义为int类型

//MyInt 基于int类型的自定义类型
type MyInt int

//NewInt int类型别名
type NewInt = int

func main() {
	var a MyInt = 50
	fmt.Printf("type:%T value:%v\n", a, a) //type:main.MyInt value:50

	var b NewInt
	fmt.Printf("type:%T value:%v\n", b, b) //type:int value:0
}
