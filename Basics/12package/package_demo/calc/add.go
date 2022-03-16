package calc

import (
	"Example/12package/package_demo/snow"
	"fmt"
)

//如果想要让其他文件访问标识符 需要首字母大写
//标识符首字母大写表示对外可见 小写其他包则访问不到
//通常不对外的标识符都是要首字母小写

//Name 定义了一个测试的全局变量
var Name string = "你好 欢迎使用calc"

//Add 是一个计算两个int类型数据和的函数
func Add(x, y int) int {
	snow.Snow()
	//同一个包中的不同文件可以直接使用
	Sub(50, 20)
	return x + y
}

//init函数在包导入的时候自动执行
//init函数没有参数也没有返回值
//首字母小写，对外不可见(只能在当前包内使用)
func init() {
	fmt.Println("calc.init()")
	fmt.Println(Name)
}
