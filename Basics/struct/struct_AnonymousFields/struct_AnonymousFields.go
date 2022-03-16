package main

import "fmt"

//结构体的匿名字段

// Person 结构体Person类型
type Person struct {
	string
	int
	//string
	//结构体匿名字段不能重复
}

func main() {
	p1 := Person{
		string: "小王霸",
		int:    18,
	}
	fmt.Println(p1)         //{小王霸 18}
	fmt.Printf("%#v\n", p1) //main.Person{string:"小王霸", int:18}
}
