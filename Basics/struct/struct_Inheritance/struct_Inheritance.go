package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s 会动", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //匿名嵌套 嵌套了一个结构体指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪", d.name)
}
func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d1.move() //狗会动
	d1.wang() //狗会wang

}
