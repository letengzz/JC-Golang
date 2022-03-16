package main

import "fmt"

type Address struct {
	Province string
	City     string
}
type Person struct {
	Name   string
	Gender string
	Age    int
	//Address Address //嵌套另外一个结构体
	Address //结构体的匿名嵌套
}

func main() {
	//嵌套结构体
	p1 := Person{
		Name:   "小明",
		Gender: "男",
		Age:    10,
		Address: Address{
			Province: "河北",
			City:     "唐山",
		},
	}
	fmt.Println(p1)
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Address.Province) //通过嵌套的匿名结构访问其内部的字段
	fmt.Println(p1.Province)         //直接访问匿名结构体中的字段

}
