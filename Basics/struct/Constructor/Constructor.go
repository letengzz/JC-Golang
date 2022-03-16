package main

import "fmt"

//结构体构造函数:构造一个结构体实例的函数
//结构体是值类型
//当结构体很大时，开销会比较大，通常使用返回值是指针 指针是固定长度 会节省开销一点
type person struct {
	name, city string
	age        int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func main() {
	p1 := newPerson("李白", "西安", int8(1))
	fmt.Printf("type:%T value:%#v\n", p1, p1)
	//type:*main.person value:&main.person{name:"李白", city:"西安", age:1}
}
