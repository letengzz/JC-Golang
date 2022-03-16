package main

import "fmt"

//结构体的定义
type person struct {
	name, city string //同样类型的字段可以写在一行
	age        int8
}

func main() {

	//键值对初始化
	p1 := person{
		name: "小李",
		city: "香港",
		age:  23,
	}
	fmt.Printf("%#v \n", p1)              //main.person{name:"小李", city:"香港", age:23}
	fmt.Println(p1.name, p1.city, p1.age) //小李 香港 23

	//对结构体指针进行键值对初始化
	p2 := &person{
		name: "小刘",
		city: "台北",
		//age:  19,	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	}
	fmt.Printf("%#v \n", p2)              //&main.person{name:"小刘", city:"台北", age:0}
	fmt.Println(p2.name, p2.city, p2.age) //小刘 台北 0

	//使用值的列表进行初始化
	//也可以用指针 方法如上相同
	//注意:
	/*1. 必须初始化结构体的所有字段。
	2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	3. 该方式不能和键值初始化方式混用。*/
	p4 := person{
		"小川",
		"青海",
		50,
	}
	fmt.Printf("%#v", p4) //main.person{name:"小川", city:"青海", age:50}
}
