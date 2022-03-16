package main

import "fmt"

func main() {
	//判断某个键存不存在
	a := make(map[string]int, 5)
	a["李白"] = 100
	a["杜甫"] = 50
	//a["曹操"] = 500

	//判断在不在
	//value, ok := map[key]
	value, ok := a["曹操"]
	//如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	fmt.Println(value, ok) //0 false
	if ok {
		fmt.Printf("在a中，值为%d", value)
	} else {
		fmt.Println("没有在a中") //没有
	}
}
