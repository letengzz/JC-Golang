# Go程序的结构及用法

## Hello World

现在我们来创建第一个Go项目——`hello`。在我们的`GOPATH`下的src目录中创建hello目录。

在该目录中创建一个`main.go`文件：

```go
package main  // 声明 main 包，表明当前是一个可执行程序

import "fmt"  // 导入内置 fmt 包

func main(){  // main函数，是程序执行的入口
	fmt.Println("Hello World!")  // 在终端打印 Hello World!
}
```

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

## go install

`go install`表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件移动到`GOPATH`的bin目录下。因为我们的环境变量中配置了`GOPATH`下的bin目录，所以我们就可以在任意地方直接执行可执行文件了。

## 跨平台编译

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，如果我想在windows下编译一个linux下可执行文件，那需要怎么做呢？

只需要指定目标操作系统的平台和处理器架构即可：

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

然后再执行`go build`命令，得到的就是能够在Linux平台运行的可执行文件了。

Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Linux 下编译 Mac 和 Windows 平台64位可执行程序：

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Windows下编译Mac平台64位可执行程序：

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```

# 变量和常量

## 标识符与关键字

**标识符**

在编程语言中标识符就是程序员定义的具有特殊意义的词，比如**变量名**、**常量名**、**函数名**等等。
Go语言中标识符由**字母数字和`_`(下划线）组成，并且只能以字母和`_`开头**。
举几个例子：`abc`, `_`, `_123`, `a123`。

**关键字**

关键字是指编程语言中预先定义好的具有特殊含义的标识符。

**关键字和保留字都不建议用作变量名**。

Go语言中有25个关键字：

![image-20220213195131533](C:\Users\LetengZzz\AppData\Roaming\Typora\typora-user-images\image-20220213195131533.png)

此外，Go语言中还有37个保留字：

![image-20220213195948979](C:\Users\LetengZzz\AppData\Roaming\Typora\typora-user-images\image-20220213195948979.png)

## 变量 

变量(**Variable**)的功能是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见变量的数据类型有：[整型](#基本数据类型/整型)、[浮点型](#基本数据类型/浮点型)、[布尔型](#基本数据类型/布尔型)等。

Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

**变量声明**：Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且**Go语言的变量声明后必须使用**。

## 标准声明

Go语言的变量**声明格式**为：

```go
var 变量名 变量类型
```

变量声明以关键字`var`开头，变量类型放在变量的后面，**行尾无需分号**。
举个例子：

```go
var name string
var age int
var isOk bool
```

## 批量声明

每声明一个变量就需要写`var`关键字会比较繁琐，go语言中还支持批量变量声明：

```go
var (
    a string
    b int
    c bool
    d float32
)
```

## 变量的初始化

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
var name string = "nihao";
var age int = 18
```

或者一次初始化多个变量

```go
var name, age = "nihao",20
```

### 类型推导

有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

```go
var name = "nihao";
var age = 18
```

### 短变量声明

在函数内部，可以使用更简略的 `:=` 方式**声明并初始化变量**。

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

**注意**：`:=`不能使用在函数外。

###  匿名变量

在使用多重赋值时，如果想要忽略某个值，可以使用**匿名变量**(`anonymous variable`)。
匿名变量用一个下划线`_`表示，例如：

```go
func foo() (int, string) {
	return 10, "Q1mi"
}
func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=",x)
	fmt.Println("y=",y)
}
```

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。
(在`Lua`等编程语言里，匿名变量也被叫做**哑元变量**。)

**注意**：

1. 函数外的每个语句都必须以关键字开始（`var`、`const`、`func`等）
2. `_`多用于占位，表示忽略值。

## 常量

常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。

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

const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
例如：

```go
const (
    n1 = 100
    n2
    n3
)
```

上面示例中，常量`n1`、`n2`、`n3`的值都是100。

## iota

`iota`是go语言的常量计数器，只能在常量的表达式中使用。

`iota`在const关键字出现时将被重置为0。const中每新增一行常量声明将使`iota`计数一次(iota可理解为const语句块中的行索引)。
使用iota能简化定义，在定义枚举时很有用。

举个例子：

```go
const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
```

## 9.1 几个常见的iota示例:

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

# 基本数据类型

## 整型

## 浮点型

## 复数

## 布尔值

## 字符串

## byte和rune

## 类型转换

# 运算符 

## 算术运算符

## 关系运算符 

## 逻辑运算符

## 位运算符

## 赋值运算符



# 流程控制

## 分支结构

## 循环结构

## 跳转语句

# 数组

## 数组的定义

## 数组的初始化

## 数组的遍历

## 多维数组



# 切片

## 切片的定义

## 切片的赋值拷贝

## 切片遍历

## 为切片添加元素

## 扩容策略

## 使用函数复制切片

## 从切片中删除元素



# 指针

# map

# 函数

# 结构体

# 包

# 接口

# 反射

# 并发

# 网络编程

# 单元测试

# 扩展

## 常用标准库

### fmt

### time

### flag

### log

### 文件操作

### strconv

### template

### net_http

### context

## 框架

### Gin

# 附录 

## Go的安装及搭建开发环境

### 下载地址

Go官网下载地址：https://golang.org/dl/

Go官方镜像站（推荐）：https://golang.google.cn/dl/

### 版本的选择 

Windows平台和Mac平台推荐下载可执行文件版，Linux平台下载压缩文件版。

![image-20220213175135834](C:\Users\LetengZzz\AppData\Roaming\Typora\typora-user-images\image-20220213175135834.png)

#### Windows安装

此安装实例以 `64位Win10`系统安装 `Go1.11.5可执行文件版本`为例。

将上一步选好的安装包下载到本地。

<img src="C:\Users\LetengZzz\AppData\Roaming\Typora\typora-user-images\image-20220213175344630.png" alt="image-20220213175344630" style="zoom:67%;" />

双击下载好的文件

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550236819662.png)

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550236972659.png)

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550237077339.png)

#### Linux下安装

我们在版本选择页面选择并下载好`go1.11.5.linux-amd64.tar.gz`文件：

```bash
wget https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz
```

将下载好的文件解压到`/usr/local`目录下：

```bash
mkdir -p /usr/local/go  # 创建目录
tar -C /usr/lcoal/go zxvf go1.11.5.linux-amd64.tar.gz. # 解压
```

如果提示没有权限，加上`sudo`以root用户的身份再运行。执行完就可以在`/usr/local/`下看到go目录了。

配置环境变量：
Linux下有两个文件可以配置环境变量，其中`/etc/profile`是对所有用户生效的；`$HOME/.profile`是对当前用户生效的，根据自己的情况自行选择一个文件打开，添加如下两行代码，保存退出。

```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

修改`/etc/profile`后要重启生效，修改`$HOME/.profile`后使用source命令加载`$HOME/.profile`文件即可生效。
检查：

```bash
~ go version
go version go1.11.5 linux/amd64
```

#### Mac下安装

下载可执行文件版，直接点击**下一步**安装即可，默认会将go安装到`/usr/local/go`目录下。
![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/mac_install_go.png)

### 检查

上一步安装过程执行完毕后，可以打开终端窗口，输入`go version`命令，查看安装的Go版本。

```shell
go version
```

 ![image-20220213163515991](C:\Users\LetengZzz\AppData\Roaming\Typora\typora-user-images\image-20220213163515991.png)

### 配置GOPATH

`GOPATH`是一个环境变量，用来表明你写的go项目的存放路径（工作目录）。

`GOPATH`路径最好只设置一个，所有的项目代码都放到`GOPATH`的`src`目录下。

Linux和Mac平台就参照上面配置环境变量的方式将自己的工作目录添加到环境变量中即可。
Windows平台按下面的步骤将`D:\code\go`添加到环境变量：

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550293810242.png)

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550293634570.png)

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550293854247.png)

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550294002369.png)

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550294111480.png)

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550294152856.png)

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550296325263.png)

在 Go 1.8 版本之前，`GOPATH`环境变量默认是空的。从 Go 1.8 版本开始，Go 开发包在安装完成后会为 `GOPATH`设置一个默认目录，参见下表。

**GOPATH在不同操作系统平台上的默认值**

|  平台   |   GOPATH默认值   |        举例        |
| :-----: | :--------------: | :----------------: |
| Windows | %USERPROFILE%/go | C:\Users\用户名\go |
|  Unix   |     $HOME/go     |  /home/用户名/go   |

同时，我们将 `GOROOT`下的bin目录及`GOPATH`下的bin目录都添加到环境变量(`path`)中。

### 查看终端运行环境

在`cmd`中，输入`go env`

```go
go  env
```

## Go项目结构

在进行Go语言开发的时候，我们的代码总是会保存在`$GOPATH/src`目录下。在工程经过`go build`、`go install`或`go get`等指令后，会将下载的第三方包源代码文件放在`$GOPATH/src`目录下， 产生的二进制可执行文件放在 `$GOPATH/bin`目录下，生成的中间缓存文件会被保存在 `$GOPATH/pkg` 下。

如果我们使用版本管理工具（`Version Control System`，`VCS`。常用如`Git`）来管理我们的项目代码时，我们只需要添加`$GOPATH/src`目录的源代码即可。`bin` 和 `pkg` 目录的内容无需版本控制。

### 5.1 适合个人开发者

我们知道源代码都是存放在`GOPATH`的`src`目录下，那我们可以按照下图来组织我们的代码。

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550805203054.png)

### 5.2 目前流行的项目结构

Go语言中也是通过包来组织代码文件，我们可以引用别人的包也可以发布自己的包，但是为了防止不同包的项目名冲突，我们通常使用`顶级域名`来作为包名的前缀，这样就不担心项目名冲突的问题了。

因为不是每个个人开发者都拥有自己的顶级域名，所以目前流行的方式是使用个人的github用户名来区分不同的包。

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550805044488.png)

举个例子：张三和李四都有一个名叫`studygo`的项目，那么这两个包的路径就会是：

```go
import "github.com/zhangsan/studygo";
```

和

```go
import "github.com/lisi/studygo";
```

以后我们从github上下载别人包的时候，如：

```bash
go get github.com/jmoiron/sqlx
```

那么，这个包会下载到我们本地`GOPATH`目录下的`src/github.com/jmoiron/sqlx`。

### 5.3 适合企业开发者

![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550806101915.png)

# Go开发编辑器

Go采用的是UTF-8编码的文本文件存放源代码，理论上使用任何一款文本编辑器都可以做Go语言开发，这里推荐使用[`VS Code`]()和[`Goland`]()。
`VS Code`是微软开源的编辑器，而`Goland`是`jetbrains`出品的付费IDE。

## Goland



## Visual Studio Code

`VS Code`全称`Visual Studio Code`，是微软公司开源的一款**免费**现代化轻量级代码编辑器，支持几乎所有主流的开发语言的语法高亮、智能代码补全、自定义热键、括号匹配、代码片段、代码对比 Diff、GIT 等特性，支持插件扩展，支持 Win、Mac 以及 Linux平台。

虽然不如某些IDE功能强大，但是它添加Go扩展插件后已经足够胜任我们日常的Go开发。

## 6.2 下载与安装

`VS Code`官方下载地址：https://code.visualstudio.com/Download

三大主流平台都支持，请根据自己的电脑平台选择对应的安装包。
![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550239338474.png?x-oss-process=style/watermark)
双击下载好的安装文件，双击安装即可。

## 6.3 配置

### 6.3.1 安装中文简体插件

点击左侧菜单栏最后一项`管理扩展`，在`搜索框`中输入`chinese` ，选中结果列表第一项，点击`install`安装。

安装完毕后右下角会提示`重启VS Code`，重启之后你的VS Code就显示中文啦！
![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/vscode1.gif?x-oss-process=style/watermark)
`VSCode`主界面介绍：
![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550240342443.png?x-oss-process=style/watermark)

### 6.3.2 安装go扩展

现在我们要为我们的VS Code编辑器安装`Go`扩展插件，让它支持Go语言开发。
![img](https://imgmd.oss-cn-shanghai.aliyuncs.com/go/1550241460281.png?x-oss-process=style/watermark)

