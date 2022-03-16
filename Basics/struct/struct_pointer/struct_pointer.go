package main

import "fmt"

type person struct {
	name, city string
	age        int8
}

func main() {
	//结构体指针
	var p = new(person)
	//(*p).name = "你好"
	//(*p).city = "nihao"
	p.name = "你好"
	p.city = "nihao"
	fmt.Println(p.name)
	fmt.Printf("%#v\n", p)

	//取结构体的地址进行实例化
	p2 := &person{}
	p2.name = "nii"
	fmt.Printf("%#v \n%T\n", p2, p2)
	fmt.Printf("%#v", p2)
}
