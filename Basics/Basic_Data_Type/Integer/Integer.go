package main

import "fmt"

/*
`uint8`就是我们熟知的`byte`型，`int16`对应C语言中的`short`型，`int64`对应C语言中的`long`型。
*/
func main() {
	//无符号类型
	var ua uint8 = 255
	var ub uint16 = 256
	var uc uint32 = 666
	var ud uint64 = 777
	fmt.Printf("uint8:%d uint16:%d uint32:%d uint64:%d\n", ua, ub, uc, ud) //uint8:255 uint16:256 uint32:666 uint64:777

	//有符号类型
	var ia int8 = -127
	var ib int16 = -128
	var ic int32 = -666
	var id int64 = -777
	fmt.Printf("int8:%d int16:%d int32:%d int64:%d\n", ia, ib, ic, id) //int8:-127 int16:-128 int32:-666 int64:-777

	//特殊整型 不同平台上的差异
	/*在使用`int`和 `uint`类型时，不能假定它是32位或64位的整型，
	而是考虑`int`和`uint`可能在不同平台上的差异。*/
	var a uint = 123
	var b int = 456
	var c uintptr = uintptr(a)
	fmt.Printf("uint:%d int:%d uintptr:%d\n", a, b, c) //uint:123 int:456 uintptr:123

	//数字字面量语法
	//十进制
	q := 10
	fmt.Printf("%d\n", q) //10
	//八进制 以0或0o开头
	w := 075
	e := 0o52
	fmt.Printf("%o %o\n", w, e) //75 52
	//十六进制 以0x开头
	r := 0xf0
	fmt.Printf("%X\n", r) //F0

	//特殊
	t := 123_456          //分割数字
	y := 0x1p-2           //代表十六进制的 1 除以 2²，也就是 0.25。
	fmt.Printf("%d\n", t) //123456
	fmt.Printf("%.2f", y) //0.25
}
