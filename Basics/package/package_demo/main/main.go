package main

//首先先导入包
import (
	//当代码都放在$GOPATH目录下时
	//导入包的路径要从$GOPATH/src后面继续写
	//前面添加名字可以用别名来使用
	clac "Example/12package/package_demo/calc" //会自动触发init函数
	"fmt"
)

//Go语言中不允许导入包而不使用
//Go语言中不允许循环引用包

//匿名导入包 只导入包 不使用包内部的数据 常用于某些数据库的驱动等
import _ "Example/12package/package_demo/calc"

func main() {
	//ret := calc.Add(10, 20)
	ret := clac.Add(10, 20)
	fmt.Println(ret)
	fmt.Println(clac.Name)
}

//多用来做一些初始化的操作
func init() {
	fmt.Println("main.init()")
}
