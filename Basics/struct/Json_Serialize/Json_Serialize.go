package main

import (
	"encoding/json"
	"fmt"
)

//如果一个Go语言包中定义的标识符是首字母大写的，那么就是对外可见的

type student struct {
	ID   int
	Name string
}
type class struct {
	Title    string    `json:"title" db:"title"`
	Students []student `json:"students_list"  xml:"students" db:"students"`
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

	//JSON序列化 Go语言的数据 —》 JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	//JSON反序列化:JSON格式的字符串-》Go语言中的数据
	//jsonStr := `{"title":"火箭101","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"},{"ID":4,"Name":"stu04"},{"ID":5,"Name":"stu05"},{"ID":6,"Name":"stu06"},{D":7,"Name":"stu07"},{"ID":8,"Name":"stu08"},{"ID":9,"Name":"stu09"}]}`
	var c2 class
	err = json.Unmarshal([]byte(data), &c2)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}
