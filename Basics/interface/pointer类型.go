package main

import "fmt"

//使用值接收者实现接口和使用指针接收者实现接口的区别

type mover interface {
	move()
}
type person1 struct {
	name string
	age  int8
}

//使用值接收者实现接口:类型的值和类型的指针都能保存到接口变量中
//func (p person1) move() {
//	fmt.Printf("%s在跑\n", p.name)
//}

//使用指针接收者实现接口:只有类型指针能保存到接口变量中
func (p *person1) move() {
	fmt.Printf("%s在跑\n", p.name)
}

type sayer1 interface {
	say()
}

func (p *person1) say() {
	fmt.Printf("%s在叫\n", p.name)
}

//接口的嵌套
type animal interface {
	mover
	sayer1
}

//空接口

func main() {
	var m mover //定义一个mover类型的变量
	//p1 := person1{	//p1 是person1类型的值
	//	name: "小白",
	//	age:  50,
	//}
	p2 := &person1{ //p2是person1类型的指针
		name: "刘备",
		age:  15,
	}
	//m = p1  无法保存 因为p1是person1类型的值没有实现mover接口
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer1 //定义一个sayer类型的变量
	s = p2
	s.say()
	fmt.Println(s)

	var a animal
	a = p2
	a.say()
	a.move()
	fmt.Println(a)
}
