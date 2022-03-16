package main

import "fmt"

type person struct {
	name, city string //同样类型的字段可以写在一行
	age        int
}

func main() {
	//var 结构体实例 结构体类型
	var p1 person
	p1.name = "小王"
	p1.age = 45
	fmt.Println(p1) //{小王 45}

	//创建指针类型结构体
	var p2 = new(person)
	p2.name = "小芳"             //支持对结构体指针直接使用`.`来访问结构体的成员
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小芳", city:"", age:0}

	//取结构体的地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "小明"             //(*p3).name = "小明"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"小明", city:"成都", age:30}

	//匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "你好好"
	user.married = false
	fmt.Println(user) //{你好好 false}
}
