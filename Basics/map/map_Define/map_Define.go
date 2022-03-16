package main

import "fmt"

func main() {
	//光声明map类型 但是没有初始化 a就是初始值nil
	//map[KeyType]ValueType
	var a map[string]int
	//没有初始化不能直接添加键值对
	//a["s"]=100
	fmt.Println(a == nil) //true

	//map的初始化
	//make(map[KeyType]ValueType, [cap])
	b := make(map[int]bool, 5)
	fmt.Println(b == nil) //false
	//map中添加键值对
	b[1] = true
	b[2] = false
	//值的Go语法表示
	fmt.Printf("%#v\n", b)     //map[int]bool{1:true, 2:false}
	fmt.Printf("type:%T\n", b) //type:map[int]bool

	//声明map的同时完成初始化
	c := map[string]int{
		"第一个": 1,
		"第二个": 2, //}
		//若后大括号在下面则添加一个`,`否则要跟在值的后面,如上表示
	}
	fmt.Printf("%#v\n", c)     //map[string]int{"第一个":1, "第二个":2}
	fmt.Printf("type:%T\n", c) //type:map[string]int
}
