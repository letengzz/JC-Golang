package main

import "fmt"

func main() {
	a := 0b001011
	b := 0b110011
	fmt.Printf("%6b\n", a&b)  //按位与
	fmt.Printf("%6b\n", a|b)  //按位或
	fmt.Printf("%6b\n", a^b)  //按位异或
	fmt.Printf("%6b\n", a<<2) //左移
	fmt.Printf("%6b\n", b>>2) //右移
}
