package main

import "fmt"

func a() {
	fmt.Println("func a")
}
func b() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}
func c() {
	fmt.Println("func c")
}
func main() {
	a()
	b()
	c()
}
