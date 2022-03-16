package main

import (
	"fmt"
	"reflect"
)

//reflect
func reflectType(x interface{}) {
	//1.通过类型断言
	//2.借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj) //*reflect.rtype
}

type Cat struct{}
type Dog struct{}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T", v, v)
	//得到一个传入时类型的变量
	k := v.Kind() //拿到值对应的类型种类
	fmt.Println(k)
	switch k {
	case reflect.Float32:
		//把反射得到的值转换成一个int32的变量
		ret := float32(v.Float())
		fmt.Printf("%v,%T\n", ret, ret)
	case reflect.Int32:
		//把反射得到的值转换成一个int32的变量
		ret := int32(v.Int())
		fmt.Printf("%v,%T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	//Elem()用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.22)
	}
}
func main() {
	var a float32 = 1.243
	reflectType(a)
	var b int8 = 2
	reflectType(b)
	//结构体类型
	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)
	//slice
	var e []int
	reflectType(e)
	var f []string
	reflectType(f)

	//valueof
	var va int32 = 50
	reflectValue(va)
	var vb float32 = 2.2
	reflectValue(vb)

	//setvalue
	var sa int32 = 200
	reflectSetValue(&sa)
}
