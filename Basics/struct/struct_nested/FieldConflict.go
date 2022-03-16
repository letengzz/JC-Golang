package main

import "fmt"

//嵌套结构体的字段冲突

//Address 地址结构体
type Address1 struct {
	Province   string
	City       string
	UpdateTime string
}

//Email 邮箱结构体
type Email struct {
	Addr       string
	UpdateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Age    int
	//Address Address //嵌套另外一个结构体
	//结构体的匿名嵌套
	Address1 //嵌套另外一个结构体
	Email    //嵌套另一个结构体
}

func main() {
	p1 := User{
		Name:   "小王",
		Gender: "男",
		Age:    45,
		Address1: Address1{
			Province:   "山东",
			City:       "菏泽",
			UpdateTime: "2013-01-03",
		},
		Email: Email{
			Addr:       "123123@163.com",
			UpdateTime: "2013-05-06",
		},
	}
	fmt.Printf("%#v\n", p1)
	//fmt.Println(p1.UpdateTime) 不明确的引用 'UpdateTime'
	fmt.Println(p1.Address1.UpdateTime) //指定Address结构体可以使用
	fmt.Println(p1.Email.UpdateTime)    //指定Email结构体可以使用

}
