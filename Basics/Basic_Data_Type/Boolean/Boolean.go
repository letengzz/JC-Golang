package main

import "fmt"

/*
Go语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true`(真)和`false`(假)两个值。

注意：
1. 布尔类型变量的默认值为`false`。
2. Go 语言中不允许将整型强制转换为布尔型.
3. 布尔型无法参与数值运算，也无法与其他类型进行转换。
*/
func main() {
	var a bool = true
	var b = false
	c := true
	fmt.Printf("%t %t %t", a, b, c) //true false true
}
