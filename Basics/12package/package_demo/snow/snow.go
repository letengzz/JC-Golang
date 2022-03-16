package snow

import "fmt"

//被calc包导入的一个包
func Snow() {
	fmt.Println("下雪了")
}

//首字母小写，对外不可见(只能在当前包内使用)
func init() {
	fmt.Println("snow.init()")
}
