package main

import "fmt"

func main() {
	a := 50   //int
	ptr := &a //指针取值 *int

	fmt.Printf("a:%d ptr:%p", a, &a)
	fmt.Printf("b:%p type:%T\n", ptr, ptr)
	fmt.Println(&ptr)
	c := *ptr
	fmt.Println(c)
}
