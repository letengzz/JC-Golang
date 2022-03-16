package main

import "fmt"

/*
Go语言规定每个`switch`只能有一个`default`分支。

一个分支可以有多个值，多个`case`值中间使用英文逗号分隔。
*/
func main() {
	f := 3
	switch f {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("其他")
	}

	//分支中有多个值
	switch c := 6; c {
	case 1, 2, 3, 4:
		fmt.Println("小于5")
	case 5:
		fmt.Println("等于5")
	case 6, 7, 8, 9:
		fmt.Println("大于5且小于10")
	}

	//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	//fallthrough语法可以执行满足条件的case的下一个case
	//是为了兼容C语言中的case设计的
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
