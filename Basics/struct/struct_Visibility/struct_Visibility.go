package main

import "fmt"

//如果一个Go语言包中定义的标识符是首字母大写的，那么就是对外可见的
//如果一个结构体中的字段名首字母是大写的，那么该字符就是对外可见的

type student struct {
	ID   int
	Name string
}
type class struct {
	Title    string
	Students []student
}

//student 的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}
func main() {
	//结构体字段的可见性与json序列化
	//创建一个班级变量c1
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}
	//往班级c1中添加
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)
}
