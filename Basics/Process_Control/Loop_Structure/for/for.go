package main

import "fmt"

/*条件表达式返回`true`时循环体不停地进行循环，直到条件表达式返回`false`时自动退出循环。*/
func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i) //0123456789
	}
	fmt.Println()

	//for循环的初始语句可以被忽略
	m := 0
	for ; m < 5; m++ {
		fmt.Print(m) //01234
	}
	fmt.Println()
	//for循环的初始语句和结束语句都可以省略
	//类似于其他语言的while
	n := 0
	for n < 5 {
		fmt.Print(n) //01234
		n++
	}
	//无限循环
	//for {

	//}
}
