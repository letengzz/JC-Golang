package main

import "fmt"

func main() {
	a := map[string]int{
		"ff": 100,
		"aa": 200,
	}
	fmt.Println(a) //map[aa:200 ff:100]

	//删除键值对
	//delete(map, key)
	delete(a, "aa")
	fmt.Println(a) //map[ff:100]
}
