package main

import "fmt"

type dog struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Printf("%s啊啊啊", p.name)
}

//接口不管你是什么类型 只管实现什么方法
//定义一个类型 一个抽象的类型只要实现了say这个方法的类型都可以称为sayer类型
type sayer interface {
	say()
}

//打的函数

func do(arg sayer) {
	arg.say()
}

func main() {
	c1 := cat{}
	do(c1)
	c2 := dog{}
	do(c2)
	p1 := person{
		name: "李白",
	}
	do(p1)

	//接口类型的变量
	var s sayer
	c3 := cat{}
	s = c3
	p2 := person{name: "p2"}
	s = p2
	fmt.Println(s)
}
