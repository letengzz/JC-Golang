package main

import "fmt"

func main() {
	// 声明切片类型
	//var name []T
	var a []string              //声明一个字符串切片
	var b = []int{1, 2, 3, 4}   //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[1 2 3 4]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	//基于数组得到切片
	var intArr = [5]int{1, 2, 3, 4}

	//简单切片表达式
	//从数组a中选出`1<=索引值<4`的元素组成切片s，
	//得到的切片`长度=high-low`
	//容量等于得到的切片的底层数组的容量-low
	intSlice := intArr[1:3]           // s := a[low:high]
	fmt.Println("容量:", cap(intSlice)) //容量: 5-1 = 4
	fmt.Printf("%#v\n", intSlice)     //[]int{2, 3}

	//完整切片表达式:
	//容量设置为`max-low`
	//在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0
	intSlice2 := intArr[0:1:3]               //a[low : high : max]
	fmt.Println("容量:", cap(intSlice2))       //容量: 3-0 = 3
	fmt.Printf("intSlice2:%#v\n", intSlice2) //intSlice2:[]int{1}

	//切片再次切片
	//再次切片后的切片容量为其所切切片的容量-low
	intSlice3 := intSlice[0:4]         //索引的上限是cap(s)而不是len(s)
	fmt.Println("容量:", cap(intSlice3)) //容量: 4 - 0 = 4
	fmt.Printf("%#v\n", intSlice3)     //[]int{2, 3, 4, 0}

	//通过make函数构造切片
	d := make([]int, 3, 10)
	fmt.Println(d)        //[0 0 0]
	fmt.Printf("%T\n", d) //[]int

}
