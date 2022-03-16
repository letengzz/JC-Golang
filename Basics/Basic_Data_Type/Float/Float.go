package main

import (
	"fmt"
)

/*
Go语言支持两种浮点型数：`float32`和`float64`。
这两种浮点型数据格式遵循`IEEE 754`标准：
	`float32` 的浮点数的最大范围约为 `3.4e38`，可以使用常量定义：`math.MaxFloat32`。
	`float64` 的浮点数的最大范围约为 `1.8e308`，可以使用一个常量定义：`math.MaxFloat64`。

打印浮点数时，可以使用`fmt`包配合动词`%f`
*/
func main() {
	var (
		f1 float32 = 4.5
		f2 float64 = 69.4
		f3         = 9.9
	)
	f4 := 90.3
	fmt.Printf("f1:%.2f f2:%.2f f3:%.2f f4:%.2f\n", f1, f2, f3, f4)   //f1:4.50 f2:69.40 f3:9.90 f4:90.30
	fmt.Printf("f1的类型:%T f2的类型:%T f3的类型:%T f4的类型:%T", f1, f2, f3, f4) //f1的类型:float32 f2的类型:float64 f3的类型:float64 f4的类型:float64
}
