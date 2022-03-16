package main

import "fmt"

/*
	`uint8`类型，或者叫 byte 型，代表了`ASCII码`的一个字符。
	`rune`类型，代表一个 `UTF-8字符`。
当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。
`rune`类型实际是一个`int32`。

Go 使用了特殊的 rune 类型来处理 Unicode，
让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，
性能和扩展性都有照顾。
*/
func main() {
	var a byte = 'a'
	var b rune = '好'
	fmt.Printf("byte:%T rune:%T", a, b) //byte:uint8 rune:int32
}
