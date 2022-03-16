package main

import "fmt"

func modifyArray(x [3]int) { //定义modifyArray方法
	x[0] = 100
}

func modifyArray2(x [3][2]int) { //定义modifyArray2方法
	x[2][0] = 100
}
func main() {
	//数组的初始化
	var testArray [3]int       //数组会初始化为int类型的零值
	c := [5]int{1, 2, 3, 4, 5} //使用指定的初始值完成初始化
	fmt.Println(testArray, c)  //[0 0 0] [1 2 3 4 5]

	//自行推断数组的长度
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string

	//使用指定索引值的方式来初始化数组
	d := [...]int{1: 1, 3: 5}
	fmt.Println(d)                  //[0 1 0 5]
	fmt.Printf("type of a:%T\n", d) //type of a:[4]int

	//字符数组定义并初始化
	var str [2]string
	str[0] = "world"
	str[1] = "hello"
	fmt.Println(str) //[world hello]

	//数组是值类型
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}
