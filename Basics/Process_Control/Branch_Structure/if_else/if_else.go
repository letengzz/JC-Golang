package main

import "fmt"

/*
基本语法:
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
注意：Go语言规定与`if`匹配的左括号`{`必须与`if和表达式`放在同一行，`{`放在其他位置会触发编译错误。

同理，与`else`匹配的`{`也必须与`else`写在同一行，`else`也必须与上一个`if`或`else if`右边的大括号在同一行。
*/
func main() {
	a := 1
	b := 2
	if a > b {
		fmt.Println("a大于b")
	} else if a == b {
		fmt.Println("a等于b")
	} else {
		fmt.Println("a小于b")
	}

	//特殊写法 在 if 表达式之前添加一个执行语句，再根据变量值进行判断
	if c := 50; c >= 90 {
		fmt.Println("c大于90")
	} else {
		fmt.Println("c小于90")
	}
}
