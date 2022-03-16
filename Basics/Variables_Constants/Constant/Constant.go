package main

import "fmt"

//全局变量
const a = 3.14 //3.14

func main() {
	const p = 3.14 //3.14
	//多个常量也可以一起声明
	const (
		e  = 50 //50
		f  = 60 //60
		n1 = 70 //70
		//`const`同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
		n2 //70
		n3 //70
	)
	fmt.Printf("a:%.2f p:%.2f e:%d f:%d n1:%d n2:%d n3:%d", a, p, e, f, n1, n2, n3) //a:3.14 p:3.14 e:50 f:60 n1:70 n2:70 n3:70
}
