# Go程序的结构及用法

## go mod init

使用go module模式新建项目时，我们需要通过`go mod init 项目名`命令对项目进行初始化，该命令会在项目根目录下生成`go.mod`文件。

```bash
go mod init hello
```

## Hello World

我们使用`hello`作为我们第一个Go项目的名称，执行如下命令。

```bash
go mod init hello
```

接下来在该目录中创建一个[`main.go`](https://github.com/letengzz/JC-GoPro/tree/main/01demo/HelloWorld/main/main.go)文件：

```go
package main  // 声明 main 包，表明当前是一个可执行程序

import "fmt"  // 导入内置 fmt 包

func main(){  // main函数，是程序执行的入口
	fmt.Println("Hello World!")  // 在终端打印 Hello World!
}
```

------

## go build

`go build`表示将源代码编译成可执行文件。

在hello目录下执行：

```bash
go build
```

或者在其他目录执行以下命令：

```bash
go build hello
```

go编译器会去 `GOPATH`的src目录下查找你要编译的`hello`项目

编译得到的可执行文件会保存在执行编译命令的当前目录下，如果是windows平台会在当前目录下找到`hello.exe`可执行文件。

可在终端直接执行该`hello.exe`文件：

```bash
d:\code\go\src\hello>hello.exe
Hello World!
```

我们还可以使用`-o`参数来指定编译后可执行文件的名字。

```bash
go build -o heiheihei.exe
```

------

## go run

`go run main.go`也可以执行程序，该命令本质上也是先编译再执行。

```bash
go run main.go
```

------

## go install

`go install`表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件移动到`GOPATH`的bin目录下。因为我们的环境变量中配置了`GOPATH`下的bin目录，所以我们就可以在任意地方直接执行可执行文件了。

------

## 跨平台编译

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，Go语言支持跨平台编译——在当前平台（例如Windows）下编译其他平台（例如Linux）的可执行文件。

#### Windows编译Linux可执行文件

如果我想在Windows下编译一个Linux下可执行文件，那需要怎么做呢？只需要在编译时指定目标操作系统的平台和处理器架构即可。

> 注意：无论你在Windows电脑上使用VsCode编辑器还是Goland编辑器，都要注意你使用的终端类型，因为不同的终端下命令不一样！！！目前的Windows通常默认使用的是`PowerShell`终端。

如果你的`Windows`使用的是`cmd`，那么按如下方式指定环境变量。

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

如果你的`Windows`使用的是`PowerShell`终端，那么设置环境变量的语法为

```bash
$ENV:CGO_ENABLED=0
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"
```

在你的`Windows`终端下执行完上述命令后，再执行下面的命令，得到的就是能够在Linux平台运行的可执行文件了。

```bash
go build
```

#### Windows编译Mac可执行文件

Windows下编译Mac平台64位可执行程序：

cmd终端下执行：

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```

PowerShell终端下执行：

```bash
$ENV:CGO_ENABLED=0
$ENV:GOOS="darwin"
$ENV:GOARCH="amd64"
go build
```

#### Mac编译Linux可执行文件

Mac电脑编译得到Linux平台64位可执行程序：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

#### Mac编译Windows可执行文件

Mac电脑编译得到Windows平台64位可执行程序：

```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

#### Linux编译Mac可执行文件

Linux平台下编译Mac平台64位可执行程序：

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
```

#### Linux编译Windows可执行文件

Linux平台下编译Windows平台64位可执行程序：

```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

------

# [变量和常量](https://github.com/letengzz/JC-GoPro/tree/main/02Variables_Constants)

## 标识符与关键字

**标识符**

在编程语言中标识符就是程序员定义的具有特殊意义的词，比如**变量名**、**常量名**、**函数名**等等。
Go语言中标识符由**字母数字和`_`(下划线）组成，并且只能以字母和`_`开头**。
举几个例子：`abc`, `_`, `_123`, `a123`。

**关键字**

关键字是指编程语言中预先定义好的具有特殊含义的标识符。

**关键字和保留字都不建议用作变量名**。

Go语言中有25个关键字：

![image-20220213195131533](https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E5%8F%98%E9%87%8F%E5%92%8C%E5%B8%B8%E9%87%8F/%E5%85%B3%E9%94%AE%E5%AD%97.png)

此外，Go语言中还有37个保留字：

![image-20220213195948979](https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E5%8F%98%E9%87%8F%E5%92%8C%E5%B8%B8%E9%87%8F/%E4%BF%9D%E7%95%99%E5%AD%97.png)

## [变量](https://github.com/letengzz/JC-GoPro/blob/main/02Variables_Constants/Variable/Variable.go) 

**变量**(`Variable`)的功能是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见变量的数据类型有：[整型](#整型)、[浮点型](#浮点型)、[布尔型](#布尔型)等。

Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

**变量声明**：Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且**Go语言的变量声明后必须使用**。

**标准声明**

Go语言的变量**声明格式**为：

```go
var 变量名 变量类型
```

变量声明以关键字`var`开头，变量类型放在变量的后面，**行尾无需分号**。
**例**：

```go
var name string
var age int
var isOk bool
```

**批量声明**

每声明一个变量就需要写`var`关键字会比较繁琐，go语言中还支持批量变量声明

**例**：

```go
var (
    a string
    b int
    c bool
    d float32
)
```

**变量的初始化**

Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的**默认值**，例如：

1. 整型和浮点型变量的默认值为`0`。
2. 字符串变量的默认值为`空字符串`。
3. 布尔型变量默认为`false`。
4. 切片、函数、指针变量的默认为`nil`。

当然我们也可在声明变量的时候为其指定初始值。变量初始化的标准格式如下：

```go
var 变量名 类型 = 表达式
```

**例**：

```go
var name string = "nihao"
var age int = 18
```

或者一次初始化多个变量

```go
var name, age = "nihao",20
```

**类型推导**

有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

```go
var name = "nihao";
var age = 18
```

**短变量声明**

在函数内部，可以使用更简略的 `:=` 方式**声明并初始化变量**。

**注意**：`:=`不能使用在函数外。

**例**：

```go
package main

import ( 
	"fmt";
)
// 全局变量m
var m = 100

func main() {
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)
}
```

**匿名变量**

在使用多重赋值时，如果想要忽略某个值，可以使用**匿名变量**(`anonymous variable`)。
匿名变量用一个下划线`_`表示

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。
(在`Lua`等编程语言里，匿名变量也被叫做**哑元变量**。)

**注意**：

1. 函数外的每个语句都必须以关键字开始（`var`、`const`、`func`等）
2. `_`多用于占位，表示忽略值。

**例**：

```go
func foo() (int, string) {
	return 10, "nihao"
}
func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=",x)
	fmt.Println("y=",y)
}
```

## [常量](https://github.com/letengzz/JC-GoPro/blob/main/02Variables_Constants/Constant/Constant.go)

**常量**是恒定不变的值，多用于定义程序运行期间不会改变的那些值。

常量的声明和变量声明非常类似，只是把`var`换成了`const`，常量在定义的时候必须赋值。

```go
const pi = 3.1415
const e = 2.7182
```

声明了`pi`和`e`这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。

多个常量也可以一起声明：

```go
const (
    pi = 3.1415
    e = 2.7182
)
```

`const`同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
**例**：

```go
const (
    n1 = 100
    n2
    n3
  //常量`n1`、`n2`、`n3`的值都是100。
)
```

## [iota](https://github.com/letengzz/JC-GoPro/blob/main/02Variables_Constants/iota/iota.go)

`iota`是go语言的常量计数器，只能在常量的表达式中使用。

`iota`在`const`关键字出现时将被重置为0。const中每新增一行常量声明将使`iota`计数一次(`iota`可理解为`const`语句块中的行索引)。
使用`iota`能简化定义，在定义枚举时很有用。

**例**：

```go
const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
```

使用`_`跳过某些值

```go
const (
		n1 = iota //0
		n2        //1
		_
		n4        //3
	)
```

`iota`声明中间插队

```go
const (
		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)
	const n5 = iota //0
```

定义数量级
（这里的`<<`表示左移操作，`1<<10`表示将1的二进制表示向左移10位，也就是由`1`变成了`10000000000`，也就是十进制的1024。同理`2<<2`表示将2的二进制表示向左移2位，也就是由`10`变成了`1000`，也就是十进制的8。）

```go
const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
```

多个`iota`定义在一行

```go
const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
```

------

# [基本数据类型](https://github.com/letengzz/JC-GoPro/tree/main/03Basic_Data_Type)

## [整型](https://github.com/letengzz/JC-GoPro/blob/main/03Basic_Data_Type/Integer/Integer.go)

**整型**分为以下两个大类： 按长度分为：`int8`、`int16`、`int32`、`int64` 对应的无符号整型：`uint8`、`uint16`、`uint32`、`uint64`

其中，`uint8`就是我们熟知的`byte`型，`int16`对应C语言中的`short`型，`int64`对应C语言中的`long`型。

![image-20220215213812364](https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E5%9F%BA%E6%9C%AC%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/%E6%95%B4%E5%9E%8B/%E6%95%B4%E5%9E%8B.png)

**特殊整型**

![image-20220215213956412](https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E5%9F%BA%E6%9C%AC%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/%E6%95%B4%E5%9E%8B/%E7%89%B9%E6%AE%8A%E6%95%B4%E5%9E%8B.png)

**注意：** 

1. 在使用`int`和 `uint`类型时，不能假定它是32位或64位的整型，而是考虑`int`和`uint`可能在不同平台上的差异。

2. 获取对象的长度的内建`len()`函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用`int`来表示。在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用`int`和 `uint`。

**数字字面量语法（Number literals syntax）**

Go 1.13版本之后引入了数字字面量语法，这样便于开发者以二进制、八进制或十六进制浮点数的格式定义数字。

允许我们用 `_` 来分隔数字，比如说： `v := 123_456` 表示 v 的值等于 123456。

我们可以借助`fmt`函数来将一个整数以不同进制形式展示。

**例**：

`v := 0b00101101`， 代表二进制的 101101，相当于十进制的 45。 `v := 0o377`，代表八进制的 377，相当于十进制的 255。 `v := 0x1p-2`，代表十六进制的 1 除以 2²，也就是 0.25。

```go
package main
 
import "fmt"
 
func main(){
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77
	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF
}
```

## [浮点型](https://github.com/letengzz/JC-GoPro/blob/main/03Basic_Data_Type/Float/Float.go)

Go语言支持两种浮点型数：`float32`和`float64`。这两种浮点型数据格式遵循`IEEE 754`标准： `float32` 的浮点数的最大范围约为 `3.4e38`，可以使用常量定义：`math.MaxFloat32`。 `float64` 的浮点数的最大范围约为 `1.8e308`，可以使用一个常量定义：`math.MaxFloat64`。

打印浮点数时，可以使用`fmt`包配合动词`%f`

**例**：

```go
package main
import (
        "fmt"
        "math"
)
func main() {
        fmt.Printf("%f\n", math.Pi)
        fmt.Printf("%.2f\n", math.Pi)
}
```

## [复数](https://github.com/letengzz/JC-GoPro/blob/main/03Basic_Data_Type/Complex/Complex.go)

复数有实部和虚部，Go语言中分为`complex64`和`complex128`，`complex64`的实部和虚部为32位，`complex128`的实部和虚部为64位。

**例**：

```go
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i
fmt.Println(c1)
fmt.Println(c2)
```

## [布尔值](https://github.com/letengzz/JC-GoPro/blob/main/03Basic_Data_Type/Boolean/Boolean.go)

Go语言中以`bool`类型进行声明布尔型数据，布尔型数据只有`true`(真)和`false`(假)两个值。

**注意：**

1. 布尔类型变量的默认值为`false`。
2. Go 语言中不允许将整型强制转换为布尔型.
3. 布尔型无法参与数值运算，也无法与其他类型进行转换。

## [字符串](https://github.com/letengzz/JC-GoPro/blob/main/03Basic_Data_Type/String/String.go)

Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用`UTF-8`编码。 字符串的值为`双引号(")`中的内容，可以在Go语言的源码中直接添加非ASCII码字符

**例**：

```go
s1 := "hello"
s2 := "你好"
```

[**字符串转义符**](#转义符)

**例**：我们要打印一个Windows平台下的一个文件路径：

```go
package main
import (
    "fmt"
)
func main() {
    fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
}
```

**多行字符串**

Go语言中要定义一个多行字符串时，就必须使用`反引号`字符

反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

**例**：

```go
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```

**字符串的常用操作**

![image-20220218143321249](https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E5%9F%BA%E6%9C%AC%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/%E5%AD%97%E7%AC%A6%E4%B8%B2/%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%9A%84%E5%B8%B8%E7%94%A8%E6%93%8D%E4%BD%9C.png)

## [byte和rune](https://github.com/letengzz/JC-GoPro/blob/main/03Basic_Data_Type/Byte_Rune/main.go)

组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（`'`）包裹起来

**例**：

```go
var a = '中'
var b = 'x'
```

Go 语言的字符有以下两种：

1. `uint8`类型，或者叫 byte 型，代表了`ASCII码`的一个字符。
2. `rune`类型，代表一个 `UTF-8字符`。

当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`。

Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。

```go
// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
```

输出：

```bash
104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³) 
104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河) 
```

因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

字符串底层是一个byte数组，所以可以和`[]byte`类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

**修改字符串**

要修改字符串，需要先将其转换成`[]rune`或`[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。

```go
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
```

## [类型转换](https://github.com/letengzz/JC-GoPro/blob/main/03Basic_Data_Type/Type_Conversion/Type_Conversion.go)

Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

强制类型转换的**基本语法**：

```bash
T(表达式)
```

其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

**例**：计算直角三角形的斜边长时使用`math`包的`Sqrt()函数`，该函数接收的是`float64`类型的参数，而变量a和b都是`int`类型的，这个时候就需要将a和b强制类型转换为`float64`类型。

```go
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
```

# [运算符](https://github.com/letengzz/JC-GoPro/tree/main/04Operator) 

## [算术运算符](https://github.com/letengzz/JC-GoPro/blob/main/04Operator/Arithmetic_operators/Arithmetic_operators.go)

<img src="https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E8%BF%90%E7%AE%97%E7%AC%A6/%E7%AE%97%E6%9C%AF%E8%BF%90%E7%AE%97%E7%AC%A6.png" alt="image-20220218152938412" style="zoom:67%;" />

**注意：** `++`（自增）和`--`（自减）在Go语言中是单独的语句，并不是运算符。

## [关系运算符](https://github.com/letengzz/JC-GoPro/blob/main/04Operator/Relational_operators/Relational_operators.go)

<img src="https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E8%BF%90%E7%AE%97%E7%AC%A6/%E5%85%B3%E7%B3%BB%E8%BF%90%E7%AE%97%E7%AC%A6.png" alt="image-20220218152954880" style="zoom:67%;" />

## [逻辑运算符](https://github.com/letengzz/JC-GoPro/blob/main/04Operator/Logical_operators/Logical_operators.go)

<img src="https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E8%BF%90%E7%AE%97%E7%AC%A6/%E9%80%BB%E8%BE%91%E8%BF%90%E7%AE%97%E7%AC%A6.png" alt="image-20220218153009186" style="zoom:67%;" />

## [位运算符](https://github.com/letengzz/JC-GoPro/blob/main/04Operator/Shift_operator/Shift_operator.go)

位运算符对整数在内存中的二进制位进行操作。

<img src="https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E8%BF%90%E7%AE%97%E7%AC%A6/%E4%BD%8D%E8%BF%90%E7%AE%97%E7%AC%A6.png" alt="image-20220218153132792" style="zoom:67%;" />

## [赋值运算符](https://github.com/letengzz/JC-GoPro/blob/main/04Operator/Assignment_operator/Assignment_operator.go)

<img src="https://cdn.jsdelivr.net/gh/letengzz/Two-C/img/Go/%E8%BF%90%E7%AE%97%E7%AC%A6/%E8%B5%8B%E5%80%BC%E8%BF%90%E7%AE%97%E7%AC%A6.png" alt="image-20220218153150658" style="zoom:67%;" />

# [流程控制](https://github.com/letengzz/JC-GoPro/tree/main/05Process_Control)

## [分支结构](https://github.com/letengzz/JC-GoPro/tree/main/05Process_Control/Branch_Structure)

### [if else语句](https://github.com/letengzz/JC-GoPro/blob/main/05Process_Control/Branch_Structure/if_else/if_else.go)

**基本写法**

```go
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
```

当表达式1的结果为`true`时，执行分支1，否则判断表达式2，如果满足则执行分支2，都不满足时，则执行分支3。 if判断中的`else if`和`else`都是可选的，可以根据实际需要进行选择。

**注意**：Go语言规定与`if`匹配的左括号`{`必须与`if和表达式`放在同一行，`{`放在其他位置会触发编译错误。

同理，与`else`匹配的`{`也必须与`else`写在同一行，`else`也必须与上一个`if`或`else if`右边的大括号在同一行。

**例**：

```go
func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

**特殊写法**

if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断

**例**：

```go
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

------

### [switch case](https://github.com/letengzz/JC-GoPro/blob/main/05Process_Control/Branch_Structure/switch_case/switch_case.go)

使用`switch`语句可方便地对大量的值进行条件判断。

```go
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
```

Go语言规定每个`switch`只能有一个`default`分支。

一个分支可以有多个值，多个`case`值中间使用英文逗号分隔。

```go
func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}
```

分支还可以使用表达式，这时候`switch`语句后面不需要再跟判断变量。

**例**：

```go
func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
```

`fallthrough`语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。

```go
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
```

输出：

```bash
a
b
```

## [循环结构](https://github.com/letengzz/JC-GoPro/tree/main/05Process_Control/Loop_Structure)

### [for循环](https://github.com/letengzz/JC-GoPro/blob/main/05Process_Control/Loop_Structure/for/for.go)

Go 语言中的所有循环类型均可以使用`for`关键字来完成。

**基本格式**：

```bash
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

条件表达式返回`true`时循环体不停地进行循环，直到条件表达式返回`false`时自动退出循环。

```go
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句可以被忽略，但是初始语句后的分号必须要写

**例**：

```go
func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}
```

for循环的初始语句和结束语句都可以省略

**例**：

```go
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
```

这种写法类似于其他编程语言中的`while`，在`while`后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。

**无限循环**

```go
for {
    循环体语句
}
```

for循环可以通过[`break`](#break跳出循环)、[`goto`](#goto跳转到指定标签)、[`return`](#)、[`panic`](#)语句强制退出循环。

------

### [for range(键值循环)](https://github.com/letengzz/JC-GoPro/blob/main/05Process_Control/Loop_Structure/for_range/for_range.go)

Go语言中可以使用`for range`遍历数组、切片、字符串、map 及通道（channel）。 通过`for range`遍历的返回值有以下规律：

1. [数组](#数组)、[切片](#切片)、[字符串](#字符串)返回索引和值。
2. [map](#map)返回键和值。
3. [通道](#)`channel`只返回通道内的值。

```go
var a = [4]int{1, 2, 3, 5}
	for index, value := range a {
      //index索引 value数组的值
		fmt.Println("第", i, "的数是", i2)
	}
```

------

## [跳转语句](https://github.com/letengzz/JC-GoPro/tree/main/05Process_Control/Jump_Statements)

### [goto(跳转到指定标签)](https://github.com/letengzz/JC-GoPro/blob/main/05Process_Control/Jump_Statements/goto/goto.go)

`goto`语句通过标签进行代码间的无条件跳转。`goto`语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用`goto`语句能简化一些代码的实现过程。 

**例**：双层嵌套的for循环要退出时：

```go
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}
```

使用`goto`语句能简化代码：

```go
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
```

------

### [break(跳出循环)](https://github.com/letengzz/JC-GoPro/blob/main/05Process_Control/Jump_Statements/break/break.go)

`break`语句可以结束`for`、`switch`和`select`的代码块。

`break`语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的`for`、`switch`和 `select`的代码块上。 

**例**：

```go
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
```

### [continue(继续下次循环)](https://github.com/letengzz/JC-GoPro/blob/main/05Process_Control/Jump_Statements/continue/continue.go)

`continue`语句可以结束当前循环，开始下一次的循环迭代过程，仅限在`for`循环内使用。

在 `continue`语句后添加标签时，表示开始标签对应的循环。

**例**：

```go
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
```

# [数组](https://github.com/letengzz/JC-GoPro/tree/main/06Arrays)

数组是**同一种数据类型元素的集合**。 

在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是**数组大小不可变化**。

## [数组的定义](https://github.com/letengzz/JC-GoPro/blob/main/06Arrays/Array_Define/ArrayDefine.go)

**定义方式**：

```bash
var 数组变量名 [元素数量]T
```

**例**：

```go
var a [3]int
var b [4]int
a = b //不可以这样做，因为此时a和b是不同的类型
```

**注意**：

1. 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 `[5]int`和`[10]int`是不同的类型。
2. 数组可以通过下标进行访问，下标是从`0`开始，最后一个元素下标是：`len-1`，访问越界（下标在合法范围之外），则触发访问越界，会`panic`。

## [数组的初始化](https://github.com/letengzz/JC-GoPro/blob/main/06Arrays/Array_Initialization/Array_Initialization.go)

1. 初始化数组时可以使用初始化列表来设置数组元素的值。

   **注意**：每次都要确保提供的初始值和数组长度一致。

```go
func main() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}
```

2.　可以让编译器根据初始值的个数**自行推断数组的长度**

**例**：

```go
func main() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}
```

3.　使用指定索引值的方式来初始化数组

**例**:

```go
func main() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}
```

## [数组的遍历](https://github.com/letengzz/JC-GoPro/blob/main/06Arrays/Array_Traversal/Array_Traversal.go)

1. `for`循环遍历

   ```go
   var a = [...]string{"北京", "上海", "深圳"}
   for i := 0; i < len(a); i++ {
   	fmt.Println(a[i])
   }    
   ```

2. `for range`遍历

   ```go
   for index, value := range a {
   	fmt.Println(index, value)
   }
   ```

## [多维数组](https://github.com/letengzz/JC-GoPro/blob/main/06Arrays/Multidimensional_Array/main.go)

### [二维数组的定义](https://github.com/letengzz/JC-GoPro/blob/main/06Arrays/Multidimensional_Array/Multidimensional_Array_Initialization.go)

```go
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}
```

### [二维数组的遍历](https://github.com/letengzz/JC-GoPro/blob/main/06Arrays/Multidimensional_Array/Multidimensional_Array_Traversal.go)

```go
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
```

输出：

```bash
北京	上海	
广州	深圳	
成都	重庆	
```

**注意：** 多维数组**只有第一层**可以使用`...`来让编译器推导数组长度。

**例**：

```go
//支持的写法
a := [...][2]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
//不支持多维数组的内层使用...
b := [3][...]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
```

## [$指针数组](https://github.com/letengzz/JC-GoPro/blob/main/06Arrays/Pointer_Array/Pointer_Array.go)

## 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```go
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
func main() {
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
```

**注意：**

1. 数组支持"`==`"、"`!=`" 操作符，因为内存总是被初始化过的。
2. `[n]*T`表示[指针数组](#指针数组)，`*[n]T`表示[数组指针](#数组指针) 。

# [切片](https://github.com/letengzz/JC-GoPro/tree/main/07Slice)

切片（`Slice`）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

## [切片的定义](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Define/Slice_Define.go)

声明切片类型的**基本语法**如下：

```go
var 变量名 []切片中的元素类型
```

**例**：

```go
func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}
```

## [切片表达式](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Define/Slice_Define.go)

​	切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。

它有两种变体：

1. 一种指定low和high两个索引界限值的[简单的形式](#简单切片表达式)
2. 除了low和high索引界限值外还指定容量的[完整的形式](#完整切片表达式)

### 简单切片表达式

切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 切片表达式中的`low`和`high`表示一个索引范围（左包含，右不包含），也就是下面代码中从数组a中选出`1<=索引值<4`的元素组成切片s，得到的切片`长度=high-low`，容量等于得到的切片的底层数组的容量。

```go
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))//s:[2 3] len(s):2 cap(s):4
}
```

为了方便起见，可以省略切片表达式中的任何索引。

省略了`low`则默认为0；省略了`high`则默认为切片操作数的长度:

```go
a[2:]  // 等同于 a[2:len(a)]
a[:3]  // 等同于 a[0:3]
a[:]   // 等同于 a[0:len(a)]
```

**注意：**

对于数组或字符串，如果`0 <= low <= high <= len(a)`，则索引合法，否则就会索引越界（`out of range`）。

对切片再执行切片表达式时（切片再切片），`high`的上限边界是切片的容量`cap(a)`，而不是长度。**常量索引**必须是非负的，并且可以用int类型的值表示;对于数组或常量字符串，常量索引也必须在有效范围内。如果`low`和`high`两个指标都是常数，它们必须满足`low <= high`。如果索引在运行时超出范围，就会发生运行时`panic`。

```go
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))//s:[2 3] len(s):2 cap(s):4
	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))//s2:[5] len(s2):1 cap(s2):1
}
```

### 完整切片表达式

对于数组，指向数组的指针，或切片a(**注意不能是字符串**)支持完整切片表达式：

```go
a[low : high : max]
```

上面的代码会构造与简单切片表达式`a[low: high]`相同类型、相同长度和元素的切片。另外，它会将得到的结果切片的容量设置为`max-low`。在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。

```go
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))//t:[2 3] len(t):2 cap(t):4
}
```

完整切片表达式需要满足的条件是`0 <= low <= high <= max <= cap(a)`，其他条件和简单切片表达式相同。

## [构造切片](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Define/Slice_Define.go)

我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的`make()`函数，格式如下：

```bash
make([]元素类型, 元素的数量, 容量)
```

**例**：

```go
func main() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}
```

上面代码中`a`的内部存储空间已经分配了10个，但实际上只用了2个。 容量并不会影响当前元素的个数，所以`len(a)`返回2，`cap(a)`则返回该切片的容量。

## [切片的长度和容量](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Length_Capacity/Slice_Length_Capacity.go)

切片拥有自己的长度和容量，我们可以通过使用内置的`len()`函数求长度，使用内置的`cap()`函数求切片的容量。

## [判断切片是否为空](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Judge_Compare/Slice_Judge_Compare.go)

要检查切片是否为空，请始终使用`len(s) == 0`来判断，而不应该使用`s == nil`来判断。

## [切片不能直接比较](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Judge_Compare/Slice_Judge_Compare.go)

切片之间是不能比较的，我们不能使用`==`操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和`nil`比较。 **一个`nil`值的切片并没有底层数组**

**注意**：

一个`nil`值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是`nil`

**例**：

```go
var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
```

所以要判断一个切片是否是空的，要是用`len(s) == 0`来判断，不应该使用`s == nil`来判断。

## [切片的赋值拷贝](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_CopyAssignment/Slice_CopyAssignment.go)

下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。

```go
func main() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}
```

## [切片遍历](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Traversal/Slice_Traversal.go)

切片的遍历方式和数组是一致的，支持索引遍历和`for range`遍历。

```go
func main() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
```

## [为切片添加元素](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Append/Slice_Append.go)

Go语言的内建函数`append()`可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加`…`）。

```go
func main(){
	var s []int
	s = append(s, 1)        // [1]
	s = append(s, 2, 3, 4)  // [1 2 3 4]
	s2 := []int{5, 6, 7}  
	s = append(s, s2...)    // [1 2 3 4 5 6 7]
}
```

**注意：**通过var声明的零值切片可以在`append()`函数直接使用，无需初始化。

```go
var s []int
s = append(s, 1, 2, 3)
```

没有必要像下面的代码一样初始化一个切片再传入`append()`函数使用，

```go
s := []int{}  // 没有必要初始化
s = append(s, 1, 2, 3)

var s = make([]int)  // 没有必要初始化
s = append(s, 1, 2, 3)
```

每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在`append()`函数调用时，所以我们通常都需要用原变量接收append函数的返回值。

**例**：

```go
func main() {
	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}
```

输出：

```bash
[0]  len:1  cap:1  ptr:0xc0000a8000
[0 1]  len:2  cap:2  ptr:0xc0000a8040
[0 1 2]  len:3  cap:4  ptr:0xc0000b2020
[0 1 2 3]  len:4  cap:4  ptr:0xc0000b2020
[0 1 2 3 4]  len:5  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5]  len:6  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6]  len:7  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6 7]  len:8  cap:8  ptr:0xc0000b6000
[0 1 2 3 4 5 6 7 8]  len:9  cap:16  ptr:0xc0000b8000
[0 1 2 3 4 5 6 7 8 9]  len:10  cap:16  ptr:0xc0000b8000
```

从上面的结果可以看出：

1. `append()`函数将元素追加到切片的最后并返回该切片。
2. 切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。

**append()函数还支持一次性追加多个元素** 

**例**：

```go
var citySlice []string
// 追加一个元素
citySlice = append(citySlice, "北京")
// 追加多个元素
citySlice = append(citySlice, "上海", "广州", "深圳")
// 追加切片
a := []string{"成都", "重庆"}
citySlice = append(citySlice, a...)
fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
```

## 切片的扩容策略

可以通过查看`$GOROOT/src/runtime/slice.go`源码，其中扩容相关代码如下：

```go
newcap := old.cap
doublecap := newcap + newcap
if cap > doublecap {
	newcap = cap
} else {
	if old.len < 1024 {
		newcap = doublecap
	} else {
		// Check 0 < newcap to detect overflow
		// and prevent an infinite loop.
		for 0 < newcap && newcap < cap {
			newcap += newcap / 4
		}
		// Set newcap to the requested cap when
		// the newcap calculation overflowed.
		if newcap <= 0 {
			newcap = cap
		}
	}
}
```

从上面的代码可以看出以下内容：

- 首先判断，如果新申请容量（`cap`）大于2倍的旧容量（`old.cap`），最终容量（`newcap`）就是新申请的容量（`cap`）。
- 否则判断，如果旧切片的长度小于1024，则最终容量(`newcap`)就是旧容量(`old.cap`)的两倍，即（`newcap=doublecap`），
- 否则判断，如果旧切片长度大于等于1024，则最终容量（`newcap`）从旧容量（`old.cap`）开始循环增加原来的1/4，即（`newcap=old.cap,for {newcap += newcap/4}`）直到最终容量（`newcap`）大于等于新申请的容量(`cap`)，即（`newcap >= cap`）
- 如果最终容量（`cap`）计算值溢出，则最终容量（`cap`）就是新申请容量（`cap`）。

**注意**：切片扩容还会根据切片中元素的类型不同而做不同的处理，比如`int`和`string`类型的处理方式就不一样。

## [复制切片](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Copy/Slice_Copy.go)

```go
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]
}
```

由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。

Go语言内建的`copy()`函数可以迅速地将一个切片的数据复制到另外一个切片空间中，`copy()`函数的使用格式如下：

```bash
copy(目标切片, 数据来源切片 []T)
```

**例**：

```go
func main() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}
```

## [从切片中删除元素](https://github.com/letengzz/JC-GoPro/blob/main/07Slice/Slice_Delete/Slice_Delete.go)

Go语言并没有提供用于删除元素的语法或接口，而是通过利用切片本身的特性来删除元素——追加元素。即 **以被删除元素为分界点，将前后两个部分的内存重新连接起来。**使用切片的追加(append)特性，利用代码实现

```go
func main() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
```

**总结**：要从切片a中删除索引为`index`的元素，操作方法是`a = append(a[:index], a[index+1:]...)`

## 切片原理

切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

举个例子，现在有一个数组`a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}`，切片`s1 := a[:5]`，相应示意图如下。![slice_01](https://www.liwenzhou.com/images/Go/slice/slice_01.png)切片`s2 := a[3:6]`，相应示意图如下：![slice_02](https://www.liwenzhou.com/images/Go/slice/slice_02.png)



# 附录

## 转义符

![image-20220223115507187](C:\Users\LetengZzz\AppData\Roaming\Typora\typora-user-images\image-20220223115507187.png)
