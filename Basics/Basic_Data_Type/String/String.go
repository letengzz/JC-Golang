package main

import (
	"fmt"
	"strings"
)

/*
 Go 语言里的字符串的内部实现使用`UTF-8`编码。
字符串的值为`双引号(")`中的内容，可以在Go语言的源码中直接添加非ASCII码字符
*/
func main() {
	//字符串的定义及初始化
	var s1 string = "Hello World"
	var s2 = "ddd"
	s3 := "你好\n"
	fmt.Printf("%s %s %s", s1, s2, s3)                  //Hello World ddd 你好
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"") //str := "c:\Code\lesson1\go.exe"

	//多行字符串
	/*反引号间换行将被作为字符串中的换行
	但是所有的转义字符均无效，文本将会原样输出。*/
	s4 := `第一行
 第二行
第三行`
	fmt.Println(s4) /*
		第一行
		 第二行
		第三行
	*/

	//字符串的常用操作
	ss1 := "Hello Lilei"
	var ss2 = "nihao"

	//求长度
	fmt.Println("s4的长度为", len(ss1)) //s4的长度为 11

	//拼接字符串
	fmt.Println(ss1 + ss2)                //Hello Lileinihao
	ss3 := fmt.Sprintf("%s+%s", ss1, ss2) //Hello Lilei+nihao
	fmt.Println(ss3)

	//分割
	fmt.Println("分割Li", strings.Split(ss1, "Li")) //分割Li [Hello  lei]

	//判断是否包含
	fmt.Println("是否包含H:", strings.Contains(ss1, "H")) //是否包含H: true

	//前后缀判断
	fmt.Println("是否前缀为H:", strings.HasPrefix(ss1, "H")) //是否前缀为H: true
	fmt.Println("是否后缀为o:", strings.HasSuffix(ss1, "o")) //是否后缀为o: false

	//串出现的位置
	fmt.Println("首次出现的位置:", strings.Index(ss1, "H"))     //首次出现的位置: 0
	fmt.Println("最后出现的位置:", strings.LastIndex(ss1, "l")) //最后出现的位置: 8

	//Jion操作
	s := []string{"how", "do", "you", "do"}
	fmt.Println("Jion操作:", strings.Join(s, "+")) //Jion操作: how+do+you+do

	//遍历字符串
	st := "hello你好"
	for i := 0; i < len(st); i++ { //byte
		fmt.Printf("%v(%c) ", st[i], st[i]) //104(h) 101(e) 108(l) 108(l) 111(o) 228(ä) 189(½) 160( ) 229(å) 165(¥) 189(½)
	}
	fmt.Println()
	for _, r := range st { //rune
		fmt.Printf("%v(%c) ", r, r) //104(h) 101(e) 108(l) 108(l) 111(o) 20320(你) 22909(好)
	}
	fmt.Println()

	//修改字符串
	//无论哪种转换，都会重新分配内存，并复制字节数组。
	ss := "big"
	// 强制类型转换
	byteS1 := []byte(ss)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1)) //pig

	sn := "白萝卜"
	runeS2 := []rune(sn)
	runeS2[0] = '红'
	fmt.Println(string(runeS2)) //红萝卜
}
