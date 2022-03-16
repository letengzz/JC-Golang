package main

import "fmt"

//定义全局变量num
var num int = 10

func testGlobal() {
	//函数内变量会在其范围内改变其值，但不会影响全局变量的值
	num := 200
	name := "刘德华"
	//可以在函数中访问全局变量
	fmt.Println("testGlobal中的变量是", num) //函数中优先使用局部变量
	fmt.Println(name)
}
func main() {
	testGlobal()
	fmt.Println(num) //10
	//外层不能访问到内层函数的内部变量(局部变量)
	//fmt.Println(name)

	//变量i只在for循环的语句块中生效
	for i := 0; i < 5; i++ {
		if i == 5 {
			z := 200
			fmt.Println(z) //200
		}
		fmt.Print(i) //01234
	}
	//fmt.Println(i) //外层访问不到内部for语句块中的变量

}
