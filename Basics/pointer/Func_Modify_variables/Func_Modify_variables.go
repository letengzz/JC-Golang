package main

import "fmt"

func modify(x int) {
	x = 100
}
func modify2(y *int) {
	*y = 100
}
func main() {
	a := 10
	modify(a)      //不会修改值
	fmt.Println(a) //10
	modify2(&a)
	fmt.Println(a) //100
}
