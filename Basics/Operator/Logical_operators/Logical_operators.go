package main

import "fmt"

func main() {
	a := 500
	b := 600
	//逻辑与 如果两边的操作数都是 true，则为 true，否则为 false。
	fmt.Println("a+b>500且b-a<=50:", a+b > 500 && b-a <= 50)
	//逻辑或 如果两边的操作数有一个 true，则为 true，否则为 false。
	fmt.Println("a+b>500或b-a<=50:", a+b > 500 && b-a <= 50)
	//逻辑非 如果条件为 true，则为 false，否则为 true。
	c := !(a+b > 500)
	fmt.Println("a+b>500的非:", c)
}
