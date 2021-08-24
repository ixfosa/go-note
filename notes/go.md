## 安装与运行环境

### Go 环境变量

Go 将被默认安装在目录 `c:/go` 下。

- **`$GOROOT`** 表示 Go 在你的电脑上的安装位置，它的值一般都是 `$HOME/go`，也可以安装在别的地方。
- **`$GOARCH`** 表示目标机器的处理器架构，它的值可以是 386、amd64 或 arm。
- **`$GOOS`** 表示目标机器的操作系统，它的值可以是 darwin、freebsd、linux 或 windows。
- **`$GOBIN`** 表示编译器和链接器的安装位置，默认是 `$GOROOT/bin`，如果你使用的是 Go 1.0.3 及以后的版本，一般情况下你可以将它的值设置为空，Go 将会使用前面提到的默认值。

> 目标机器是指你打算运行你的 Go 应用程序的机器。



Go 编译器支持`交叉编译`，也就是说你可以在一台机器上构建运行在具有不同操作系统和处理器架构上运行的应用程序，也就是说编写源代码的机器可以和目标机器有完全不同的特性（操作系统与处理器架构）。

为了区分本地机器和目标机器，你可以使用 `$GOHOSTOS` 和 `$GOHOSTARCH` 设置本地机器的操作系统名称和编译体系结构，这两个变量只有在进行交叉编译的时候才会用到，如果你不进行显示设置，他们的值会和本地机器（`$GOOS` 和 `$GOARCH`）一样。

- `$GOPATH`默认采用和 `$GOROOT` 一样的值，但从 Go 1.1 版本开始，你必须修改为其它路径。它可以包含多个 Go 语言源码文件、包文件和可执行文件的路径，而这些路径下又必须分别包含三个规定的目录：`src`、`pkg` 和 `bin`，这三个目录分别用于存放源码文件、包文件和可执行文件。
- **$GOARM** 专门针对基于 arm 架构的处理器，它的值可以是 5 或 6，默认为 6。
- **$GOMAXPROCS** 用于设置应用程序可使用的处理器个数与核数

```go
go env //打印Go所有默认环境变量
go env GOPATH //打印某个环境变量的值
```



### Windows上安装

下载安装包：`go1.9.7.windows-amd64.zip`

环境变量配置：

```go
GOFROOT  D:\env\go  //表示Go语言的安装目录。
Path     %GOFROOT%\bin  
GOPATH   D:\code\GoProjects  //开发工作区,是存放源代码、测试文件、库静态文件、可执行文件的工作。
```

> 注意，`GOPATH`的值不能与`GOROOT`相同。

### 安装目录清单

你的 Go 安装目录（` D:\env\go`）的文件夹结构应该如下所示：

README.md, AUTHORS, CONTRIBUTORS, LICENSE

- `/bin`：包含可执行文件，如：编译器，Go 工具
- `/doc`：包含示例程序，代码工具，本地文档等
- `/lib`：包含文档模版
- `/misc`：包含与支持 Go 编辑器有关的配置文件以及 cgo 的示例
- `/os_arch`：包含标准库的包的对象文件（`.a`）
- `/src`：包含源代码构建脚本和标准库的包的完整源代码（Go 是一门开源语言）
- `/src/cmd`：包含 Go 和 C 的编译器和命令行脚本

### Go 运行时（runtime）

尽管 Go 编译器产生的是本地可执行代码，这些代码仍旧运行在 Go 的 runtime（这部分的代码可以在 runtime 包中找到）当中。这个 `runtime` 类似 Java 和 .NET 语言所用到的`虚拟机`，它负责管理包括内存分配、垃圾回收、栈处理、goroutine、channel、切片（slice）、map 和反射（reflection）等等。

runtime 主要由 C 语言编写（Go 1.5 开始自举），并且是每个 Go 包的最顶级包。你可以在目录 `D:\env\go\src\runtime`中找到相关内容。

`垃圾回收器` Go 拥有简单却高效的标记-清除回收器。它的主要思想来源于 IBM 的可复用垃圾回收器，旨在打造一个高效、低延迟的并发回收器。目前 gccgo 还没有回收器，同时适用 gc 和 gccgo 的新回收器正在研发中。使用一门具有垃圾回收功能的编程语言不代表你可以避免内存分配所带来的问题，分配和回收内容都是消耗 CPU 资源的一种行为。

Go 的可执行文件都比相对应的源代码文件要大很多，这恰恰说明了 Go 的 runtime 嵌入到了每一个可执行文件当中。当然，在部署到数量巨大的集群时，较大的文件体积也是比较头疼的问题。但总的来说，Go 的部署工作还是要比 Java 和 Python 轻松得多。因为 Go 不需要依赖任何其它文件，它只需要一个单独的静态文件，这样你也不会像使用其它语言一样在各种不同版本的依赖文件之间混淆。



## 语言结构

###  Hello World 实例

Go 语言的基础组成有以下几个部分：

- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

 hello.go

```go
package main

import "fmt"

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
}
```

1. 第一行代码 `package main` 定义了`包名`。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
2. 下一行 `import "fmt"` 告诉 Go 编译器这个程序需要使用 `fmt 包`（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
3. 下一行 `func main()` 是程序开始执行的函数。main 函数是每一个可执行程序所`必须包含`的，一般来说都是在启动后第一个执行的函数（如果有 `init()` 函数则会先执行该函数）。
4. 下一行` /*...*/ `是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。
5. 下一行 `fmt.Println(...)*`可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
   使用 fmt.Print("hello, world\n") 可以得到相同的结果。
   `Print` 和 `Println` 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
6. 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个`大写字母开头`，如：Group1，那么使用这种形式的标识符的对象就可以被`外部包`的代码所使用（客户端程序需要先导入这个包），这被称为`导出`（像面向对象语言中的 `public`）；标识符如果以`小写字母开头`，则对包外是`不可见`的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 `protected` ）。



-  文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
-  文件夹名与包名没有直接关系，并非需要一致。
-  **同一个文件夹下的文件只能有一个包名，否则编译报错。**



### 执行 Go 程序

`.go文件` --> `go build`(编译) --> `可执行文件`(.exe) --> 运行(.exe) -->结果

`.go文件` --> `go run`(编译+运行) --> 结果



-  文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
-  文件夹名与包名没有直接关系，并非需要一致。
-  同一个文件夹下的文件只能有一个包名，否则编译报错。

#### go run

打开命令行，并进入程序文件保存的目录中。

```go
go run hello.go
Hello, World!
```

#### go build

```go
#还可以使用 go build 命令来生成二进制文件： Windows下

>go build hello.go 
>dir
>hello.exe   hello.go
>hello.exe
Hello, World!
```

### 标准库简介

在Go语言的安装文件里包含了一些可以直接使用的包，即标准库。Go语言的标准库（通常被称为语言自带的电池），提供了清晰的构建模块和公共接口，包含 I/O 操作、文本处理、图像、密码学、网络和分布式应用程序等，并支持许多标准化的文件格式和编解码协议。

在 Windows 下，标准库的位置在Go语言根目录下的子目录 pkg\windows_amd64 中；在 Linux 下，标准库在Go语言根目录下的子目录 pkg\linux_amd64 中（如果是安装的是 32 位，则在 linux_386 目录中）。一般情况下，标准包会存放在 $GOROOT/pkg/$GOOS_$GOARCH/ 目录下。

Go语言的编译器也是标准库的一部分，通过词法器扫描源码，使用语法树获得源码逻辑分支等。Go语言的周边工具也是建立在这些标准库上。在标准库上可以完成几乎大部分的需求。

Go语言的标准库以包的方式提供支持，下表列出了Go语言标准库中常见的包及其功能。



| Go语言标准库包名 | 功  能                                                       |
| ---------------- | ------------------------------------------------------------ |
| bufio            | 带缓冲的 I/O 操作                                            |
| bytes            | 实现字节操作                                                 |
| container        | 封装堆、列表和环形列表等容器                                 |
| crypto           | 加密算法                                                     |
| database         | 数据库驱动和接口                                             |
| debug            | 各种调试文件格式访问及调试功能                               |
| encoding         | 常见算法如 JSON、XML、Base64 等                              |
| flag             | 命令行解析                                                   |
| fmt              | 格式化操作                                                   |
| go               | Go语言的词法、语法树、类型等。可通过这个包进行代码信息提取和修改 |
| html             | HTML 转义及模板系统                                          |
| image            | 常见图形格式的访问及生成                                     |
| io               | 实现 I/O 原始访问接口及访问封装                              |
| math             | 数学库                                                       |
| net              | 网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等                |
| os               | 操作系统平台不依赖平台操作封装                               |
| path             | 兼容各操作系统的路径操作实用函数                             |
| plugin           | Go 1.7 加入的插件系统。支持将代码编译为插件，按需加载        |
| reflect          | 语言反射支持。可以动态获得代码中的类型信息，获取和修改变量的值 |
| regexp           | 正则表达式封装                                               |
| runtime          | 运行时接口                                                   |
| sort             | 排序接口                                                     |
| strings          | 字符串转换、解析及实用函数                                   |
| time             | 时间接口                                                     |
| text             | 文本模板及 Token 词法器                                      |

## 基础语法

### 注释

单行注释: ` //` 开头的单行注释

多行注释: 也叫块注释，均已以` /* `开头，并以 `*/ `结尾。

```go
// 单行注释

/*
 我是多行注释
 */
```

### 行分隔符

在 Go 程序中，`一行代表一个语句结束`。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。

如果你打算将多个语句写在同一行，它们则必须使用` ; `人为区分，但在实际开发中我们并不鼓励这种做法。

以下为两个语句：

```go
fmt.Println("Hello, World!")
fmt.Println("菜鸟教程：runoob.com")

fmt.Println("Hello, World!")fmt.Println("菜鸟教程：runoob.com") //不允许
```



### 标识符

个标识符实际上就是一个或是多个字母`(A~Z`和`a~z`)数字(`0~9`)、下划线`_`组成的序列，但是第`一个字符`必须是`字母`或`下划线`而不能是数字。

```go
//以下是有效的标识符：
mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal

#以下是无效的标识符：
1ab //以数字开头
case //Go 语言的关键字
a+b //运算符是不允许的
```



### 关键字

关键字不能用于自定义名字，只能在特定语法结构中使用。

下面列举了 Go 代码中会使用到的` 25` 个关键字或保留字：

```go
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var
```

此外，还有大约30多个预定义的名字，比如int和true等，主要对应内建的常量、类型和函数。

```go
内建常量: true false iota nil

内建类型: int int8 int16 int32 int64
          uint uint8 uint16 uint32 uint64 uintptr
          float32 float64 complex128 complex64
          bool byte rune string error

内建函数: make len cap new append copy close delete
          complex real imag
          panic recover
```

这些**内部预先定义的名字并不是关键字**，你可以再定义中重新使用它们。在一些特殊的场景中重新定义它们也是有意义的，但是也要注意避免过度而引起语义混乱。



程序一般由`关键字`、`常量`、`变量`、`运算符`、`类型`和`函数`组成。

程序中可能会使用到这些分隔符：括号` ()`，中括号` [] `和大括号` {}`。

程序中可能会使用到这些标点符号：`.`、`,`、`;`、`: `和` …`。



### 变量-var

#### 概述

变量可以通过变量名访问。

Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。

`严格区分大小写`， `不能是关键字`

Go语言程序员推荐使用 `驼峰式` 命名，当名字有几个单词组成的时优先使用大小写分隔，而不是优先用下划线分隔。

名字的长度没有逻辑限制，但是Go语言的风格是尽量使用短小的名字，对于局部变量尤其是这样；

声明变量的一般形式是使用` var` 关键字：

```go
var identifier type

//可以一次声明多个变量：
var identifier1, identifier2 type

//实例
package main
import "fmt"
func main() {
    var a string = "Runoob"
    fmt.Println(a) //Runoob1

    var b, c int = 1, 2
    fmt.Println(b, c) //1 2
}
```



需要注意的是，Go 和许多编程语言不同，它在声明变量时将变量的类型放在变量的名称之后。Go 为什么要选择这么做呢？

> 首先，它是为了避免像 C 语言中那样含糊不清的声明形式，例如：`int* a, b;`。在这个例子中，只有 a 是指针而 b 不是。如果你想要这两个变量都是指针，则需要将它们分开书写

而在 Go 中，则可以很轻松地将它们都声明为指针类型：

```go
var a, b *int
```

其次，这种语法能够按照从左至右的顺序阅读，使得代码更加容易理解。

示例：

```go
var a int
var b bool
var str string

//你也可以改写成这种形式：
//这种因式分解关键字的写法一般用于声明全局变量。
var (
	a int
	b bool
	str string
)
```

- 当一个变量被声明之后，系统自动赋予它该类型的零值：`int 为 0`，`float 为 0.0`，`bool 为 false`，`string 为空字符串`，`指针为 nil`。记住，所有的内存在 Go 中都是经过初始化的。

- 变量的命名规则遵循`骆驼命名法`，即首个单词小写，每个新单词的首字母大写，例如：`numShips` 和 `startDate`。

- 全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写
- 一个变量（常量、类型或函数）在程序中都有一定的作用范围，称之为`作用域`。如果一个变量在函数体外声明，则被认为是全局变量，可以在整个包甚至外部包（被导出后）使用，不管你声明在哪个源文件里或在哪个源文件里调用该变量。
- 在函数体内声明的变量称之为`局部变量`，它们的作用域只在函数体内，**参数和返回值变量也是局部变量。**像 if 和 for 这些控制结构，而在这些结构中声明的变量的作用域只在相应的代码块内。一般情况下，局部变量的作用域可以通过代码块（用大括号括起来的部分）判断。
- `标识符必须是唯一的`，但你可以在某个代码块的`内层代码块中使用相同名称的变量`，则此时外部的同名变量将会暂时隐藏（结束内部代码块的执行后隐藏的外部同名变量又会出现，而内部同名变量则被释放），你任何的操作都只会影响内部代码块的局部变量。
- 变量可以编译期间就被赋值，赋值给变量使用运算符等号 `=`，当然你也可以在运行时对变量进行赋值操作。



#### 变量声明初始化



`声明语句`定义了程序的各种`实体对象`以及部分或全部的`属性`。Go语言主要有四种类型的声明语句：`var`、`const`、`type`和`func`，分别对应`变量`、`常量`、`类型`和`函数实体`对象的声明。



##### 指定变量类型，如果没有初始化，则变量默认为零值

```go
var v_name v_type
v_name = value
```



```go
//零值就是变量没有做初始化时系统默认设置的值。
package main
import "fmt"
func main() {
    // 声明一个变量并初始化
    var a = "RUNOOB"
    fmt.Println(a) //RUNOOB

    // 没有初始化就为零值
    var b int
    fmt.Println(b) //0

    // bool 零值为 false
    var c bool
    fmt.Println(c) //false
}
```

- 值类型（包括complex64/128）为 **0**

- 布尔类型为 **false**

- 字符串为 **""**（空字符串）

- 以下几种类型为 **nil**：

- ```go
  var a *int
  var a []int
  var a map[string] int
  var a chan int
  var a func(string) int
  var a error // error 是接口
  ```

  ```go
  package main
  
  import "fmt"
  
  func main() {
      var i int
      var f float64
      var b bool
      var s string
      fmt.Printf("%v %v %v %q\n", i, f, b, s) //0 0 false ""
  }
  ```

##### 根据值自行判定变量类型

```go
var v_name = value
```

```go
package main
import "fmt"
func main() {
    var d = true
    fmt.Println(d) //true
}
```

##### 省略 var, 注意 :=左侧如果没有声明新的变量，就产生编译错误

```go
v_name := value
```

```go
var intVal int 
intVal :=1 // 这时候会产生编译错误
intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句


//可以将 var f string = "Runoob" 简写为 f := "Runoob"：
package main
import "fmt"
func main() {
    f := "Runoob" // var f string = "Runoob"

    fmt.Println(f) //Runoob
}
```

##### 多变量声明

```go
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

```go
package main

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h) //0 0 0 false 1 2 123 hello 123 hello
}
```

#### 使用 := 赋值操作符

可以在变量的初始化时省略变量的类型而由系统自动推断，声明语句写上 `var` 关键字其实是显得有些多余了，因此我们可以将它们简写为 a := 50 或 b := false。

a 和 b 的类型（int 和 bool）将由编译器自动推断。

这是使用变量的首选形式，但是它`只能被用在函数体内`，而不可以用于全局变量的声明与赋值。使用操作符` :=` 可以高效地创建一个新的变量，称之为`初始化声明`。



**注意事项**

1. 如果在相同的代码块中，我们不可以再次对于相同名称的变量使用`初始化声明`
   例如：`a := 20` 就是不被允许的，编译器会提示错误 no new variables on left side of :=，但是
   ` a = 20` 是可以的，因为这是给相同的变量赋予一个新的值。

2. `局部变量，声明必须使用`。如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误

3. 如果你想要交换两个变量的值，则可以简单地使用 **a, b = b, a**，两个变量的`类型必须是相同`。

4. 空白标识符 `_ `也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。

   ​		`_ `实际上是一个`只写变量`，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。

5. `并行赋值`也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。

```go
package main

import "fmt"

func main() {
   var a string = "abc"
   fmt.Println("hello, world")
}
//尝试编译这段代码将得到错误 a declared and not used。
//此外，单纯地给 a 赋值也是不够的，这个值必须被使用，所以使用
fmt.Println("hello, world", a) //会移除错误。


//但是全局变量是允许声明但不使用的。 同一类型的多个变量可以声明在同一行，如：
var a, b, c int

//多变量可以在同一行进行赋值，如：
var a, b int
var c string
a, b, c = 5, 7, "abc"

//上面这行假设了变量 a，b 和 c 都已经被声明，否则的话应该这样使用：
//这被称为 并行 或 同时 赋值。
a, b, c := 5, 7, "abc"
```

```go
//空白标识符在函数返回值时的使用：
package main
import "fmt"
func main() {
  _,numb,strs := numbers() //只获取函数返回值的后两个
  fmt.Println(numb,strs) //2 str
}

//一个可以返回多个值的函数
func numbers()(int,int,string){
  a , b , c := 1 , 2 , "str"
  return a,b,c
}
```

#### 匿名变量

匿名变量的特点是一个下画线“`_`”，“`_`”本身就是一个特殊的标识符，被称为`空白标识符`。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被`抛弃`，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。使用匿名变量时，只需要在变量声明的地方使用下画线替换即可。例如：

```go
func GetData() (int, int) {
    return 100, 200
}
func main(){
    a, _ := GetData()
    _, b := GetData()
    fmt.Println(a, b) //100 200
}
```

GetData() 是一个函数，拥有两个整型返回值。每次调用将会返回 100 和 200 两个数值。

代码说明如下：
第 5 行只需要获取第一个返回值，所以将第二个返回值的变量设为下画线（匿名变量）。
第 6 行将第一个返回值的变量设为匿名变量。

> **匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。**





#### 变量作用域

`作用域`为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：

- 函数内定义的变量称为`局部变量`
- 函数外定义的变量称为`全局变量`
- 函数定义中的变量称为`形式参数`

##### 局部变量

在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。

```go
//以下实例中 main() 函数使用了局部变量 a, b, c：
package main

import "fmt"

func main() {

   /* 声明局部变量 */
   var a, b, c int

   /* 初始化参数 */
   a = 10
   b = 20
   c = a + b

   fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)
   //结果： a = 10, b = 20 and c = 30    
}
```

##### 全局变量

在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

```go
//全局变量可以在任何函数中使用，以下实例演示了如何使用全局变量：
package main

import "fmt"

/* 声明全局变量 */
var g int

func main() {

   /* 声明局部变量 */
   var a, b int

   /* 初始化参数 */
   a = 10
   b = 20
   g = a + b

   fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
   //结果： a = 10, b = 20 and g = 30
}


//Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：
package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
   /* 声明局部变量 */
   var g int = 10

   fmt.Printf ("结果： g = %d\n",  g) //结果： g = 10
}
```

全局变量可以在整个包甚至外部包（被导出后）使用。

```go
//下述代码在 test.go 中定义了全局变量 Total_sum，然后在 hello.go 中引用。

//test.go:
package main
import "fmt"
var Total_sum int = 0
func Sum_test(a int, b int) int {
    fmt.Printf("%d + %d = %d\n", a, b, a+b)
    Total_sum += (a + b)
    fmt.Printf("Total_sum: %d\n", Total_sum)
    return a+b
}


//hello.go：
package main
import (
    "fmt"
)

func main() {
    var sum int
    sum = Sum_test(2, 3)
    fmt.Printf("sum: %d; Total_sum: %d\n", sum, Total_sum)
}
```



##### 形式参数

```go
//形式参数会作为函数的局部变量来使用。实例如下：
package main

import "fmt"

/* 声明全局变量 */
var a int = 20;

func main() {

   /* main 函数中声明局部变量 */
   var a int = 10
   var b int = 20
   var c int = 0

   fmt.Printf("main()函数中 a = %d\n",  a);
   c = sum( a, b);
   fmt.Printf("main()函数中 c = %d\n",  c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
   fmt.Printf("sum() 函数中 a = %d\n",  a);
   fmt.Printf("sum() 函数中 b = %d\n",  b);

   return a + b;
}

//以上实例执行输出结果为：
main()函数中 a = 10
sum() 函数中 a = 10
sum() 函数中 b = 20
main()函数中 c = 30
```



```go
//形参使用，比较 sum 函数中的 a 和 main 函数中的 a，sum 函数中虽然加了 1，但是 main 中还是原值 10:

package main

import "fmt"

/* 声明全局变量 */
var a int = 20

func main() {
    /* main 函数中声明局部变量 */
    var a int = 10
    var b int = 20
    var c int = 0

    fmt.Printf("main()函数中 a = %d\n", a)
    c = sum(a, b)
    fmt.Printf("main()函数中 a = %d\n", a)
    fmt.Printf("main()函数中 c = %d\n", c)
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
    a = a + 1
    fmt.Printf("sum() 函数中 a = %d\n", a)
    fmt.Printf("sum() 函数中 b = %d\n", b)
    return a + b
}

//输出为：
main()函数中 a = 10
sum() 函数中 a = 11
sum() 函数中 b = 20
main()函数中 a = 10
main()函数中 c = 31
```

##### 注意

```go
package main

import "fmt"

func main(){
  var a int = 0
  fmt.Println("for start")
  for a:=0; a < 10; a++ {
    fmt.Println(a)
  }
  fmt.Println("for end")

  fmt.Println(a)
}
//输出为：
for start
0
1
2
3
4
5
6
7
8
9
for end
0
//在 for 循环的 initialize（a:=0） 中，此时 initialize 中的 a 与外层的 a 不是同一个变量，initialize 中的 a 为 for 循环中的局部变量，因此在执行完 for 循环后，输出 a 的值仍然为 0。


package main

import "fmt"

func main(){
  var a int = 0
  fmt.Println("for start")
  for a = 0; a < 10; a++ {
    fmt.Println(a)
  }
  fmt.Println("for end")

  fmt.Println(a)
}
//输出为：
for start
0
1
2
3
4
5
6
7
8
9
for end
10
//此时 initialize 中的 a 便于外层的 a 为同一个变量，因此在执行完 for 循环后，输出 a 的值为 10。
```



##### 花括号控制变量的作用域

可通过花括号来控制变量的作用域，花括号中的变量是单独的作用域，同名变量会覆盖外层。

```go
//1.
a := 5
{
    a := 3
    fmt.Println("in a = ", a) /in a = 3
}
fmt.Println("out a = ", a) //out a = 5


//2.
a := 5
{
    fmt.Println("in a = ", a) //in a = 5
}
fmt.Println("out a = ", a) //out a = 5
```



#### 变量的生命周期

变量的生命周期指的是在程序运行期间变量有效存在的时间间隔。

变量的生命周期与变量的作用域有着不可分割的联系：

- 全局变量：它的生命周期和整个程序的运行周期是一致的；
- 局部变量：它的生命周期则是动态的，从创建这个变量的声明语句开始，到这个变量不再被引用为止；
- 形式参数和函数返回值：它们都属于局部变量，在函数被调用的时候创建，函数调用结束后被销毁。

```go
for t := 0.0; t < cycles*2*math.Pi; t += res {
    x := math.Sin(t)
    y := math.Sin(t*freq + phase)
    img.SetColorIndex(
        size+int(x*size+0.5), size+int(y*size+0.5),
        blackIndex, // 最后插入的逗号不会导致编译错误，这是Go编译器的一个特性
    )               // 小括号另起一行缩进，和大括号的风格保存一致
}
```

上面代码中，在每次循环的开始会创建临时变量 t，然后在每次循环迭代中创建临时变量 x 和 y。临时变量 x、y 存放在栈中，随着函数执行结束（执行遇到最后一个`}`），释放其内存。



栈和堆的区别在于：

- 堆（heap）：堆是用于存放进程执行中被动态分配的内存段。它的大小并不固定，可动态扩张或缩减。当进程调用 malloc 等函数分配内存时，新分配的内存就被动态加入到堆上（堆被扩张）。当利用 free 等函数释放内存时，被释放的内存从堆中被剔除（堆被缩减）；
- 栈(stack)：栈又称堆栈， 用来存放程序暂时创建的局部变量，也就是我们函数的大括号`{ }`中定义的局部变量。

在程序的编译阶段，编译器会根据实际情况自动选择在栈或者堆上分配局部变量的存储空间，不论使用 var 还是 new 关键字声明变量都不会影响编译器的选择。

```go
var global *int
func f() {
    var x int
    x = 1
    global = &x
}
func g() {
    y := new(int)
    *y = 1
}
```

上述代码中，函数 f 里的变量 x 必须在堆上分配，因为它在函数退出后依然可以通过包一级的 global 变量找到，虽然它是在函数内部定义的。用Go语言的术语说，这个局部变量 x 从函数 f 中逃逸了。

相反，当函数 g 返回时，变量 *y 不再被使用，也就是说可以马上被回收的。因此，*y 并没有从函数 g 中逃逸，编译器可以选择在栈上分配 *y 的存储空间，也可以选择在堆上分配，然后由Go语言的 GC（垃圾回收机制）回收这个变量的内存空间。

在实际的开发中，并不需要刻意的实现变量的逃逸行为，因为逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响。

虽然Go语言能够帮助我们完成对内存的分配和释放，但是为了能够开发出高性能的应用我们任然需要了解变量的声明周期。例如，如果将局部变量赋值给全局变量，将会阻止 GC 对这个局部变量的回收，导致不必要的内存占用，从而影响程序的性能。

### 常量-const-iota

#### const

Go语言中的常量使用关键字 const 定义，用于存储不会改变的数据，常量是在编译时被创建的，即使定义在函数内部也是如此，并且只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。由于编译时的限制，定义常量的表达式必须为能被编译器求值的常量表达式。

存储在常量中的数据类型只可以是`布尔型`、`数字型`（`整数型`、`浮点型`和`复数`）和`字符串型`。

常量的定义格式：`const identifier [type] = value`



```go
package main //必须有一个main包

import "fmt"

func main() {
	//变量：程序运行期间，可以改变的量， 变量声明需要var
	//常量：程序运行期间，不可以改变的量，常量声明需要const

	const a int = 10
	//a = 20 //err, 常量不允许修改
	fmt.Println("a = ", a)

	const b = 11.2 //没有使用:=
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)
}

```

在 Go 语言中，你可以省略类型说明符 `[type]`，因为编译器可以根据变量的值来推断其类型。

- 显式类型定义： `const b string = "abc"`
- 隐式类型定义： `const b = "abc"`



一个没有指定类型的常量被使用时，会根据其使用环境而推断出它所需要具备的类型。换句话说，未定义类型的常量会在必要时刻根据上下文来获得相关类型。

```go
var n int
f(n + 5) // 无类型的数字型常量 “5” 它的类型在这里变成了 int
```

常量的值必须是能够在编译时就能够确定的；你可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。

- 正确的做法：`const c1 = 2/3`
- 错误的做法：`const c2 = getNumber()` // 引发构建错误: `getNumber() used as value`

**因为在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用，如：len()。**

数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出：

```go
//反斜杠 \ 可以在常量表达式中作为多行的连接符使用。
const Ln2 = 0.693147180559945309417232121458\
			176568075500134360255254120680009

const Log2E = 1/Ln2 // this is a precise reciprocal
const Billion = 1e9 // float constant
const hardEight = (1 << 100) >> 97
```

> 当常量赋值给一个精度过小的数字型变量时，可能会因为无法正确表达常量所代表的数值而导致溢出，这会在编译期间就引发错误。

另外，常量也允许使用并行赋值的形式：

```go
const beef, two, c = "eat", 2, "veg"
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)
```

常量还可以用作枚举：

```go
const (
	Unknown = 0
	Female = 1
	Male = 2
)
```

数字 0、1 和 2 分别代表未知性别、女性和男性。

常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：

```go
package main

import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c) //abc 3 16
}
//字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
```

#### iota-常量生成器

`iota`，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

iota 可以被用作枚举值：

```go
const (
    a = iota
    b = iota
    c = iota
)
```

`第一个 iota 等于 0`，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 `a=0, b=1, c=2 `可以简写为如下形式：

```go
const (
    a = iota
    b
    c
)
```

iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始。

```go
package main

const (
    i = iota
    j = iota
    x = iota
)
const xx = iota
const yy = iota
func main(){
    println(i, j, x, xx, yy)
}

// 输出是 0 1 2 0 0
```

iota 用法

```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i) //0 1 2 ha ha 100 100 7 8
}
```



```go
package main

import "fmt"
const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() { 
    fmt.Println("i=",i) //i= 1
    fmt.Println("j=",j) //j= 6
    fmt.Println("k=",k) //k= 12
    fmt.Println("l=",l) //l= 24
}
```

iota 表示从 0 开始自动加 1，所以 **i=1<<0**, **j=3<<1**（**<<** 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 **k=3<<2**，**l=3<<3**。

简单表述:

- i=1：左移 0 位,不变仍为 1;

- **j=3**：左移 1 位,变为二进制 110, 即 6;
- **k=3**：左移 2 位,变为二进制 1100, 即 12;
- **l=3**：左移 3 位,变为二进制 11000,即 24。

注：**<<n==\*(2^n)**。

```go
package main //必须有一个main包

import "fmt"

func main() {
	//1、iota常量自动生成器，每个一行，自动累加1
	//2、iota给常量赋值使用
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	//3、iota遇到const，重置为0
	const d = iota
	fmt.Printf("d = %d\n", d)

	//4、可以只写一个iota
	const (
		a1 = iota //0
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

	//5、如果是同一行，值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d\n", i, j1, j2, j3, k)

}
```

#### 无类型常量

Go语言的常量有个不同寻常之处。虽然一个常量可以有任意一个确定的基础类型，例如 int 或 float64，或者是类似 time.Duration 这样的基础类型，但是许多常量并没有一个明确的基础类型。

编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算，可以认为至少有 256bit 的运算精度。这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。

通过延迟明确常量的具体类型，不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。

math.Pi 无类型的浮点数常量，可以直接用于任意需要浮点数或复数的地方：

```go
var x float32 = math.Pivar y float64 = math.Pivar z complex128 = math.Pi
```

如果 math.Pi 被确定为特定类型，比如 float64，那么结果精度可能会不一样，同时对于需要 float32 或 complex128 类型值的地方则需要一个明确的强制类型转换：

```go
const Pi64 float64 = math.Pivar x float32 = float32(Pi64)var y float64 = Pi64var z complex128 = complex128(Pi64)
```

对于常量面值，不同的写法可能会对应不同的类型。例如 0、0.0、0i 和 \u0000 虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符等不同的常量类型。同样，true 和 false 也是无类型的布尔类型，字符串面值常量是无类型的字符串类型。

### 运算符

运算符用于在程序运行时执行数学或逻辑运算。

Go 语言内置的运算符有：

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 其他运算符

#### 算术运算符

下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述 | 实例               |
| :----- | :--- | :----------------- |
| +      | 相加 | A + B 输出结果 30  |
| -      | 相减 | A - B 输出结果 -10 |
| *      | 相乘 | A * B 输出结果 200 |
| /      | 相除 | B / A 输出结果 2   |
| %      | 求余 | B % A 输出结果 0   |
| ++     | 自增 | A++ 输出结果 11    |
| --     | 自减 | A-- 输出结果 9     |

Go 的自增，自减只能作为表达式使用，而不能用于赋值语句。

```go
a++ // 这是允许的，类似 a = a + 1,结果与 a++ 相同
a-- //与 a++ 相似
a = a++ // 这是不允许的，会出现变异错误 syntax error: unexpected ++ at end of statement
```

```go
//以下实例演示了各个算术运算符的用法：
package main

import "fmt"

func main() {

   var a int = 21
   var b int = 10
   var c int

   c = a + b
   fmt.Printf("第一行 - c 的值为 %d\n", c ) //第一行 - c 的值为 31
   c = a - b
   fmt.Printf("第二行 - c 的值为 %d\n", c ) //第二行 - c 的值为 11
   c = a * b 
   fmt.Printf("第三行 - c 的值为 %d\n", c ) //第三行 - c 的值为 210
   c = a / b
   fmt.Printf("第四行 - c 的值为 %d\n", c ) //第四行 - c 的值为 2
   c = a % b
   fmt.Printf("第五行 - c 的值为 %d\n", c ) //第五行 - c 的值为 1
   a++
   fmt.Printf("第六行 - a 的值为 %d\n", a ) //第六行 - a 的值为 22
   a=21   // 为了方便测试，a 这里重新赋值为 21
   a--
   fmt.Printf("第七行 - a 的值为 %d\n", a ) //第七行 - a 的值为 20
}
```

#### 关系运算符

下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述                                                         | 实例              |
| :----- | :----------------------------------------------------------- | :---------------- |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       | (A == B) 为 False |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   | (A != B) 为 True  |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   | (A > B) 为 False  |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   | (A < B) 为 True   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | (A >= B) 为 False |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | (A <= B) 为 True  |

```go
//以下实例演示了关系运算符的用法：
package main

import "fmt"

func main() {
   var a int = 21
   var b int = 10

   if( a == b ) {
      fmt.Printf("第一行 - a 等于 b\n" )
   } else {
      fmt.Printf("第一行 - a 不等于 b\n" )
   }
   if ( a < b ) {
      fmt.Printf("第二行 - a 小于 b\n" )
   } else {
      fmt.Printf("第二行 - a 不小于 b\n" )
   }
   
   if ( a > b ) {
      fmt.Printf("第三行 - a 大于 b\n" )
   } else {
      fmt.Printf("第三行 - a 不大于 b\n" )
   }
   /* Lets change value of a and b */
   a = 5
   b = 20
   if ( a <= b ) {
      fmt.Printf("第四行 - a 小于等于 b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("第五行 - b 大于等于 a\n" )
   }
}

//以上实例运行结果：
第一行 - a 不等于 b
第二行 - a 不小于 b
第三行 - a 大于 b
第四行 - a 小于等于 b
第五行 - b 大于等于 a
```

#### 逻辑运算符

下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。

| 运算符 | 描述                                                         | 实例               |
| :----- | :----------------------------------------------------------- | :----------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 | (A && B) 为 False  |
| \|\|   | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 | (A \|\| B) 为 True |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 | !(A && B) 为 True  |

```go
//以下实例演示了逻辑运算符的用法：
package main

import "fmt"

func main() {
   var a bool = true
   var b bool = false
   if ( a && b ) {
      fmt.Printf("第一行 - 条件为 true\n" )
   }
   if ( a || b ) {
      fmt.Printf("第二行 - 条件为 true\n" )
   }
   /* 修改 a 和 b 的值 */
   a = false
   b = true
   if ( a && b ) {
      fmt.Printf("第三行 - 条件为 true\n" )
   } else {
      fmt.Printf("第三行 - 条件为 false\n" )
   }
   if ( !(a && b) ) {
      fmt.Printf("第四行 - 条件为 true\n" )
   }
}

//以上实例运行结果：
第二行 - 条件为 true
第三行 - 条件为 false
第四行 - 条件为 true
```

#### 位运算符

Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：

| 运算符 | 描述                                                         | 实例                                   |
| :----- | :----------------------------------------------------------- | :------------------------------------- |
| &      | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 | (A & B) 结果为 12, 二进制为 0000 1100  |
| \|     | 按位或运算符"\|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 | (A \| B) 结果为 61, 二进制为 0011 1101 |
| ^      | 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 | (A ^ B) 结果为 49, 二进制为 0011 0001  |
| <<     | 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >>     | 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 | A >> 2 结果为 15 ，二进制为 0          |

位运算符对整数在内存中的二进制位进行操作。

下表列出了位运算符 &, |, 和 ^ 的计算：

| p    | q    | p & q | p \| q | p ^ q |
| :--- | :--- | :---- | :----- | :---- |
| 0    | 0    | 0     | 0      | 0     |
| 0    | 1    | 0     | 1      | 1     |
| 1    | 1    | 1     | 1      | 0     |
| 1    | 0    | 0     | 1      | 1     |

```go
//假定 A = 60; B = 13; 其二进制数转换为：

A = 0011 1100
B = 0000 1101

-----------------

A&B = 0000 1100

A|B = 0011 1101

A^B = 0011 0001
```

```go
以下实例演示了位运算符的用法：

实例
package main

import "fmt"

func main() {

   var a uint = 60      /* 60 = 0011 1100 */  
   var b uint = 13      /* 13 = 0000 1101 */
   var c uint = 0          

   c = a & b       /* 12 = 0000 1100 */
   fmt.Printf("第一行 - c 的值为 %d\n", c ) //第一行 - c 的值为 12

   c = a | b       /* 61 = 0011 1101 */ 
   fmt.Printf("第二行 - c 的值为 %d\n", c ) //第二行 - c 的值为 61

   c = a ^ b       /* 49 = 0011 0001 */
   fmt.Printf("第三行 - c 的值为 %d\n", c ) //第三行 - c 的值为 49

   c = a << 2     /* 240 = 1111 0000 */
   fmt.Printf("第四行 - c 的值为 %d\n", c ) //第四行 - c 的值为 240

   c = a >> 2     /* 15 = 0000 1111 */
   fmt.Printf("第五行 - c 的值为 %d\n", c ) //第五行 - c 的值为 15
}
```

#### 赋值运算符

下表列出了所有Go语言的赋值运算符。

| 运算符 | 描述                                           | 实例                                  |
| :----- | :--------------------------------------------- | :------------------------------------ |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 | C = A + B 将 A + B 表达式结果赋值给 C |
| +=     | 相加后再赋值                                   | C += A 等于 C = C + A                 |
| -=     | 相减后再赋值                                   | C -= A 等于 C = C - A                 |
| *=     | 相乘后再赋值                                   | C *= A 等于 C = C * A                 |
| /=     | 相除后再赋值                                   | C /= A 等于 C = C / A                 |
| %=     | 求余后再赋值                                   | C %= A 等于 C = C % A                 |
| <<=    | 左移后赋值                                     | C <<= 2 等于 C = C << 2               |
| >>=    | 右移后赋值                                     | C >>= 2 等于 C = C >> 2               |
| &=     | 按位与后赋值                                   | C &= 2 等于 C = C & 2                 |
| ^=     | 按位异或后赋值                                 | C ^= 2 等于 C = C ^ 2                 |
| \|=    | 按位或后赋值                                   | C \|= 2 等于 C = C \| 2               |

```go
//以下实例演示了赋值运算符的用法：
package main

import "fmt"

func main() {
   var a int = 21
   var c int

   c =  a
   fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )

   c +=  a
   fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

   c -=  a
   fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

   c *=  a
   fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

   c /=  a
   fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

   c  = 200;

   c <<=  2
   fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )

   c >>=  2
   fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

   c &=  2
   fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

   c ^=  2
   fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

   c |=  2
   fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )

}

//以上实例运行结果：
第 1 行 - =  运算符实例，c 值为 = 21
第 2 行 - += 运算符实例，c 值为 = 42
第 3 行 - -= 运算符实例，c 值为 = 21
第 4 行 - *= 运算符实例，c 值为 = 441
第 5 行 - /= 运算符实例，c 值为 = 21
第 6行  - <<= 运算符实例，c 值为 = 800
第 7 行 - >>= 运算符实例，c 值为 = 200
第 8 行 - &= 运算符实例，c 值为 = 0
第 9 行 - ^= 运算符实例，c 值为 = 2
第 10 行 - |= 运算符实例，c 值为 = 2
```

#### 其他运算符

下表列出了Go语言的其他运算符。

| 运算符 | 描述             | 实例                       |
| :----- | :--------------- | :------------------------- |
| &      | 返回变量存储地址 | &a; 将给出变量的实际地址。 |
| *      | 指针变量。       | *a; 是一个指针变量         |

```go
//指针变量 * 和地址值 & 的区别：指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。当变量前面有 * 标识时，才等同于 & 的用法，否则会直接输出一个整型数字。

func main() {
   var a int = 4
   var ptr *int
   ptr = &a
   println("a的值为", a);    // 4
   println("*ptr为", *ptr);  // 4
   println("ptr为", ptr);    // 824633794744
}
```



```go
//以下实例演示了其他运算符的用法：
package main

import "fmt"

func main() {
   var a int = 4
   var b int32
   var c float32
   var ptr *int

   /* 运算符实例 */
   fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a ); //第 1 行 - a 变量类型为 = int
   fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b ); //第 2 行 - b 变量类型为 = int32
   fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c ); //第 3 行 - c 变量类型为 = float32

   /*  & 和 * 运算符实例 */
   ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
   fmt.Printf("a 的值为  %d\n", a); //a 的值为  4
   fmt.Printf("*ptr 为 %d\n", *ptr); //*ptr 为 4
}
```

#### 运算符优先级

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

| 优先级 | 运算符           |
| :----- | :--------------- |
| 5      | * / % << >> & &^ |
| 4      | + - \| ^         |
| 3      | == != < <= > >=  |
| 2      | &&               |
| 1      | \|\|             |

使用括号来临时提升某个表达式的整体运算优先级。

```go
package main

import "fmt"

func main() {
   var a int = 20
   var b int = 10
   var c int = 15
   var d int = 5
   var e int;

   e = (a + b) * c / d;      // ( 30 * 15 ) / 5
   fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );

   e = ((a + b) * c) / d;    // (30 * 15 ) / 5
   fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );

   e = (a + b) * (c / d);   // (30) * (15/5)
   fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );

   e = a + (b * c) / d;     //  20 + (150/5)
   fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e );  
}

//以上实例运行结果：
(a + b) * c / d 的值为 : 90
((a + b) * c) / d 的值为  : 90
(a + b) * (c / d) 的值为  : 90
a + (b * c) / d 的值为  : 50
```

### fmt包的格式化输出输入

#### 格式说明

| ***\*格式\**** | ***\*含义\****                                               |
| -------------- | ------------------------------------------------------------ |
| %%             | 一个%字面量                                                  |
| %b             | 一个二进制整数值(基数为2)，或者是一个(高级的)用科学计数法表示的指数为2的浮点数 |
| %c             | 字符型。可以把输入的数字按照ASCII码相应转换为对应的字符      |
| %d             | 一个十进制数值(基数为10)                                     |
| %e             | 以科学记数法e表示的浮点数或者复数值                          |
| %E             | 以科学记数法E表示的浮点数或者复数值                          |
| %f             | 以标准记数法表示的浮点数或者复数值                           |
| %g             | 以%e或者%f表示的浮点数或者复数，任何一个都以最为紧凑的方式输出 |
| %G             | 以%E或者%f表示的浮点数或者复数，任何一个都以最为紧凑的方式输出 |
| %o             | 一个以八进制表示的数字(基数为8)                              |
| %p             | 以十六进制(基数为16)表示的一个值的地址，前缀为0x,字母使用小写的a-f表示 |
| %q             | 使用Go语法以及必须时使用转义，以双引号括起来的字符串或者字节切片[]byte，或者是以单引号括起来的数字 |
| %s             | 字符串。输出字符串中的字符直至字符串中的空字符（字符串以'\0‘结尾，这个'\0'即空字符） |
| %t             | 以true或者false输出的布尔值                                  |
| %T             | 使用Go语法输出的值的类型                                     |
| %U             | 一个用Unicode表示法表示的整型码点，默认值为4个数字字符       |
| %v             | 使用默认格式输出的内置或者自定义类型的值，或者是使用其类型的String()方式输出的自定义值，如果该方法存在的话 |
| %x             | 以十六进制表示的整型值(基数为十六)，数字a-f使用小写表示      |
| %X             | 以十六进制表示的整型值(基数为十六)，数字A-F使用小写表示      |

#### 输出

```go
package main //必须有一个main包

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14
	//%T操作变量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

	//%d 整型格式
	//%s 字符串格式
	//%c 字符个数
	//%f 浮点型个数
	fmt.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)
	//%v自动匹配格式输出
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)
}
```



#### 输人

```go
package main //必须有一个main包

import "fmt"

func main() {
	var a int //声明变量
	fmt.Printf("请输入变量a: ")

	//阻塞等待用户的输入
	//fmt.Scanf("%d", &a) //别忘了&
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
```

Printf和Println的区别

```go
package main //必须有一个main包

import "fmt"

func main() {
	a := 10
	//一段一段处理，自动加换行
	fmt.Println("a = ", a)

	//格式化输出， 把a的内容放在%d的位置
	// "a = 10\n" 这个字符串输出到屏幕，"\n"代表换行符
	fmt.Printf("a = %d\n", a)

	b := 20
	c := 30
	fmt.Println("a = ", a, ", b = ", b, ", c = ", c)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

}
```

### 随机数的使用

```go
package main 

import "fmt"
import "math/rand"
import "time"

func main() {
	//设置种子， 只需一次
	//如果种子参数一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano()) //以当前系统时间作为种子参数

	for i := 0; i < 5; i++ {

		//产生随机数
		//fmt.Println("rand = ", rand.Int()) //随机很大的数
		fmt.Println("rand = ", rand.Intn(100)) //限制在100内的数
	}
}
```



## 控制结构

### 条件语句

Go 语言提供了以下几种条件判断语句：

| 语句          | 描述                                                         |
| :------------ | :----------------------------------------------------------- |
| `if `         | **if 语句** 由一个布尔表达式后紧跟一个或多个语句组成。       |
| `if...else`   | **if 语句** 后可以使用可选的 **else 语句**, else 语句中的表达式在布尔表达式为 false 时执行。 |
| `if 嵌套`     | 你可以在 **if** 或 **else if** 语句中嵌入一个或多个 **if** 或 **else if** 语句。 |
| `switch `     | **switch** 语句用于基于不同条件执行不同动作。                |
| `select 语句` | **select** 语句类似于 **switch** 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。 |

> 注意：Go 没有`三目运算符`，所以不支持 **?:** 形式的条件判断。



#### if 语句

if 语句由布尔表达式后紧跟一个或多个语句组成。

If 在布尔表达式为` true `时，其后紧跟的语句块执行，如果为 `false `则不执行。

if 语句使用 tips：

**（1）** 不需使用括号将条件包含起来

**（2）** 大括号{}必须存在，即使只有一行语句

**（3）** 左括号必须在if或else的同一行

**（4）** 在if之后，条件语句之前，可以添加变量初始化语句，使用；进行分隔

**（5）** 在有返回值的函数中，最终的return不能在条件语句中

```go
//Go 编程语言中 if 语句的语法如下：
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}
```

```go
package main //必须有一个main包

import "fmt"

func main() {
	s := "屌丝"

	//if和{就是条件，条件通常都是关系运算符
	if s == "王思聪" { //左括号和if在同一行
		fmt.Println("左手一个妹子，右手一个大妈")
	}

	//if支持1个初始化语句, 初始化语句和判断条件以分号分隔
	if a := 10; a == 10 { //条件为真，指向{}语句
		fmt.Println("a == 10")
	}
}
```

```go
//用 If 语句判断偶数:
package main

import "fmt"

func main() {
    var s int ;    // 声明变量 s 是需要判断的数
    fmt.Println("输入一个数字：")
    fmt.Scan(&s)

    if s%2 == 0  { //     取 s 处以 2 的余数是否等于 0
        fmt.Print("s 是偶数\n") ;//如果成立
    }else {
        fmt.Print("s 不是偶数\n") ;//否则
    }
    fmt.Print("s 的值是：",s) ;
}
```

```go
//if 还有另外一种形式，它包含一个 statement 可选语句部分，该组件在条件判断之前运行。它的语法是：
if statement; condition {  
}

package main

import (  
    "fmt"
)

func main() {  
    if num := 10; num % 2 == 0 { // 判断数字是否为偶数
        fmt.Println(num,"偶数") 
    }  else {
        fmt.Println(num,"奇数")
    }
}
```

#### if...else

if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。

If 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则执行 else 语句块。

```go
//Go 编程语言中 if...else 语句的语法如下：

if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else {
  /* 在布尔表达式为 false 时执行 */
}

```

```go
使用 if else 判断一个数的大小：

实例
package main

import "fmt"

func main() {
   /* 局部变量定义 */
   var a int = 100;
 
   /* 判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" );
   } else {
       /* 如果条件为 false 则执行以下语句 */
       fmt.Printf("a 不小于 20\n" ); //a 不小于 20
   }
   fmt.Printf("a 的值为 : %d\n", a); //a 的值为 : 100
}
```

#### if ... else if ... else

```go
//if ... else if ... else... 实例：

package main

import "fmt"

func main() {
    var age int = 23
    if age == 25 {
        fmt.Println("true")
    } else if age < 25 {
        fmt.Println("too small")
    } else {
        fmt.Println("too big")
    }
}
```

#### if 语句嵌套

可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。

也可以在 if 语句中嵌套 else if...else 语句

```go
//Go 编程语言中 if...else 语句的语法如下：
if 布尔表达式 1 {
   /* 在布尔表达式 1 为 true 时执行 */
   if 布尔表达式 2 {
      /* 在布尔表达式 2 为 true 时执行 */
   }
}
```

```go
//嵌套使用 if 语句：

package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
 
   /* 判断条件 */
   if a == 100 {
       /* if 条件语句为 true 执行 */
       if b == 200 {
          /* if 条件语句为 true 执行 */
          fmt.Printf("a 的值为 100 ， b 的值为 200\n" ); //a 的值为 100 ， b 的值为 200
       }
   }
   fmt.Printf("a 值为 : %d\n", a ); //a 值为 : 100
   fmt.Printf("b 值为 : %d\n", b ); //b 值为 : 200
}
```

```go
//判断用户密码输入：
package main 

import"fmt"

func main(){
    var a int 
    var b int
    fmt.Printf("请输入密码：   \n")
    fmt.Scan(&a)
    if a == 666 {
    fmt.Printf("请再次输入密码：")
    fmt.Scan(&b)
        if b == 999 {
            fmt.Printf("密码正确，门锁已打开")
        }else{
            fmt.Printf("非法入侵，已自动报警")
        }
    }else{
        fmt.Printf("非法入侵")
    }
}
```

#### switch- case语句

##### switch- case

`switch 语句`用于基于不同条件执行不同动作，每一个 `case` 分支都是唯一的，从上至下逐一测试，直到匹配为止。

switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 **fallthrough** 。

变量 var1 可以是`任何类型`，而 val1 和 val2 则可以是`同类型的任意值`。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。`

可以多个可能符合条件的值，使用逗号分割它们，例如：`case val1, val2, val3`。

```go
//Go 编程语言中 switch 语句的语法如下：

switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

   switch {
      case grade == "A" :
         fmt.Printf("优秀!\n" )   //优秀!  
      case grade == "B", grade == "C" :
         fmt.Printf("良好\n" )      
      case grade == "D" :
         fmt.Printf("及格\n" )      
      case grade == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
   fmt.Printf("你的等级是 %s\n", grade );    //你的等级是 A    
}
```



```go
package main //必须有一个main包

import "fmt"

func main() {
	var num int
	fmt.Printf("请按下楼层：")
	fmt.Scan(&num)

	switch num { //switch后面写的是变量本身
	case 1:
		fmt.Println("按下的是1楼")
		//break //go语言保留了break关键字，跳出switch语言， 不写，默认就包含
		fallthrough //不跳出switch语句，后面的无条件执行
	case 2:
		fmt.Println("按下的是2楼")
		//break
		fallthrough
	case 3:
		fmt.Println("按下的是3楼")
		//break
		fallthrough
	case 4:
		fmt.Println("按下的是4楼")
		//break
		fallthrough
	default:
		fmt.Println("按下的是xxx楼")
	}

}

/**************************************************/

package main //必须有一个main包

import "fmt"

func main() {
	//支持一个初始化语句， 初始化语句和变量本身， 以分号分隔
	switch num := 4; num { //switch后面写的是变量本身
	case 1:
		fmt.Println("按下的是1楼")

	case 2:
		fmt.Println("按下的是2楼")

	case 3, 4, 5:
		fmt.Println("按下的是yyy楼")

	case 6:
		fmt.Println("按下的是4楼")

	default:
		fmt.Println("按下的是xxx楼")
	}

	score := 85
	switch { //可以没有条件
	case score > 90: //case后面可以放条件
		fmt.Println("优秀")
	case score > 80: //case后面可以放条件
		fmt.Println("良好")
	case score > 70: //case后面可以放条件
		fmt.Println("一般")
	default:
		fmt.Println("其它")
	}
}
```



##### Type Switch

switch 语句还可以被用于 `type-switch `来判断某个 interface 变量中实际存储的变量类型。

```go
//Type Switch 语法格式如下：
switch x.(type){
    case type:
       statement(s);      
    case type:
       statement(s); 
    /* 你可以定义任意个数的case */
    default: /* 可选 */
       statement(s);
}
```

```go
package main

import "fmt"

func main() {
   var x interface{}
     
   switch i := x.(type) {
      case nil:  
         fmt.Printf(" x 的类型 :%T",i)      //x 的类型 :<nil>           
      case int:  
         fmt.Printf("x 是 int 型")                      
      case float64:
         fmt.Printf("x 是 float64 型")          
      case func(int) float64:
         fmt.Printf("x 是 func(int) 型")                      
      case bool, string:
         fmt.Printf("x 是 bool 或 string 型" )      
      default:
         fmt.Printf("未知型")    
   }  
}
```

##### fallthrough

使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。

switch 从第一个判断表达式为 true 的 case 开始执行，如果 case 带有 fallthrough，程序会继续执行下一条 case，且它不会去判断下一个 case 的表达式是否为 true。

```go
package main

import "fmt"

func main() {

    switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")
    case false:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
    }
}

//以上代码执行结果为：
2、case 条件语句为 true
3、case 条件语句为 false
4、case 条件语句为 true
```

##### 注意

```go
//1. 支持多条件匹配
switch {
    case 1,2,3,4:
    default:
}

//2. 不同的 case 之间不使用 break 分隔，默认只会执行一个 case。

//3. 如果想要执行多个 case，需要使用 fallthrough 关键字，也可用 break 终止。
switch {
    case 1:
    ...
    if(...){
        break
    }

    fallthrough // 此时switch(1)会执行case1和case2，但是如果满足if条件，则只执行case1

    case 2:
    ...
    case 3:
}
```

####  select 语句

select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

```go
//Go 编程语言中 select 语句的语法如下：
select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s);
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s);
}
```

以下描述了 select 语句的语法：

- 每个 case 都必须是一个通信

- 所有 channel 表达式都会被求值

- 所有被发送的表达式都会被求值

- 如果任意某个通信可以进行，它就执行，其他被忽略。

- 如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。

  否则：

  1. 如果有 default 子句，则执行该语句。
  2. 如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

```go
//select 语句应用演示：
package main

import "fmt"

func main() {
   var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n") //no communication
   }    
}
```

select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测。

```go
//select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测。

package main

import (
    "fmt"
    "time"
)

func Chann(ch chan int, stopCh chan bool) {
    var i int
    i = 10
    for j := 0; j < 10; j++ {
        ch <- i
        time.Sleep(time.Second)
    }
    stopCh <- true
}

func main() {

    ch := make(chan int)
    c := 0
    stopCh := make(chan bool)

    go Chann(ch, stopCh)

    for {
        select {
        case c = <-ch:
            fmt.Println("Recvice", c)
            fmt.Println("channel")
        case s := <-ch:
            fmt.Println("Receive", s)
        case _ = <-stopCh:
            goto end
        }
    }
end:
}
```



### 循环语句

Go 语言提供了以下几种类型循环处理语句：

| 循环类型 | 描述                                 |
| :------- | :----------------------------------- |
| for 循环 | 重复执行语句块                       |
| 循环嵌套 | 在 for 循环中嵌套一个或多个 for 循环 |

#### 语法

Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。

和 C 语言的 for 一样：

```go
for init; condition; post { }
```

和 C 的 while 一样：

```go
for condition { }
```

和 C 的 for(;;) 一样：

```go
for { }
```

- init： 一般为赋值表达式，给控制变量赋初值；
- condition： 关系表达式或逻辑表达式，循环控制条件；
- post： 一般为赋值表达式，给控制变量增量或减量。

for语句执行过程如下：

- 1、先对表达式 1 赋初值；
- 2、判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

#### 实例

计算 1 到 100 的数字之和：

```go
package main

import "fmt"

func main() {

	//for 初始化条件 ;  判断条件 ;  条件变化 {
	//}
	//1+2+3 …… 100累加

	sum := 0

	//1) 初始化条件  i := 1
	//2) 判断条件是否为真， i <= 100， 如果为真，执行循环体，如果为假，跳出循环
	//3) 条件变化 i++
	//4) 重复2， 3， 4
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println("sum = ", sum)
}


```

init 和 post 参数是可选的，我们可以直接省略它，类似 `While` 语句。

以下实例在 sum 小于 10 的时候计算 sum 自相加后的值：

```go
package main

import "fmt"

func main() {
        sum := 1
        for ; sum <= 10; {
                sum += sum
        }
        fmt.Println(sum)

        // 这样写也可以，更像 While 语句形式
        for sum <= 10{
                sum += sum
        }
        fmt.Println(sum)
}

//输出结果为：
16
16
```



#### 循环嵌套

以下为 Go 语言嵌套循环的格式：

```go
for [condition |  ( init; condition; increment ) | Range]
{
   for [condition |  ( init; condition; increment ) | Range]
   {
      statement(s);
   }
   statement(s);
}
```

以下实例使用循环嵌套来输出 2 到 100 间的素数：

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var i, j int

   for i=2; i < 100; i++ {
      for j=2; j <= (i/j); j++ {
         if(i%j==0) {
            break; // 如果发现因子，则不是素数
         }
      }
      if(j > (i/j)) {
         fmt.Printf("%d  是素数\n", i);
      }
   }  
}
```



#### 循环控制语句

循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：

| 控制语句      | 描述                                             |
| :------------ | :----------------------------------------------- |
| break 语句    | 经常用于中断当前 for 循环或跳出 switch 语句      |
| continue 语句 | 跳过当前循环的剩余语句，然后继续进行下一轮循环。 |
| goto 语句     | 将控制转移到被标记的语句。                       |



#### 无限循环

如果循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：

```go
package main

import "fmt"

func main() {
    for true  {
        fmt.Printf("这是无限循环。\n");
    }
}

/***************************************************/
package main

import "fmt"
import "time"

func main() {

	i := 0

	for { //for后面不写任何东西，这个循环条件永远为真，死循环
		i++
		time.Sleep(time.Second) //演示1s

		if i == 5 {
			//break //跳出循环，如果嵌套多个循环，跳出最近的那个内循环
			continue //跳过本次循环，下一次继续
		}
		fmt.Println("i = ", i)
	}
}
```

#### range的使用

Go 语言中 `range` 关键字用于 for 循环中迭代`数组`(array)、`切片`(slice)、`通道`(channel)或`集合`(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

```go
package main //必须有一个main包

import "fmt"

func main() {

	str := "abc"

	//通过for打印每个字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	//迭代打印每个元素，默认返回2个值: 一个是元素的位置，一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	for i := range str { //第2个返回值，默认丢弃，返回元素的位置(下标)
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	for i, _ := range str { //第2个返回值，默认丢弃，返回元素的位置(下标)
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
}
```



```go
package main

import "fmt"

func main() {

    //这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    
    fmt.Println("sum:", sum)
    //在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    //range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

//以上实例运行输出结果为：
sum: 9
index: 1
a -> apple
b -> banana
0 103
1 111
```



##### Range 简单循环

```go
package main

import "fmt"

func main(){
    nums := []int{1,2,3,4};
    length := 0;
    for range nums {                                                  
        length++;
    }
    fmt.Println( length);
}
```



##### 循环键值对

```go
package main

import "fmt"

func main(){
   nums := []int{1,2,3,4}
   for i,num := range nums {
      fmt.Printf("索引是%d,长度是%d\n",i, num)
   }
}

//输出结果为：
索引是0,长度是1
索引是1,长度是2
索引是2,长度是3
索引是3,长度是4
```

#####  range 获取参数列表

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(len(os.Args))
    for _, arg := range os.Args {
        fmt.Println(arg)
    }
}
```



#### 案例

##### 1-100 素数：

```go
package main
import "fmt"
func main() {
    var C, c int//声明变量
    C=1 /*这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）*/
    A: for C < 100 {
           C++ //C=1不能写入for这里就不能写入
           for c=2; c < C ; c++ {
               if C%c==0 {
                   goto A //若发现因子则不是素数
               }
           }
           fmt.Println(C,"是素数")
    }
}

//另一个方法输出 1-100 素数:
package main
import "fmt"

func main() {
    var a, b int
    for a = 2; a <= 100; a++ {
        for b = 2; b <= (a / b); b++ {
            if a%b == 0 {
                break
            }
        }
        if b > (a / b) {
            fmt.Printf("%d\t是素数\n", a)
        }
    }
}
```

##### 99乘法表

```go
package main
import "fmt"

func main() {
    var a, b int
    for a = 2; a <= 100; a++ {
        for b = 2; b <= (a / b); b++ {
            if a%b == 0 {
                break
            }
        }
        if b > (a / b) {
            fmt.Printf("%d\t是素数\n", a)
        }
    }
}
```

##### 成绩查询

```go
package main
import "fmt"
import "strconv"
import "os"

func main(){
    var score int = 0
    var fenshu string = "A"
    fmt.Printf("欢迎进入成绩查询系统\n")
    
    ZHU: for true{
        var xuanzhe int = 0
        fmt.Println("1.进入成绩查询 2.退出程序")
        fmt.Printf("请输入序号选择:")
        fmt.Scanln(&xuanzhe)
        fmt.Printf("\n")
        if xuanzhe == 1{
             goto CHA
        }
        if xuanzhe == 2{
            os.Exit(1)
        }
    }

    CHA: for true {
        fmt.Printf("请输入一个学生的成绩:")
        fmt.Scanln(&score)

        switch {
            case score >= 90:fenshu = "A"

            case score >= 80&&score < 90:fenshu = "B"

            case score >= 60&&score < 80:fenshu = "C"

            default: fenshu = "D"
        }

        //fmt.Println(fenshu)
        var c string  = strconv.Itoa(score)
        switch{
            case fenshu == "A":
                fmt.Printf("你考了%s分,评价为%s,成绩优秀\n",c,fenshu)
            case fenshu == "B" || fenshu == "C":
                fmt.Printf("你考了%s分,评价为%s,成绩良好\n",c,fenshu)
            case fenshu == "D":
                fmt.Printf("你考了%s分,评价为%s,成绩不合格\n",c,fenshu)
        }
        fmt.Printf("\n")
        goto ZHU
}
    //fmt.Println(score)
}
```

### break 语句

Go 语言中 break 语句用于以下两方面：

- 用于循环语句中跳出循环，并开始执行循环之后的语句。
- break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。
- 在多重循环中，可以用标号 label 标出想 break 的循环。

```go
//break 语法格式如下：
break;
```

```go
//在变量 a 大于 15 的时候跳出循环：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* for 循环 */
   for a < 20 {
      fmt.Printf("a 的值为 : %d\n", a);
      a++;
      if a > 15 {
         /* 使用 break 语句跳出循环 */
         break;
      }
   }
}

//以上实例执行结果为：
a 的值为 : 10
a 的值为 : 11
a 的值为 : 12
a 的值为 : 13
a 的值为 : 14
a 的值为 : 15

```

```go
//以下实例有多重循环，演示了使用标记和不使用标记的区别：
package main

import "fmt"

func main() {

    // 不使用标记
    fmt.Println("---- break ----")
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
                for i2 := 11; i2 <= 13; i2++ {
                        fmt.Printf("i2: %d\n", i2)
                        break
                }
        }

    // 使用标记
    fmt.Println("---- break label ----")
    re:
        for i := 1; i <= 3; i++ {
            fmt.Printf("i: %d\n", i)
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2: %d\n", i2)
                break re
            }
        }
}

//以上实例执行结果为：
---- break ----
i: 1
i2: 11
i: 2
i2: 11
i: 3
i2: 11
---- break label ----
i: 1
i2: 11    
```



### continue 语句

Go 语言的 continue 语句 有点像 break 语句。但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。

for 循环中，执行 continue 语句会触发 for 增量语句的执行。

在多重循环中，可以用标号 label 标出想 continue 的循环。

```go
//continue 语法格式如下：
continue;
```

```go
//在变量 a 等于 15 的时候跳过本次循环执行下一次循环：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* for 循环 */
   for a < 20 {
      if a == 15 {
         /* 跳过此次循环 */
         a = a + 1;
         continue;
      }
      fmt.Printf("a 的值为 : %d\n", a);
      a++;    
   }  
}

//以上实例执行结果为：
a 的值为 : 10
a 的值为 : 11
a 的值为 : 12
a 的值为 : 13
a 的值为 : 14
a 的值为 : 16
a 的值为 : 17
a 的值为 : 18
a 的值为 : 19
```

```go
//以下实例有多重循环，演示了使用标记和不使用标记的区别：
package main

import "fmt"

func main() {

    // 不使用标记
    fmt.Println("---- continue ---- ")
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2: %d\n", i2)
                continue
            }
    }

    // 使用标记
    fmt.Println("---- continue label ----")
    re:
        for i := 1; i <= 3; i++ {
            fmt.Printf("i: %d\n", i)
                for i2 := 11; i2 <= 13; i2++ {
                    fmt.Printf("i2: %d\n", i2)
                    continue re
                }
        }
}

//以上实例执行结果为：
---- continue ---- 
i: 1
i2: 11
i2: 12
i2: 13
i: 2
i2: 11
i2: 12
i2: 13
i: 3
i2: 11
i2: 12
i2: 13
---- continue label ----
i: 1
i2: 11
i: 2
i2: 11
i: 3
i2: 11
```



### goto 语句

Go 语言的` goto` 语句可以无条件地转移到过程中指定的行。

goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。

但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。

```go
//goto 语法格式如下：
goto label;
..
.
label: statement;
```

```go
//在变量 a 等于 15 的时候跳过本次循环并回到循环的开始语句 LOOP 处：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* 循环 */
   LOOP: for a < 20 {
      if a == 15 {
         /* 跳过迭代 */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", a)
      a++    
   }  
}

//以上实例执行结果为：
a的值为 : 10
a的值为 : 11
a的值为 : 12
a的值为 : 13
a的值为 : 14
a的值为 : 16
a的值为 : 17
a的值为 : 18
a的值为 : 19
```

```go
//打印九九乘法表:
package main 

import "fmt"

func main() {
    //print9x()
    gotoTag()
}

//嵌套for循环打印九九乘法表
func print9x() {
    for m := 1; m < 10; m++ {
        for n := 1; n <= m; n++ {
      fmt.Printf("%dx%d=%d ",n,m,m*n)
        }
        fmt.Println("")
    }
}

//for循环配合goto打印九九乘法表
func gotoTag() {
    for m := 1; m < 10; m++ {
    	n := 1
        LOOP: if n <= m {
            fmt.Printf("%dx%d=%d ",n,m,m*n)
            n++
            goto LOOP
        } else {
            fmt.Println("")
        }
    	n++
    }
}
```



## 数据类型

### 概述

在 Go 编程语言中，数据类型用于声明函数和变量。

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

Go 语言按类别有以下几种数据类型：

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | `布尔型`布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 |
| 2    | `数字类型` 整型 int 和浮点型 float32、float64，Go 语言支持`整型`和`浮点型`数字，并且支持`复数`，其中位的运算采用补码。 |
| 3    | `字符串类型`: 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 `UTF-8` 编码标识 Unicode 文本。 |
| 4    | `派生类型`: 包括： `指针类型`（Pointer）， `数组类型`，`结构化类型`(struct)， `Channel类型`，`函数类型`， `切片类型`， `接口类型`（interface）， `Map 类型` |



### 类型转换

类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：

```go
type_name(expression)
//type_name 为类型，expression 为表达式。
```

```go
//以下实例中将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量：
package main

import "fmt"

func main() {
   var sum int = 17
   var count int = 5
   var mean float32
   
   mean = float32(sum)/float32(count)
   fmt.Printf("mean 的值为: %f\n",mean) //mean 的值为: 3.400000
}
```

go` 不支持隐式转换类型`，比如 :

```go
package main
import "fmt"

func main() {  
    var a int64 = 3
    var b int32
    b = a
    fmt.Printf("b 为 : %d", b)
}
//此时会报错
cannot use a (type int64) as type int32 in assignment
cannot use b (type int32) as type string in argument to fmt.Printf

//但是如果改成 b = int32(a) 就不会报错了:
package main
import "fmt"

func main() {  
    var a int64 = 3
    var b int32
    b = int32(a)
    fmt.Printf("b 为 : %d", b)
}
```

### 类型别名

```go
package main //必须有一个main包

import "fmt"

func main() {
	//给int64起一个别名叫bigint
	type bigint int64

	var a bigint // 等价于var a int64
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d, ch = %c\n", b, ch)
}
```



### 值类型和引用类型

队列：先进先出
`栈`：先进后出，在程序调用的时候从栈空间去分配
`堆`：在程序调用的时候从系统的内存区分配

值类型和引用类型

- 类型分别有：int系列、float系列、bool、string、数组和结构体

```go
值类型：变量直接存储值，内容通常在栈中分配
 var i = 5       i -----> 5
```

- 引用类型有：指针、slice切片、管道channel、接口interface、map、函数等

```go
引用类型：变量存储的是一个地址，这个地址存储最终的值，内容通常在堆上分配，通过GC回收
ref r ------> 内存地址 -----> 值
```

值类型的特点是：变量直接存储值，内存通常在栈中分配

引用类型的特点是：变量存储的是一个地址，这个地址对应的空间里才是真正存储的值，内存通常在堆中分配

```go
package main

import (
    "fmt"
)

var a int = 5
var b = make(chan int, 3)

func main() {
    fmt.Println("值类型：", a)
    fmt.Println("引用类型：", b)
}

go run example1.go
值类型： 5
引用类型： 0xc00001c100

2. example2.go
package main

import (
    "fmt"
)

// func swap(a int, b int) {
//  tmp := a
//  a = b
//  b = tmp
// }

// a *int, b *int表示接受a,b的内存地址
func swap(a *int, b *int) {
    // *a，*b表示a,b内存地址所指向的值
    tmp := *a
    *a = *b
    *b = tmp
}

// func swap(a int, b int) (int, int) {
//  return b, a
// }

func main() {
    a := 5
    b := 10

    // swap(a, b)
    swap(&a, &b) // &a，&b表示引用a,b的内存地址
    // a, b = swap(a, b)
    // b, a = a, b

    fmt.Println("a=", a)
    fmt.Println("b=", b)
}

go run example2.go
a= 10
b= 5

3. example3.go
package main

import (
    "fmt"
)

func modify(a int) {
    a = 10
    return
}

func modify1(a *int) {
    *a = 10
    return
}

func main() {
    a := 5
    b := make(chan int, 1)

    fmt.Println("a=", a)
    fmt.Println("b=", b)

    modify(a) // 这个是复制一份到栈中改变a，作用域不一样，所有改变modify中的a不影响main中的a的值
    fmt.Println("a=", a)

    // &变量名：是引用变量名的内存地址
    // *变量名：是引用变量名的内存地址的值
    modify1(&a) // 这个是引用a的内存地址到modify1的作用域，然后在modify1中改变了和main中的a一样内存地址的空间的值，所以main中a的值改变了
    fmt.Println("a=", a)
}

go run example3.go
a= 5
b= 0xc0000160e0
a= 5
a= 10
```



### 基础数据类型

#### 分类

Go语言内置以下这些基础类型：

| 类型          | 名称     | 长度 | 零值  | 说明                                          |
| ------------- | -------- | ---- | ----- | --------------------------------------------- |
| bool          | 布尔类型 | 1    | false | 其值不为真即为家，不可以用数字代表true或false |
| byte          | 字节型   | 1    | 0     | uint8别名                                     |
| rune          | 字符类型 | 4    | 0     | 专用于存储unicode编码，等价于uint32           |
| int, uint     | 整型     | 4或8 | 0     | 32位或64位                                    |
| int8, uint8   | 整型     | 1    | 0     | -128 ~ 127, 0 ~ 255                           |
| int16, uint16 | 整型     | 2    | 0     | -32768 ~ 32767, 0 ~ 65535                     |
| int32, uint32 | 整型     | 4    | 0     | -21亿 ~ 21 亿, 0 ~ 42 亿                      |
| int64, uint64 | 整型     | 8    | 0     |                                               |
| float32       | 浮点型   | 4    | 0.0   | 小数位精确到7位                               |
| float64       | 浮点型   | 8    | 0.0   | 小数位精确到15位                              |
| complex64     | 复数类型 | 8    |       |                                               |
| complex128    | 复数类型 | 16   |       |                                               |
| uintptr       | 整型     | 4或8 |       | ⾜以存储指针的uint32或uint64整数              |
| string        | 字符串   |      | ""    | utf-8字符串                                   |

#### 整型

```go
package main //必须有一个main包

import "fmt"

func main() {
	//声明变量
	var v1 int32
    
    v1 = 123
    v2 := 64 // v1将会被自动推导为int类型
    
	fmt.Printf(v1, v2)

	

}
```



#### 浮点数

```go
package main //必须有一个main包

import "fmt"

func main() {
	//声明变量
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)

	//自动推导类型
	f2 := 3.14
	fmt.Printf("f2 type is %T\n", f2) //f2 type is float64

	//float64存储小数比float32更准确

}
```



#### 布尔型

```go
package main //必须有一个main包

import "fmt"

func main() {
	//1、声明变量, 没有初始化，零值（初始值）为false
	var a bool
	fmt.Println("a0 = ", a)

	a = true
	fmt.Println("a = ", a)

	//2、自动推导类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)
}
```



#### 字符串和字符

字符串类型

```go
package main //必须有一个main包

import "fmt"

func main() {
	var str1 string //声明变量
	str1 = "abc"
	fmt.Println("strl = ", str1)

	//自动推导类型
	str2 := "mike"
	fmt.Printf("str2 类型是 %T\n", str2)

	//内建函数，len()可以测字符串的长度，有多少个字符
	fmt.Println("len(str2) = ", len(str2))
}
```

字符类型

在Go语言中支持两个字符类型，一个是`byte`（实际上是uint8的别名），代表utf-8字符串的单个字节的值；另一个是`rune`，代表单个unicode字符。

```go
package main //必须有一个main包

import "fmt"

func main() {
	var ch byte //声明字符类型
	ch = 97
	//fmt.Println("ch = ", ch)
	//格式化输出，%c以字符方式打印，%d以整型方式打印
	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'a' //字符， 单引号
	fmt.Printf("%c, %d\n", ch, ch)

	//大写转小写，小写转大写, 大小写相差32，小写大
	fmt.Printf("大写：%d， 小写：%d\n", 'A', 'a')
	fmt.Printf("大写转小写：%c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)

	//'\'以反斜杠开头的字符是转义字符, '\n'代表换行
	fmt.Printf("hello go%c", '\n')
	fmt.Printf("hello itcast")

}
```

字符和字符串的区别

```go
package main //必须有一个main包

import "fmt"

func main() {
	var ch byte
	var str string

	//字符
	//1、单引号
	//2、字符，往往都只有一个字符，转义字符除外'\n'
	ch = 'a'
	fmt.Println("ch =", ch)

	//字符串
	//1、双引号
	//2、字符串有1个或多个字符组成
	//3、字符串都是隐藏了一个结束符，'\0'
	str = "a" // 由'a'和'\0'组成了一个字符串
	fmt.Println("str = ", str)

	str = "hello go"
	//只想操作字符串的某个字符，从0开始操作
	fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])

}
```



#### 复数

复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag）。

```go
package main //必须有一个main包

import "fmt"

func main() {
	var t complex128 //声明
	t = 2.1 + 3.14i  //赋值
	fmt.Println("t = ", t)

	//自动推导类型
	t2 := 3.3 + 4.4i
	fmt.Printf("t2 type is %T\n", t2)

	//通过内建函数，取实部和虚部
	fmt.Println("real(t2) = ", real(t2), ", imag(t2) = ", imag(t2))

}
```



### 复合数据类型

#### 分类

| 类型    | 名称   | 长度 | 默认值 | 说明     |
| ------- | ------ | ---- | ------ | -------- |
| pointer | 指针   |      | nil    |          |
| array   | 数组   |      | 0      |          |
| slice   | 切片   |      | nil    | 引⽤类型 |
| map     | 字典   |      | nil    | 引⽤类型 |
| struct  | 结构体 |      |        |          |



#### 指针

##### 概述

指针是一个代表着某个`内存地址的值`。这个内存地址往往是在内存中存储的另一个变量的值的起始位置。Go语言对指针的支持介于Java语言和C/C++语言之间，它既没有想Java语言那样取消了代码对指针的直接操作的能力，也避免了C/C++语言中由于对指针的滥用而造成的安全和可靠性问题。

变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 `&`，放到一个变量前使用就会返回相应变量的内存地址。

```go
//以下实例演示了变量在内存中地址：
package main

import "fmt"

func main() {
   var a int = 10  

   fmt.Printf("变量的地址: %x\n", &a  ) //变量的地址: 20818a220
}
```

一个指针变量指向了一个值的内存地址。

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

```go
var var_name *type
//type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：

var ip *int        /* 指向整型*/
var fp *float32    /* 指向浮点型 */
```

##### 使用指针

指针使用流程：

1. 定义指针变量。
2. 为指针变量赋值。
3. 访问指针变量中指向地址的值。
4. 在指针类型前面加上` * `号（前缀）来获取指针所指向的内容。

```go
package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  ) //a 变量的地址是: 20818a220

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip ) //ip 变量储存的指针地址: 20818a220

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip ) //*ip 变量的值: 20
}
```

##### Go 空指针

当一个指针被定义后没有分配到任何变量时，它的值为 `nil`。

**nil 指针也称为空指针。**

nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 `ptr`。

```go
package main

import "fmt"

func main() {
   var  ptr *int

   fmt.Printf("ptr 的值为 : %x\n", ptr  ) //ptr 的值为 : 0
}


//空指针判断：
if(ptr != nil)     /* ptr 不是空指针 */
if(ptr == nil)    /* ptr 是空指针 */
```



##### 基本操作

Go语言虽然保留了指针，但与其它编程语言不同的是：

- 默认值`nil`，没有 NULL 常量
- 操作符 `"&"` 取变量地址， "`*`" 通过指针访问目标对象
- 不支持指针运算，不支持 "->" 运算符，直接⽤` "." `访问目标成员

```go
package main

import "fmt"

func main() {
    
	var a int = 10
	//每个变量有2层含义：变量的内存，变量的地址
	fmt.Printf("a = %d\n", a) //变量的内存
	fmt.Printf("&a = %v\n", &a)

	//保存某个变量的地址，需要指针类型   *int 保存int的地址， **int 保存 *int 地址
	//声明(定义)， 定义只是特殊的声明
	//定义一个变量p, 类型为*int
	var p *int
	p = &a //指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p = %v, &a = %v\n", p, &a)

	*p = 666 //*p操作的不是p的内存，是p所指向的内存(就是a)
	fmt.Printf("*p = %v, a = %v\n", *p, a)
}
```

不要操作没有合法指向的内存

```go
package main 

import "fmt"

func main() {
    
	var p *int
    
    fmt.Println("p = ", p)
    
	p = nil
	fmt.Println("p = ", p)

	//*p = 666 //err, 因为p没有合法指向

	var a int
	p = &a //p指向a
	*p = 666
	fmt.Println("a = ", a)

}
```



#####  new函数

表达式`new(T)`将创建一个T类型的匿名变量，所做的是为T类型的新值分配并清零一块内存空间，然后将这块内存空间的地址作为结果返回，而这个结果就是指向这个新的T类型值的指针值，返回的指针类型为`*T`。

```go
package main

import "fmt"

func main() {
    
	//a := 10 //整型变量a

	var p *int
	//指向一个合法内存
	//p = &a

	//p是*int, 指向int类型
	p = new(int)

	*p = 666
	fmt.Println("*p = ", *p)

	q := new(int) //自动推导类型
	*q = 777
	fmt.Println("*q = ", *q)
}
```

我们只需使用new()函数，无需担心其内存的生命周期或怎样将其删除，因为Go语言的内存管理系统会帮我们打理一切。

##### 指针做函数参数

普通变量做函数参数

```go
package main //必须有个main包

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap: a = %d, b = %d\n", a, b)
}

func main() {
	a, b := 10, 20

	//通过一个函数交换a和b的内容
	swap(a, b) //变量本身传递，值传递（站在变量角度）
	fmt.Printf("main: a = %d, b = %d\n", a, b)
}
```

指针做函数参数

```go
package main //必须有个main包

import "fmt"

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

func main() {
    
	a, b := 10, 20

	//通过一个函数交换a和b的内容
	swap(&a, &b) //地址传递
	fmt.Printf("main: a = %d, b = %d\n", a, b)
}
```

##### 多重指針

如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。

当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址

![指向指针的指针](images\指向指针的指针.png)

```go
//指向指针的指针变量声明格式如下：
//指向指针的指针变量为整型。
var ptr **int;
```



访问指向指针的指针变量值需要使用两个 * 号：

```go
//访问指向指针的指针变量值需要使用两个 * 号：
package main

import "fmt"

func main() {

   var a int
   var ptr *int
   var pptr **int

   a = 3000

   /* 指针 ptr 地址 */
   ptr = &a

   /* 指向指针 ptr 地址 */
   pptr = &ptr

   /* 获取 pptr 的值 */
   fmt.Printf("变量 a = %d\n", a ) //变量 a = 3000
   fmt.Printf("指针变量 *ptr = %d\n", *ptr ) //指针变量 *ptr = 3000
   fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr) //指向指针的指针变量 **pptr = 3000
}
```



三重指针及其对应关系:

```go
ptr3 - >ptr2- > ptr1 - >变量a
```

```go
package main

import "fmt"
func main(){
    
    var a int = 5
    
    //把ptr指针 指向ss所在地址
    var ptr *int = &a
    
    //开辟一个新的指针，指向ptr指针指向的地方
    var pts *int = ptr 
    
    //二级指针，指向一个地址，这个地址存储的是一级指针的地址
    var pto **int = &ptr
    
    //三级指针，指向一个地址，这个地址存储的是二级指针的地址，二级指针同上
    var pt3 ***int = &pto
    
    fmt.Println("a的地址:",&a,
                "\n 值", a, "\n\n",

                "ptr指针所在地址:",&ptr,
                "\n ptr指向的地址:",ptr,
                "\n ptr指针指向地址对应的值",*ptr,"\n\n", 

                "pts指针所在地址:",&pts,
                "\n pts指向的地址:", pts,
                "\n pts指针指向地址对应的值:",*pts,"\n\n", 

                "pto指针所在地址:",&pto,
                "\n pto指向的指针（ptr）的存储地址:",pto, 
                "\n pto指向的指针（ptr）所指向的地址: " ,*pto, 
                "\n pto最终指向的地址对应的值（a）",**pto,"\n\n",

                "pt3指针所在的地址:",&pt3,
                "\n pt3指向的指针（pto）的地址:",pt3,//等于&*pt3,
                "\n pt3指向的指针（pto）所指向的指针的（ptr）地址", *pt3, //等于&**pt3,
                "\n pt3指向的指针（pto）所指向的指针（ptr）所指向的地址（a）:",**pt3, //等于&***pt3,
                "\n pt3最终指向的地址对应的值（a）", ***pt3)


}

//执行结果：
a的地址: 0xc00009a008 
 值 5 

 ptr指针所在地址: 0xc000092010 
 ptr指向的地址: 0xc00009a008 
 ptr指针指向地址对应的值 5 

 pts指针所在地址: 0xc000092018 
 pts指向的地址: 0xc00009a008 
 pts指针指向地址对应的值: 5 

 pto指针所在地址: 0xc000092020 
 pto指向的指针（ptr）的存储地址: 0xc000092010 
 pto指向的指针（ptr）所指向的地址:  0xc00009a008 
 pto最终指向的地址对应的值（a） 5 

 pt3指针所在的地址: 0xc000092028 
 pt3指向的指针（pto）的地址: 0xc000092020 
 pt3指向的指针（pto）所指向的指针的（ptr）地址 0xc000092010 
 pt3指向的指针（pto）所指向的指针（ptr）所指向的地址（a）: 0xc00009a008 
 pt3最终指向的地址对应的值（a） 5
```

##### 指针数组

定义了长度为 3 的整型数组：

```go
package main

import "fmt"

const MAX int = 3

func main() {

   a := []int{10,100,200}
   var i int

   for i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i, a[i] )
   }
}

a[0] = 10
a[1] = 100
a[2] = 200
```

有一种情况，我们可能需要保存数组，这样我们就需要使用到指针。

以下声明了整型`指针数组`：

```go
var ptr [MAX]*int;
```

ptr 为整型指针数组。因此每个元素都指向了一个值。以下实例的三个整数将存储在指针数组中：

```go
package main

import "fmt"

const MAX int = 3

func main() {
    
   a := []int{10,100,200}
   var i int
    
   var ptr [MAX]*int;

   for  i = 0; i < MAX; i++ {
      ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
   }

   for  i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
   }
}

//以上代码执行输出结果为：
a[0] = 10
a[1] = 100
a[2] = 200
```



#### 数组

##### 概述

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

数组中包含的每个数据被称为数组元素（element），一个数组包含的元素个数被称为数组的长度。

数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

![Go 语言数组](images\Go 语言数组.png)



数组`⻓度必须是常量`，且是类型的组成部分。 [2]int 和 [3]int 是不同类型。

```go
var n int = 10
var a [n]int  //err, non-constant array bound n
var b [10]int //ok

const MAX int = 6
var c [MAX]int //ok
```

```go
package main //必须有个main包

import "fmt"

func main() {
	//	id1 := 1
	//	id2 := 2
	//	id3 := 3

	//数组，同一个类型的集合
	var id [50]int

	//操作数组，通过下标， 从0开始，到len()-1
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}
}
```



##### 声明,初始化数组

声明数组

```go
//Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：
var variable_name [SIZE] variable_type
//以上为一维数组的定义方式。

//例如以下定义了数组 balance 长度为 10 类型为 float32：
var balance [10]float32



//声明数组：
nums := [3]int{1,2,3,}
//声明切片：
nums := []int{1,2,3}
```



Go 语言的数组是值，其长度是其类型的一部分，作为函数参数时，是 值传递，函数中的修改对调用者不可见

Go 语言中对数组的处理，一般采用 切片 的方式，切片包含对底层数组内容的引用，作为函数参数时，类似于 指针传递，函数中的修改对调用者可见

```go


// 数组
b := [...]int{2, 3, 5, 7, 11, 13}

func boo(tt [6]int) {
    tt[0], tt[len(tt)-1] = tt[len(tt)-1], tt[0]
}

boo(b)
fmt.Println(b) // [2 3 5 7 11 13]
// 切片
p := []int{2, 3, 5, 7, 11, 13}

func poo(tt []int) {
    tt[0], tt[len(tt)-1] = tt[len(tt)-1], tt[0]
}
poo(p)
fmt.Println(p)  // [13 3 5 7 11 2]
```

初始化数组

```go
//初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
//如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：


//该实例与上面的实例是一样的，虽然没有设置数组的大小。
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

//读取了第五个元素。数组元素可以通过索引（位置）来读取（或者修改），索引从0开始，第一个元素索引为 0，第二个索引为 1
 balance[4] = 50.0

/************************************************/
package main //必须有个main包

import "fmt"

func main() {
	//声明定义同时赋值，叫初始化
	//1、全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	//部分初始化，没有初始化的元素，自动赋值为0
	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)

	//指定某个元素初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Println("d = ", d)
}
```

##### 操作数组

数组的每个元素可以通过索引下标来访问，索引下标的范围是从0开始到数组长度减1的位置。

```go
 var a [10]int
    for i := 0; i < 10; i++ {
        a[i] = i + 1
        fmt.Printf("a[%d] = %d\n", i, a[i])
    }

    //range具有两个返回值，第一个返回值是元素的数组下标，第二个返回值是元素的值
    for i, v := range a {
        fmt.Println("a[", i, "]=", v)
    }
```

内置函数 len(长度) 和 cap(容量) 都返回数组⻓度 (元素数量)：

```go

a := [10]int{}
fmt.Println(len(a), cap(a))//10 10
```

相同类型的数组之间可以使用 == 或 != 进行比较，但不可以使用 < 或 >，也可以相互赋值：

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{1, 2}
fmt.Println(a == b, b == c) //true false

var d [3]int
d = a
fmt.Println(d) //[1 2 3]
```

以下演示了数组完整操作（声明、赋值、访问）的实例：

```go
package main

import "fmt"

func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   /* 为数组 n 初始化元素 */        
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}
```



##### 在函数间传递数组

根据内存和性能来看，在函数间传递数组是一个开销很大的操作。在函数之间传递变量时，总是以值的方式传递的。如果这个变量是一个数组，意味着整个数组，不管有多长，都会完整复制，并传递给函数。

```go
//方式一
//形参设定数组大小：
void myFunction(param [10]int)
{
.
.
.
}

//方式二
//形参未设定数组大小：
void myFunction(param []int)
{
.
.
.
}
```

初始化数组长度后，元素可以不进行初始化，或者不进行全部初始化，但未进行数组大小初始化的数组初始化结果元素大小就为多少。

```go
func main() {

    var array = []int{1, 2, 3, 4, 5}
    
    /* 未定义长度的数组只能传给不限制数组长度的函数 */
    setArray(array)
    
    /* 定义了长度的数组只能传给限制了相同数组长度的函数 */
    var array2 = [5]int{1, 2, 3, 4, 5}
    setArray2(array2)
}

func setArray(params []int) {
    fmt.Println("params array length of setArray is : ", len(params))
}

func setArray2(params [5]int) {
    fmt.Println("params array length of setArray2 is : ", len(params))
}
//输出：
params array length of setArray is :  5
params array length of setArray2 is :  5
```



```go
package main

import "fmt"

//数组做函数参数，它是值传递
//实参数组的每个元素给形参数组拷贝一份
//形参的数组是实参数组的复制品
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify a = ", a)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5} //初始化

	modify(a) //数组传递过去
	fmt.Println("main: a = ", a)
}
```

##### 数组指针做函数参数

```go
package main

import "fmt"

//p指向实现数组a，它是指向数组，它是数组指针
//*p代表指针所指向的内存，就是实参a
func modify(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("modify *a = ", *p)
}

func main() {
    
	a := [5]int{1, 2, 3, 4, 5} //初始化

	modify(&a) //地址传递
	fmt.Println("main: a = ", a)
}
```

##### 多维数组

###### 概述

Go 语言支持多维数组，以下为常用的多维数组声明方式：

```go
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

//以下实例声明了三维的整型数组：
var threedim [5][10][4]int
```

###### 二维数组

二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：

二维数组中的元素可通过 `a[ i ][ j ]`来访问。

```go
var arrayName [ x ][ y ] variable_type
```

variable_type 为 Go 语言的数据类型，arrayName 为数组名，二维数组可认为是一个表格，x 为行，y 为列，下图演示了一个二维数组 a 为三行四列：初始化二维数组

![二维数组](images\二维数组.png)

```go
package main //必须有个main包

import "fmt"

func main() {
	//有多少个[]就是多少维
	//有多少个[]就用多少个循环
	var a [3][4]int

	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d, ", i, j, a[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Println("a = ", a)

	//有3个元素，每个元素又是一维数组[4]int
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b = ", b)

	//部分初始化，没有初始化的值为0
	c := [3][4]int{{1, 2, 3}, {5, 6, 7, 8}, {9, 10}}
	fmt.Println("c = ", c)

	d := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println("d = ", d)

	e := [3][4]int{1: {5, 6, 7, 8}}
	fmt.Println("e = ", e)
}
```



###### 初始化二维数组

多维数组可通过大括号来初始值。以下实例为一个 3 行 4 列的二维数组：

```go
a = [3][4]int{  
 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
 {8, 9, 10, 11},   /* 第三行索引为 2 */
}

注意：以上代码中倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样：
a = [3][4]int{  
 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
 {8, 9, 10, 11}}   /* 第三行索引为 2 */
```

多维数组初始化或赋值时需要注意 Go 语法规范，该写在一行就写在一行，一行一条语句。

```go
package main 

import "fmt"

func main() {
  var a = [3][5]int {{1, 2, 3, 4, 5}, {0, 9, 8, 7, 6}, {3, 4, 5, 6, 7}}
  var i, j int
  for i = 0; i < 3; i++ {
    for j = 0; j < 5; j++ {
    fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j])
    }
  }
}
```

###### 访问二维数组 

二维数组通过指定坐标来访问。如数组中的行索引与列索引，例如：

```go
val := a[2][3]
//或
var value int = a[2][3]
//以上实例访问了二维数组 val 第三行的第四个元素。
```

二维数组可以使用循环嵌套来输出元素：

```go
package main

import "fmt"

func main() {
   /* 数组 - 5 行 2 列*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   var i, j int

   /* 输出数组元素 */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
}

//以上实例运行输出结果为：
a[0][0] = 0
a[0][1] = 0
a[1][0] = 1
a[1][1] = 2
a[2][0] = 2
a[2][1] = 4
a[3][0] = 3
a[3][1] = 6
a[4][0] = 4
a[4][1] = 8
```

将二维数组按行输出:

```go
//将二维数组按行输出:
package main

import "fmt"
func main()  {
    
    /* 数组 - 5 行 2 列 */
    var a = [5][2]int{{0,0}, {1,2}, {2,4}, {3,6}, {4,8}}
    
    var i, j int    /* 输出数组元素 */
    
    for i =0; i < 5; i++ {
        fmt.Printf("第 %d 行：", i)
        for j = 0; j < 2; j++ {
            fmt.Printf("%d， ", a[i][j])
        }
        // 换行
        fmt.Println()
    }
}

//输出结果为：
第 0 行：0， 0， 
第 1 行：1， 2， 
第 2 行：2， 4， 
第 3 行：3， 6， 
第 4 行：4， 8， 
```

##### 冒泡排序

```go
package main //必须有个main包

import "fmt"
import "math/rand"
import "time"

func main() {
    
	//设置种子， 只需一次
	rand.Seed(time.Now().UnixNano())

	var a [10]int
	n := len(a)

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100) //100以内的随机数
		fmt.Printf("%d, ", a[i])
	}
	fmt.Printf("\n")

	//冒泡排序，挨着的2个元素比较，升序（大于则交换）
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Printf("\n排序后:\n")
	for i := 0; i < n; i++ {
		fmt.Printf("%d, ", a[i])
	}
	fmt.Printf("\n")
}
```



#### Slice

##### 概述

Go 语言切片是`对数组的抽象`。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

切片并不是数组或数组指针，它通过内部指针和相关属性引⽤数组⽚段，以实现变⻓⽅案。

`slice`并不是真正意义上的动态数组，而是一个`引用类型`。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

![slice](images\slice.png)



##### 定义切片

声明一个未指定大小的数组来定义切片：

```go
var identifier []type
```

切片不需要说明长度。

或使用make()函数来创建切片:

```go
var slice1 []type = make([]type, len)
//也可以简写为
slice1 := make([]type, len)
```

也可以指定容量，其中capacity为可选参数。

```go
make([]T, length, capacity)
```

这里 len 是数组的长度并且也是切片的初始长度。



##### 切片初始化

```go
// 声明一个空切片
var numListEmpty = []int{}

//直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
s :=[] int {1,2,3} 

//初始化切片s,是数组arr的引用
s := arr[:] 

//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:endIndex] 

//默认 endIndex 时将表示一直到arr的最后一个元素
s := arr[startIndex:] 

//默认 startIndex 时将表示从arr的第一个元素开始
s := arr[:endIndex] 

//通过切片s初始化切片s1
s1 := s[startIndex:endIndex] 

//通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
s :=make([]int,len,cap) 
```



slice和数组的区别：声明数组时，方括号内写明了数组的长度或使用`...`自动计算长度，而声明slice时，方括号内没有任何字符。

```go
package main //必须有个main包

import "fmt"

func main() {
    
	//自动推导类型， 同时初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)

	//借助make函数, 格式 make(切片类型, 长度, 容量)
	s2 := make([]int, 5, 10)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	//没有指定容量，容量和长度一样
	s3 := make([]int, 5)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

}

func main01() {
	//切片和数组的区别
	//数组[]里面的长度时固定的一个常量， 数组不能修改长度， len和cap永远都是5
	a := [5]int{}
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))

	//切片， []里面为空，或者为...，切片的长度或容易可以不固定
	s := []int{}
	fmt.Printf("1: len = %d, cap = %d\n", len(s), cap(s))

	s = append(s, 11) //给切片末尾追加一个成员
	fmt.Printf("append: len = %d, cap = %d\n", len(s), cap(s))
}
```

##### make()函数

```go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```

注意：`make`只能创建`slice`、`map`和`channel`，并且返回一个有初始值(非零)。



在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量。在第一种语句中，slice是整个数组的view。在第二个语句中，slice只引用了底层数组的前len个元素，但是容量将包含整个的数组。额外的元素是留给未来的增长用的。

##### 空(nil)切片

一个切片在未初始化之前默认为` nil`，长度为 0，实例如下：

```go
package main

import "fmt"

func main() {
   var numbers []int

   printSlice(numbers) ////len=0 cap=0 slice=[]

   if(numbers == nil){
      fmt.Printf("切片是空的") //切片是空的
   }
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x) 
}
```

##### 多维切片

Go语言中同样允许使用多维切片，声明一个多维数组的语法格式如下：

var sliceName [][]...[]sliceType

其中，sliceName 为切片的名字，sliceType为切片的类型，每个`[ ]`代表着一个维度，切片有几个维度就需要几个`[ ]`。

下面以二维切片为例，声明一个二维切片并赋值，代码如下所示。

```go
//声明一个二维切片var slice [][]int//为二维切片赋值slice = [][]int{{10}, {100, 200}}
```

上面的代码也可以简写为下面的样子。

```go
// 声明一个二维整型切片并赋值slice := [][]int{{10}, {100, 200}}
```

上面的代码中展示了一个包含两个元素的外层切片，同时每个元素包又含一个内层的整型切片，切片 slice 的值如下图所示。

![切片 slice](images\切片 slice.png)

通过上图可以看到外层的切片包括两个元素，每个元素都是一个切片，第一个元素中的切片使用单个整数 10 来初始化，第二个元素中的切片包括两个整数，即 100 和 200。

【示例】组合切片的切片

```go
// 声明一个二维整型切片并赋值slice := [][]int{{10}, {100, 200}}// 为第一个切片追加值为 20 的元素slice[0] = append(slice[0], 20)
```

Go语言里使用 append() 函数处理追加的方式很简明，先增长切片，再将新的整型切片赋值给外层切片的第一个元素，当上面代码中的操作完成后，再将切片复制到外层切片的索引为 0 的元素，如下图所示。

![组合切片的切片](images\组合切片的切片.png)

即便是这么简单的多维切片，操作时也会涉及众多的布局和值，在函数间这样传递数据结构会很复杂，不过切片本身结构很简单，可以用很小的成本在函数间传递。

##### len() 和 cap() 函数

切片是可索引的，并且可以由`len()` 方法获取长度。

切片提供了计算容量的方法 `cap() `可以测量切片最长可以达到多少。

```go
package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)

   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x) //len=3 cap=5 slice=[0 0 0]
}


/**********************************************/

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[0:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) //长度  3
	fmt.Println("cap(s) = ", cap(s)) //容量  5

	s = a[1:4:5]
	fmt.Println("s = ", s)           //从下标1开始，取4-1=3个
	fmt.Println("len(s) = ", len(s)) //长度  4-1
	fmt.Println("cap(s) = ", cap(s)) //容量  5-1
}
```



##### 切片截取

| 操作            | 含义                                                         |
| --------------- | ------------------------------------------------------------ |
| s[n]            | 切片s中索引位置为n的项                                       |
| s[:]            | 从切片s的索引位置0到len(s)-1处所获得的切片                   |
| s[low:]         | 从切片s的索引位置low到len(s)-1处所获得的切片                 |
| s[:high]        | 从切片s的索引位置0到high处所获得的切片，len=high             |
| s[low:high]     | 从切片s的索引位置low到high处所获得的切片，len=high-low       |
| s[low:high:max] | 从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low |
| len(s)          | 切片s的长度，总是<=cap(s)                                    |
| cap(s)          | 切片s的容量，总是>=len(s)                                    |

示例说明：

```go
 array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
```

 

| 操作        | 结果                  | len  | cap  | 说明            |
| ----------- | --------------------- | ---- | ---- | --------------- |
| array[:6:8] | [0 1 2 3 4 5]         | 6    | 8    | 省略 low        |
| array[5:]   | [5 6 7 8 9]           | 5    | 5    | 省略 high、 max |
| array[:3]   | [0 1 2]               | 3    | 10   | 省略 high、 max |
| array[:]    | [0 1 2 3 4 5 6 7 8 9] | 10   | 10   | 全部省略        |

```go
package main

import "fmt"

func main() {
   /* 创建切片 */
   numbers := []int{0,1,2,3,4,5,6,7,8}  
   printSlice(numbers) //len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

   /* 打印原始切片 */
   fmt.Println("numbers ==", numbers) //numbers == [0 1 2 3 4 5 6 7 8]

   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4]) //numbers[1:4] == [1 2 3]

   /* 默认下限为 0*/
   fmt.Println("numbers[:3] ==", numbers[:3]) //numbers[:3] == [0 1 2]

   /* 默认上限为 len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:]) //numbers[4:] == [4 5 6 7 8]

   numbers1 := make([]int,0,5)
   printSlice(numbers1) //len=0 cap=5 slice=[]

   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
   number2 := numbers[:2]
   printSlice(number2) //len=2 cap=9 slice=[0 1]

   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   number3 := numbers[2:5]
   printSlice(number3) //len=3 cap=7 slice=[2 3 4]

}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

##### 删除元素

Go语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素，根据要删除元素的位置有三种情况，分别是从开头位置删除、从中间位置删除和从尾部删除，其中删除切片尾部的元素速度最快。

###### 从开头位置删除

删除开头的元素可以直接移动数据指针：

```go
a = []int{1, 2, 3}a = a[1:] // 删除开头1个元素a = a[N:] // 删除开头N个元素
```

也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：

```go
a = []int{1, 2, 3}a = append(a[:0], a[1:]...) // 删除开头1个元素a = append(a[:0], a[N:]...) // 删除开头N个元素
```

还可以用 copy() 函数来删除开头的元素：

```go
a = []int{1, 2, 3}a = a[:copy(a, a[1:])] // 删除开头1个元素a = a[:copy(a, a[N:])] // 删除开头N个元素
```

###### 从中间位置删除

对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用 append 或 copy 原地完成：

```go
a = []int{1, 2, 3, ...}a = append(a[:i], a[i+1:]...) // 删除中间1个元素a = append(a[:i], a[i+N:]...) // 删除中间N个元素a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素
```

###### 从尾部删除

```go
a = []int{1, 2, 3}a = a[:len(a)-1] // 删除尾部1个元素a = a[:len(a)-N] // 删除尾部N个元素
```


删除开头的元素和删除尾部的元素都可以认为是删除中间元素操作的特殊情况，下面来看一个示例。

【示例】删除切片指定位置的元素。

```go
package main

import "fmt"

func main() {    
    seq := []string{"a", "b", "c", "d", "e"}    
    // 指定删除位置
    index := 2    
    // 查看删除位置之前的元素和之后的元素    
    fmt.Println(seq[:index], seq[index+1:])    
    // 将删除点前后的元素连接起来    
    seq = append(seq[:index], seq[index+1:]...)    
    fmt.Println(seq)}
                                     
//代码输出结果：
[a b] [d e]
[a b d e]                                    
```

代码说明如下：

- 第 1 行，声明一个整型切片，保存含有从 a 到 e 的字符串。
- 第 4 行，为了演示和讲解方便，使用 index 变量保存需要删除的元素位置。
- 第 7 行，seq[:index] 表示的就是被删除元素的前半部分，值为 [1 2]，seq[index+1:] 表示的是被删除元素的后半部分，值为 [4 5]。
- 第 10 行，使用 append() 函数将两个切片连接起来。
- 第 12 行，输出连接好的新切片，此时，索引为 2 的元素已经被删除。

代码的删除过程可以使用下图来描述。

![删除过程](images\删除过程.png)

Go语言中删除切片元素的本质是，以被删除元素为分界点，将前后两个部分的内存重新连接起来。

**提示**

连续容器的元素删除无论在任何语言中，都要将删除点前后的元素移动到新的位置，随着元素的增加，这个过程将会变得极为耗时，因此，当业务需要大量、频繁地从一个切片中删除元素时，如果对性能要求较高的话，就需要考虑更换其他的容器了（如双链表等能快速从删除点删除元素）。

##### 切片和底层数组关系

```go
package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//新切片
	s1 := a[2:5] //从a[2]开始，取3个元素
	s1[1] = 666
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)

	//另外新切片
	s2 := s1[2:7]
	s2[2] = 777
	fmt.Println("s2 = ", s2)
	fmt.Println("a = ", a)
}

```

##### 内建函数-append()

`append`函数向 slice 尾部添加数据，返回新的 slice 对象：



append函数的使用

```go
package main

import "fmt"

func main() {
	s1 := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	//在原切片的末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)

	s2 := []int{1, 2, 3}
	fmt.Println("s2 = ", s2)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Println("s2 = ", s2)
}
```

append函数会智能地底层数组的容量增长，一旦超过原底层数组容量，通常以2倍容量重新分配底层数组，并复制原来的数据：

```go
package main

import "fmt"

func main() {
    
	//如果超过原来的容量，通常以2倍容量扩容
    
	s := make([]int, 0, 1) //容量为1
    
	oldCap := cap(s)
    
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap: %d ===> %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}
```



##### 内建函数-copy() 

函数 `copy `在两个 slice 间复制数据，复制⻓度以 len 小的为准，两个 slice 可指向同⼀底层数组。

```go
package main

import "fmt"

func main() {
    
   var numbers []int
   printSlice(numbers) ////len=0 cap=0 slice=[]

   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers) //len=1 cap=1 slice=[0]

   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers) //len=2 cap=2 slice=[0 1]

   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
   printSlice(numbers) //len=5 cap=6 slice=[0 1 2 3 4]

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)   //len=5 cap=12 slice=[0 1 2 3 4]
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

GOROOT/src/runtime/slice.go源码，其中扩容相关代码如下：

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

- 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
- 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
- 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
- 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如`int`和`string`类型的处理方式就不一样。

##### 切片做函数参数

在做函数调用时，slice 按引用传递，array 按值传递：

```go
package main

import "fmt"

func main(){
  changeSliceTest()
}

func changeSliceTest() {
    arr1 := []int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    arr3 := [3]int{1, 2, 3}

    fmt.Println("before change arr1, ", arr1) //before change arr1,  [1 2 3]
    changeSlice(arr1) // slice 按引用传递
    fmt.Println("after change arr1, ", arr1) //after change arr1,  [9999 2 3]

    fmt.Println("before change arr2, ", arr2) //before change arr2,  [1 2 3]
    changeArray(arr2) // array 按值传递
    fmt.Println("after change arr2, ", arr2) //after change arr2,  [1 2 3]

    fmt.Println("before change arr3, ", arr3) //before change arr3,  [1 2 3]
    changeArrayByPointer(&arr3) // 可以显式取array的 指针
    fmt.Println("after change arr3, ", arr3) //after change arr3,  [6666 2 3]
}

func changeSlice(arr []int) {
    arr[0] = 9999
}

func changeArray(arr [3]int) {
    arr[0] = 6666
}

func changeArrayByPointer(arr *[3]int) {
    arr[0] = 6666
}
```

##### 切片内部结构

```go
struct Slice
{   
    byte*    array;       // actual data
    uintgo    len;        // number of elements
    uintgo    cap;        // allocated number of elements

};
```

第一个字段表示 array 的指针，是真实数据的指针第二个是表示 slice 的长度，第三个是表示 slice 的容量。

所以 `unsafe.Sizeof(切片)`永远都是 24。

当把 slice 作为参数，本身传递的是值，但其内容就 **byte\* array**，实际传递的是引用，所以可以在函数内部修改，但如果对 slice 本身做 append，而且导致 slice 进行了扩容，实际扩容的是函数内复制的一份切片，对于函数外面的切片没有变化。

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
    
    slice_test := []int{1, 2, 3, 4, 5}
    
    fmt.Println(unsafe.Sizeof(slice_test)) //24
    
    //main:[]int{1, 2, 3, 4, 5},5,5
    fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
    
    //slice_value:[]int{1, 100, 3, 4, 5, 6},6,10
    slice_value(slice_test)
    
    //main:[]int{1, 100, 3, 4, 5},5,5
    fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
    
    //slice_ptr:[]int{1, 100, 3, 4, 5, 7},6,10
    slice_ptr(&slice_test)
    
    //main:[]int{1, 100, 3, 4, 5, 7},6,10
    fmt.Printf("main:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
    
    //24
    fmt.Println(unsafe.Sizeof(slice_test))
}

func slice_value(slice_test []int) {
    slice_test[1] = 100                // 函数外的slice确实有被修改
    slice_test = append(slice_test, 6) // 函数外的不变
    fmt.Printf("slice_value:%#v,%#v,%#v\n", slice_test, len(slice_test), cap(slice_test))
}

func slice_ptr(slice_test *[]int) { // 这样才能修改函数外的slice
    *slice_test = append(*slice_test, 7)
    fmt.Printf("slice_ptr:%#v,%#v,%#v\n", *slice_test, len(*slice_test), cap(*slice_test))
}
```

##### 猜数字游戏

```go
package main

import "fmt"
import "math/rand"
import "time"

func CreatNum(p *int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())

	var num int
	for {
		num = rand.Intn(10000) //一定是4位数
		if num >= 1000 {
			break
		}
	}

	//fmt.Println("num = ", num)

	*p = num

}

func GetNum(s []int, num int) {
	s[0] = num / 1000       //取千位
	s[1] = num % 1000 / 100 //取百位
	s[2] = num % 100 / 10   //取百位
	s[3] = num % 10         //取个位
}

func OnGame(randSlice []int) {
	var num int
	keySlice := make([]int, 4)

	for {
		for {
			fmt.Printf("请输入一个4位数：")
			fmt.Scan(&num)

			// 999 < num < 10000
			if 999 < num && num < 10000 {
				break
			}

			fmt.Println("请输入的数不符合要求")
		}
		//fmt.Println("num = ", num)
		GetNum(keySlice, num)
		//fmt.Println("keySlice = ", keySlice)

		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了一点\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了一点\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}

		if n == 4 { //4位都猜对了
			fmt.Println("全部猜对!!!")
			break //跳出循环
		}
	}
}

func main() {
	var randNum int

	//产生一个4位的随机数
	CreatNum(&randNum)
	//fmt.Println("randNum: ", randNum)

	randSlice := make([]int, 4)
	//保存这个4位数的每一位
	GetNum(randSlice, randNum)
	//fmt.Println("randSlice = ", randSlice)

	/*
		n1 := 1234 / 1000 //取商
		//(1234 % 1000) //取余数，结果为234   234/100取商得到2
		n2 := 1234 % 1000 / 100
		fmt.Println("n1 = ", n1)
		fmt.Println("n2 = ", n2)
	*/

	OnGame(randSlice) //游戏
}
```



#### Map

##### 概述

哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同的，然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。

在Go语言中，`一个map就是一个哈希表的引用`，map类型可以写为`map[K]V`，其中K和V分别对应key和value。**map中所有的key都有相同的类型，所有的value也有着相同的类型，但是key和value之间可以是不同的数据类型。**其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。虽然浮点数类型也是支持相等运算符比较的，但是将浮点数用做key类型则是一个坏的想法，最坏的情况是可能出现的NaN和任何浮点数都不相等。对于V对应的value数据类型则没有任何的限制。



`Map` 是一种`无序`的`键值对`的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

![map](images\map.png)



map格式为：

```go
map[keyType]valueType
```

在一个map里所有的键都是`唯一`的，而且必须是支持`==`和`!=`操作符的类型，切片、函数以及包含切片的结构类型这些类型由于具有`引用语义`，不能作为映射的键，使用这些类型会造成编译错误：

```go
 dict := map[ []string ]int{} //err, invalid map key type []string
```

map值可以是任意类型，没有限制。**map里所有键的数据类型必须是相同的**，值也必须如何，但键和值的数据类型可以不相同。

> 注意：map是无序的，我们无法决定它的返回顺序，所以，每次打印结果的顺利有可能不同。

```go
package main

import "fmt"

func main() {
    
	//定义一个变量， 类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len，没有cap
	fmt.Println("len = ", len(m1))

	//可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	//可以通过make创建，可以指定长度，只是指定了容量，但是里面却是一个数据也没有
	m3 := make(map[int]string, 2)
	m3[1] = "mike" //元素的操作
	m3[2] = "go"
	m3[3] = "c++"

	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))

	//初始化
	//键值是唯一的
	m4 := map[int]string{1: "mike", 2: "go", 3: "c++"}
	fmt.Println("m4 = ", m4)
}
```



##### 创建和初始化

可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

```go
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable := make(map[key_data_type]value_data_type)


//用map字面值的语法创建map，同时还可以指定一些最初的key/value：
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
//这相当于
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34


//创建空的map的表达式
map[string]int{}

//Map中的元素通过key对应的下标语法访问：
ages["alice"] = 32
fmt.Println(ages["alice"]) // "32"

//内置的delete函数可以删除元素：
delete(ages, "alice") // remove element ages["alice"]

//所有这些操作是安全的，即使这些元素不在map中也没有关系；如果一个查找失败将返回value类型对应的零值，例如，即使map中不存在“bob”下面的代码也可以正常工作，因为ages["bob"]失败时将返回0。
ages["bob"] = ages["bob"] + 1 // happy birthday!

//而且x += y和x++等简短赋值语法也可以用在map上，所以上面的代码可以改写成
ages["bob"] += 1
//更简单的写法
ages["bob"]++

//map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作：
_ = &ages["bob"] // compile error: cannot take address of map element
//禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。
```

```go
package main

import "fmt"

func main() {
    var countryCapitalMap map[string]string /*创建集合 */
    countryCapitalMap = make(map[string]string)

    /* map插入key - value对,各个国家对应的首都 */
    countryCapitalMap [ "France" ] = "巴黎"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India " ] = "新德里"

    /*使用键输出地图值 */
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    /*查看元素在集合中是否存在 */
    capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("American 的首都是", capital)
    } else {
        fmt.Println("American 的首都不存在")
    }
}

//以上实例运行结果为：
France 首都是 巴黎
Italy 首都是 罗马
Japan 首都是 东京
India  首都是 新德里
American 的首都不存在
```

##### 常用操作-赋值

```go
package main

import "fmt"

func main() {
    
	m1 := map[int]string{1: "mike", 2: "yoyo"}
    
	//赋值，如果已经存在的key值，修改内容
	fmt.Println("m1 = ", m1)
    
	m1[1] = "c++"
	m1[3] = "go" //追加，map底层自动扩容，和append类似
    
	fmt.Println("m1 = ", m1)
}
```



##### 常用操作-遍历

Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同。

```go
package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}

	//第一个返回值为key, 第二个返回值为value, 遍历结果是无序的
	for key, value := range m {
		fmt.Printf("%d =======> %s\n", key, value)
	}

	//如何判断一个key值是否存在
	//第一个返回值为key所对应的value, 第二个返回值为key是否存在的条件，存在ok为true
	value, ok := m[0]
	if ok == true {
		fmt.Println("m[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}
}
```



##### 常用操作-删除delete() 

即使这些元素不在map中也没有关系；如果一个查找失败将返回value类型对应的零值

```go
package main //必须有个main包

import "fmt"

func main() {
    
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	fmt.Println("m = ", m)

	delete(m, 1) //删除key为1的内容
    
	fmt.Println("m = ", m)
}
```

##### map做函数参数

在函数间传递映射并不会制造出该映射的一个副本，不是值传递，而是引用传递：

```go
package main //必须有个main包

import "fmt"

func test(m map[int]string) {
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	fmt.Println("m = ", m)

	test(m) //在函数内部删除某个key

	fmt.Println("m = ", m)
}
```



#### 结构体

##### 概述

**结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。**

结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。用结构体的经典案例处理公司的员工信息，每个员工信息包含一个唯一的员工编号、员工的名字、家庭住址、出生日期、工作岗位、薪资、上级领导等等。所有的这些信息都需要绑定到一个实体中，可以作为一个整体单元被复制，作为函数的参数或返回值，或者是被存储到数组中，等等。

下面两个语句声明了一个叫Employee的命名的结构体类型，并且声明了一个Employee类型的变量dilbert：

```go
type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

var dilbert Employee

//dilbert结构体变量的成员可以通过点操作符访问，因为dilbert是一个变量，它所有的成员也同样是变量，我们可以直接对每个成员赋值：
dilbert.Salary -= 5000 // demoted, for writing too few lines of code

//或者是对成员取地址，然后通过指针访问：
position := &dilbert.Position
*position = "Senior " + *position // promoted, for outsourcing to Elbonia

//点操作符也可以和指向结构体的指针一起工作：
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
//相当于下面语句
(*employeeOfTheMonth).Position += " (proactive team player)"
```

##### 结构体初始化

结构体定义需要使用 `type` 和 `struct` 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

```go
variable_name := structure_variable_type {value1, value2...valuen}
//或
variable_name := structure_variable_type { key1: value1, ..., keyn: valuen}
```

###### 普通变量

结构体普通变量初始化

```go
package main

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
    
	//顺序初始化，每个成员必须初始化
	var s1 Student = Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("s1 = ", s1)

	//指定成员初始化，没有初始化的成员，自动赋值为 0 或 空
	s2 := Student{name: "mike", addr: "bj"}
	fmt.Println("s2 = ", s2)
}
```

###### 指针变量

结构体指针变量初始化

```go
package main

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
    
	//顺序初始化，每个成员必须初始化, 别忘了&
	var p1 *Student = &Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("p1 = ", p1)

	//指定成员初始化，没有初始化的成员，自动赋值为 0 或 空
	p2 := &Student{name: "mike", addr: "bj"}
	fmt.Printf("p2 type is %T\n", p2)
	fmt.Println("p2 = ", p2)
}
```

##### 结构体成员的使用

###### 普通变量

```go
package main

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	//定义一个结构体普通变量
	var s Student

	//操作成员，需要使用点(.)运算符
	s.id = 1
	s.name = "mike"
	s.sex = 'm' //字符
	s.age = 18
	s.addr = "bj"
	fmt.Println("s = ", s)
}
```



###### 指针变量

```go
package main //必须有个main包

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	//1、指针有合法指向后，才操作成员
	//先定义一个普通结构体变量
	var s Student
	//在定义一个指针变量，保存s的地址
	var p1 *Student
	p1 = &s

	//通过指针操作成员  p1.id 和(*p1).id完全等价，只能使用.运算符
	p1.id = 1
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.age = 18
	p1.addr = "bj"
	fmt.Println("p1 = ", p1)

	//2、通过new申请一个结构体
	p2 := new(Student)
	p2.id = 1
	p2.name = "mike"
	p2.sex = 'm'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2 = ", p2)
}
```



##### 结构体比较

如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用 == 或 != 运算符进行比较，但不支持 > 或 < 。

```go
package main //必须有个main包

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	s1 := Student{1, "mike", 'm', 18, "bj"}
	s2 := Student{1, "mike", 'm', 18, "bj"}
	s3 := Student{2, "mike", 'm', 18, "bj"}
	fmt.Println("s1 == s2 ", s1 == s2)
	fmt.Println("s1 == s3 ", s1 == s3)

	//同类型的2个结构体变量可以相互赋值
	var tmp Student
	tmp = s3
	fmt.Println("tmp = ", tmp)
}
```



#####  结构体作为函数参数

```go
package main

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func test02(p *Student) {
	p.id = 666
}

func main() {
	s := Student{1, "mike", 'm', 18, "bj"}

	test02(&s) //地址传递（引用传递），形参可以改实参
	fmt.Println("main: ", s)

}

func test01(s Student) {
	s.id = 666
	fmt.Println("test01: ", s)
}

func main01() {
	s := Student{1, "mike", 'm', 18, "bj"}

	test01(s) //值传递，形参无法改实参
	fmt.Println("main: ", s)
}
```



##### 可见性

Go语言对关键字的增加非常吝啬，其中没有private、 protected、 public这样的关键字。

要使某个符号对其他包（package）可见（即可以访问），需要将该符号定义为以大写字母
开头。

目录结构：

```go
|--src
	|--test
		|--test.go
	|--main.go
```

test.go示例代码如下：

```go
package test

import "fmt"

//如果首字母是小写，只能在同一个包里使用
type stu struct {
	id int
}

type Stu struct {
	//id int //如果首字母是小写，只能在同一个包里使用
	Id int
}

//如果首字母是小写，只能在同一个包里使用
func myFunc() {
	fmt.Println("this is myFunc")
}

func MyFunc() {
	fmt.Println("this is MyFunc -=======")
}
```

main.go示例代码如下：

```go
package main

import "test"
import "fmt"

func main() {
    
	//包名.函数名
	test.MyFunc()

	//包名.结构体里类型名
	var s test.Stu
	s.Id = 666
	fmt.Println("s = ", s)
}
```



## 函数

函数是基本的代码块，用于执行一个任务。

Go 语言最少有个 `main()` 函数。

函数声明告诉了编译器`函数名称`，`返回类型`，和`参数`。

Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

### 定义格式

函数构成代码执行的逻辑结构。在Go语言中，函数的基本组成为：关键字`func`、函数名、参数列表、返回值、函数体和返回语句。

```go
//Go 语言函数定义格式如下：
func FuncName( [参数列表] ) [(o1 type1, o2 type2/*返回类型*/)] {
    
    //函数体
    
    return v1, v2 //返回多个值
}
```

函数定义说明：

1. func：函数由关键字 func 开始声明
2. FuncName：函数名称，根据约定，函数名首字母小写即为private，大写即为public
3. 参数列表：函数可以有0个或多个参数，参数格式为：变量名 类型，如果有多个参数通过逗号分隔，不支持默认参数
4. 返回类型：

①　上面返回值声明了两个变量名o1和o2(命名返回参数)，这个不是必须，可以只有类型没有变量名

②　如果只有一个返回值且不声明返回值变量，那么你可以省略，包括返回值的括号

③　如果没有返回值，那么就直接省略最后的返回信息

④　如果有返回值， 那么必须在函数的内部添加`return`语句



````go
/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 声明局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}
````

### 函数调用

当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。

调用函数，向函数传递参数，并返回值，例如：

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
   var ret int

   /* 调用函数并返回最大值 */
   ret = max(a, b)

   fmt.Printf( "最大值是 : %d\n", ret ) //最大值是 : 200
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 定义局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}
```

### 自定义函数

#### 无参无返回值

```go
package main //必须

import "fmt"

func main() {
	//无参无返回值函数的调用： 函数名()
	MyFunc()
}

//无参无返回值函数的定义
func MyFunc() {
	a := 666
	fmt.Println("a = ", a)
}
```

#### 有参无返回值

##### 普通参数列表

```go
package main //必须

import "fmt"

//有参无返回值函数的定义， 普通参数列表
//定义函数时， 在函数名后面()定义的参数叫形参
//参数传递，只能由实参传递给形参，不能反过来， 单向传递
func MyFunc01(a int) {
	//a = 111
	fmt.Println("a = ", a)
}

func MyFunc02(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc03(a, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc04(a int, b string, c float64) {
}

func MyFunc05(a, b string, c float64, d, e int) {
}

func MyFunc06(a string, b string, c float64, d int, e int) {
}

func main() {
	//有参无返回值函数调用：  函数名(所需参数)
	//调用函数传递的参数叫实参
	MyFunc01(666)

	MyFunc02(666, 777)

}
```



#### 不定参数列表

###### 不定参数类型

不定参数是指函数传入的参数个数为不定数量。为了做到这点，首先需要将函数定义为接受不定参数类型：

```go
package main //必须

import "fmt"

func MyFunc01(a int, b int) { //固定参数

}

//...int类型这样的类型， ...type不定参数类型
//注意：不定参数，一定（只能）放在形参中的最后一个参数
func MyFunc02(args ...int) { //传递的实参可以是0或多个
	fmt.Println("len(args) = ", len(args)) //获取用户传递参数的个数
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}

	fmt.Println("==========================================")

	//返回2个值，第一个是下标，第二个是下标所对应的数
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

func main01() {
    
	//MyFunc01(666, 777)

	MyFunc02()
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++")
	MyFunc02(1)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++")
	MyFunc02(1, 2, 3)
}

//固定参数一定要传参，不定参数根据需求传递
func MyFunc03(a int, args ...int) {
}

//注意：不定参数，一定（只能）放在形参中的最后一个参数
//func MyFunc04(args ...int, a int) {
//}

func main() {
	MyFunc03(111, 1, 2, 3)
}

```



###### 不定参数的传递

```go
package main //必须

import "fmt"

func myfunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}

}

func myfunc2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}

}

func test(args ...int) {
    
	//全部元素传递给myfunc
	//myfunc(args...)

	//只想把后2个参数传递给另外一个函数使用
	myfunc2(args[:2]...) //args[0]~args[2]（不包括数字args[2]）， 传递过去
	myfunc2(args[2:]...) //从args[2]开始(包括本身)，把后面所有元素传递过去
}

func main() {
	test(1, 2, 3, 4)
}
```

#### 无参有返回值

##### 一个返回值

```go
package main //必须

import "fmt"

//无参有返回值：只有一个返回值
//有返回值的函数需要通过return中断函数，通过return返回
func myfunc01() int {
	return 666
}

//给返回值起一个变量名，go推荐写法
func myfunc02() (result int) {

	return 666
}

//给返回值起一个变量名，go推荐写法
//常用写法
func myfunc03() (result int) {

	result = 666
	return
}

func main() {
	//无参有返回值函数调用
	var a int
	a = myfunc01()
	fmt.Println("a = ", a)

	b := myfunc01()
	fmt.Println("b = ", b)

	c := myfunc03()
	fmt.Println("c = ", c)
}
```



##### 多个返回值

```go
package main //必须

import "fmt"

//多个返回值
func myfunc01() (int, int, int) {
	return 1, 2, 3
}

//go官方推荐写法
func myfunc02() (a int, b int, c int) {
	a, b, c = 111, 222, 333
	return
}

func myfunc03() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}

func main() {
	//函数调用
	a, b, c := myfunc02()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
```

#### 有参有返回值

```go
package main //必须

import "fmt"

//函数定义
func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return //有返回值的函数，必须通过return返回
}

func main() {
	max, min := MaxAndMin(10, 20)
	fmt.Printf("max = %d, min = %d\n", max, min)

	//通过匿名变量丢弃某个返回值
	a, _ := MaxAndMin(10, 20)
	fmt.Printf("a = %d\n", a)
}
```

### 递归函数

#### 概述

递归指函数可以直接或间接的调用自身。

```go
//语法格式如下：
func recursion() {
   recursion() /* 函数调用自身 */
}

func main() {
   recursion()
}
```

Go 语言支持递归。但我们在使用递归时，开发者需要`设置退出条件`，否则递归将陷入无限循环中。

递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。

#### 普通函数的调用流程

```go
package main //必须

import "fmt"

func funcc(c int) {
	fmt.Println("c = ", c)
}

func funcb(b int) {

	funcc(b - 1)
	fmt.Println("b = ", b)
}

func funca(a int) {
	funcb(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	funca(3) //函数调用
	fmt.Println("main")
}
```



#### 函数递归调用的流程

```go
package main //必须

import "fmt"

func test(a int) {
	if a == 1 { //函数终止调用的条件，非常重要
		fmt.Println("a = ", a)
		return //终止函数调用
	}

	//函数调用自身
	test(a - 1)

	fmt.Println("a = ", a)
}

func main() {
	test(3)
	fmt.Println("main")
}
```

#### 实例

##### 阶乘

```go
package main

import "fmt"

func Factorial(n uint64)(result uint64) {
    if (n > 0) {
        result = n * Factorial(n-1)
        return result
    }
    return 1
}

func main() {  
    var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i))) //15 的阶乘是 1307674368000
}
```

##### 斐波那契数列

```go
package main

import "fmt"

func fibonacci(n int) int {
  if n < 2 {
   return n
  }
  return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
    var i int
    for i = 0; i < 10; i++ {
       fmt.Printf("%d\t", fibonacci(i))
    }
}

//以上实例执行输出结果为：
0    1    1    2    3    5    8    13    21    34



//更好的一种 fibonacci 实现，用到多返回值特性，降低复杂度：
func fibonacci2(n int) (int,int) {
  if n < 2 {
    return 0,n
  }
  a,b := fibonacci2(n-1)
  return b,a+b
}


func fibonacci(n int) int {
  a,b := fibonacci2(n)
  return b
}
```

##### 求平方根

```go
//原理: 计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：

z -= (z*z - x) / (2*z)
//重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。


package main
import "fmt"

func sqrt(x float64,i float64) (float64,float64){
    remain:=(i*i-x)/(2*i);
    i=i-remain
    if(remain>0){
        return sqrt(x,i);
    }else{ 
        return i,remain  
    }
}
func get_sqrt(x float64) float64{   
    i,_ :=sqrt(x,x);  
    return i;
}
func main(){  
    fmt.Println(get_sqrt(2))
    fmt.Println(get_sqrt(3))
}



//求平方根算法复杂度改成 O(1)：
package main
import "fmt"
import "unsafe"

func get_sqrt(x float32) float32{   
    xhalf := 0.5*x;
    var i int32 = *(*int32)(unsafe.Pointer(&x));
    i = 0x5f375a86 - (i>>1);
    x = *(*float32)(unsafe.Pointer(&i));
    
    x = x * (1.5 - xhalf*x*x);
    x = x * (1.5 - xhalf*x*x);
    x = x * (1.5 - xhalf*x*x);
    
    return 1/x;
}

func main(){  
    fmt.Println(get_sqrt(4))
}
```



### 函数类型

```go
package main //必须

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型， 通过type给一个函数类型起名
//FuncType它是一个函数类型
type FuncType func(int, int) int  //没有函数名字，没有{}

func main() {
    
	var result int
	result = Add(1, 1) //传统调用方式
	fmt.Println("result = ", result)

	//声明一个函数类型的变量，变量名叫fTest
	var fTest FuncType
	fTest = Add            //是变量就可以赋值
	result = fTest(10, 20) //等价于Add(10, 20)
	fmt.Println("result2 = ", result)

	fTest = Minus
	result = fTest(10, 5) //等价于Minus(10, 5)
	fmt.Println("result3 = ", result)
}

```



### 回调函数

回调函数，函数有一个参数是函数类型，这个函数就是回调函数

```go
package main //必须

import "fmt"

type FuncType func(int, int) int

//实现加法
func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
//计算器，可以进行四则运算
//多态，多种形态，调用同一个接口，不同的表现，可以实现不同表现，加减乘除
//现有想法，后面再实现功能
func Calc(a, b int, fTest FuncType) (result int) {
    
	fmt.Println("Calc")
    
	result = fTest(a, b) //这个函数还没有实现
	//result = Add(a, b) //Add()必须先定义后，才能调用
	return
}

func main() {
	a := Calc(1, 1, Mul)
	fmt.Println("a = ", a)
}
```



### 值传递和引用传递

函数如果使用参数，该变量可称为函数的形参。

形参就像定义在函数体内的局部变量。

调用函数，可以通过两种方式来传递参数：

| 传递类型 | 描述                                                         |
| :------- | :----------------------------------------------------------- |
| 值传递   | 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。 |
| 引用传递 | 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。 |

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

#### 值传递值

传递是指在调用函数时将`实际参数复制一份传递到函数中`，这样在函数中如果对参数进行修改，将不会影响到实际参数。

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

```go
//使用值传递来调用 swap() 函数：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200

   fmt.Printf("交换前 a 的值为 : %d\n", a )
   fmt.Printf("交换前 b 的值为 : %d\n", b )

   /* 通过调用函数来交换值 */
   swap(a, b)

   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )
}


/* 定义相互交换值的函数 */
func swap(x, y int) int {
   var temp int

   temp = x /* 保存 x 的值 */
   x = y    /* 将 y 值赋给 x */
   y = temp /* 将 temp 值赋给 y*/

   return temp;
}

//以下代码执行结果为：
交换前 a 的值为 : 100
交换前 b 的值为 : 200
交换后 a 的值 : 100
交换后 b 的值 : 200
```

#### 引用传递值

引用传递是指在调用函数时将实际参数的`地址`传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递：

```go
//使用引用传递来调用 swap() 函数：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200

   fmt.Printf("交换前，a 的值 : %d\n", a )
   fmt.Printf("交换前，b 的值 : %d\n", b )

   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap(&a, &b)

   fmt.Printf("交换后，a 的值 : %d\n", a )
   fmt.Printf("交换后，b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}

//以上代码执行结果为：
交换前，a 的值 : 100
交换前，b 的值 : 200
交换后，a 的值 : 200
交换后，b 的值 : 100
```



### 匿名函数与闭包

Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

```go
package main //必须

import "fmt"

func main() {
	a := 10
	str := "mike"

	//匿名函数，没有函数名字, 函数定义，还没有调用
	f1 := func() { //:= 自动推导类型
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	f1()

	//给一个函数类型起别名
	type FuncType func() //函数没有参数，没有返回值
	//声明变量
	var f2 FuncType
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}() //后面的()代表调用此匿名函数

	//带参数的匿名函数
	f3 := func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	f3(1, 2)

	//定义匿名函数，同时调用
	func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(10, 20)

	//匿名函数，有参有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}

		return
	}(10, 20)

	fmt.Printf("x = %d, y = %d\n", x, y)

}
```

闭包捕获外部变量的特点

```go
package main //必须

import "fmt"

func main() {
	a := 10
	str := "mike"

	func() {
		//闭包以引用方式捕获外部变量
		a = 666
		str = "go"
		fmt.Printf("内部：a = %d, str = %s\n", a, str)
	}() //()代表直接调用

	fmt.Printf("外部：a = %d, str = %s\n", a, str)

}
```

闭包的特点

```go
package main //必须

import "fmt"

//函数的返回值是一个匿名函数，返回一个函数类型
func test02() func() int {
	var x int //没有初始化，值为0

	return func() int {
		x++
		return x * x
	}
}

func main() {

	//返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	//它不关心这些捕获了的变量和常量是否已经超出了作用域
	//所以只有闭包还在使用它，这些变量就还会存在。
	f := test02()
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9
	fmt.Println(f()) //16
	fmt.Println(f()) //25

}

func test01() int {
	//函数被调用时，x才分配空间，才初始化为0
	var x int //没有初始化，值为0
	x++
	return x * x //函数调用完毕，x自动释放
}

func main01() {
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
}

```



```go
package main

import "fmt"

func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
   }
}

func main(){
   /* nextNumber 为一个函数，函数 i 为 0 */
   nextNumber := getSequence()  

   /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   
   /* 创建新的函数 nextNumber1，并查看结果 */
   nextNumber1 := getSequence()  
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}
//以上代码执行结果为：
1
2
3
1
2


//带参数的闭包函数调用:
package main

import "fmt"
func main() {
    add_func := add(1,2)
    fmt.Println(add_func())
    fmt.Println(add_func())
    fmt.Println(add_func())
}

// 闭包使用方法
func add(x1, x2 int) func()(int,int)  {
    i := 0
    return func() (int,int){
        i++
        return i,x1+x2
    }
}

//闭包带参数补充:
package main
import "fmt"
func main() {
    add_func := add(1,2)
    fmt.Println(add_func(1,1))
    fmt.Println(add_func(0,0))
    fmt.Println(add_func(2,2))
} 
// 闭包使用方法
func add(x1, x2 int) func(x3 int,x4 int)(int,int,int)  {
    i := 0
    return func(x3 int,x4 int) (int,int,int){ 
       i++
       return i,x1+x2,x3+x4
    }
}

//闭包带参数继续补充：
package main

import "fmt"

// 闭包使用方法，函数声明中的返回值(闭包函数)不用写具体的形参名称
func add(x1, x2 int) func(int, int) (int, int, int) {
  i := 0
  return func(x3, x4 int) (int, int, int) {
    i += 1
    return i, x1 + x2, x3 + x4
  }
}

func main() {
  add_func := add(1, 2)
  fmt.Println(add_func(4, 5))
  fmt.Println(add_func(1, 3))
  fmt.Println(add_func(2, 2)) 
}
```

### 延迟调用defer

####  defer作用

关键字` defer `⽤于延迟一个函数或者方法（或者当前所创建的匿名函数）的执行。注意，defer语句只能出现在函数或方法的内部。

```go
package main //必须

import "fmt"

func main() {
	//defer延迟调用，main函数结束前调用
	defer fmt.Println("bbbbbbbbbbbbb")

	fmt.Println("aaaaaaaaaaaaaaa")
}
```

defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。释放资源的defer应该直接跟在请求资源的语句后。

#### 多个defer执行顺序

如果一个函数中有多个defer语句，它们会以`LIFO`（后进先出）的顺序执行。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执⾏。

```go
package main //必须

import "fmt"

func test(x int) {
	result := 100 / x

	fmt.Println("result = ", result)
}

func main() {

	defer fmt.Println("aaaaaaaaaaaaaa")

	defer fmt.Println("bbbbbbbbbbbbbb")

	//调用一个函数，导致内存出问题
	defer test(0)

	defer fmt.Println("ccccccccccccccc")
}
```

#### defer和匿名函数结合使用

```go
package main //必须

import "fmt"

func main() {
	a := 10
	b := 20

	//	defer func(a, b int) {
	//		fmt.Printf("a = %d, b = %d\n", a, b)
	//	}(a, b) //()代表调用此匿名函数, 把参数传递过去，已经先传递参数，只是没有调用

	defer func(a, b int) {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(10, 20) //()代表调用此匿名函数, 把参数传递过去，已经先传递参数，只是没有调用

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)
}

func main01() {
	a := 10
	b := 20

	defer func() {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}() //()代表调用此匿名函数

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)
}
```

### 获取命令行参数

```go
package main //必须

import "fmt"
import "os"

func main() {
	//接收用户传递的参数，都是以字符串方式传递
	list := os.Args

	n := len(list)
	fmt.Println("n = ", n)

	//xxx.exe a b
	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}
```



```go
package main

import (
    "fmt"
    "os"    //os.Args所需的包
)

func main() {
    args := os.Args //获取用户输入的所有参数

    //如果用户没有输入,或参数个数不够,则调用该函数提示用户
    if args == nil || len(args) < 2 {
        fmt.Println("err: xxx ip port")
        return
    }
    
    ip := args[1]   //获取输入的第一个参数
    port := args[2] //获取输入的第二个参数
    fmt.Printf("ip = %s, port = %s\n", ip, port)
}
```



## 工程管理

### 概述

在实际的开发工作中，直接调用编译器进行编译和链接的场景是少而又少，因为在工程中不会简单到只有一个源代码文件，且源文件之间会有相互的依赖关系。如果这样一个文件一个文件逐步编译，那不亚于一场灾难。 早期Go语言使用makefile作为临时方案，到了Go 1发布时引入了强大无比的Go命令行工具。

 Go命令行工具的革命性之处在于彻底消除了工程文件的概念，完全用`目录结构`和`包名`来推导工程结构和构建顺序。针对只有一个源文件的情况讨论工程管理看起来会比较多余，因为这可以直接用`go run`和`go build`搞定。



### 工作区

#### 工作区介绍

Go代码必须放在工作区中。工作区其实就是一个对应于特定工程的目录，它应包含3个子目录：src目录、pkg目录和bin目录。

`src`目录：用于以代码包的形式组织并保存Go源码文件。（比如：.go .c .h .s等）

`pkg`目录：用于存放经由go install命令构建安装后的代码包（包含Go库源码文件）的“.a”归档文件。

`bin`目录：与pkg目录类似，在通过go install命令完成安装后，保存由Go命令源码文件生成的可执行文件。

目录src用于包含所有的源代码，是Go命令行工具一个强制的规则，而pkg和bin则无需手动创建，如果必要Go命令行工具在构建过程中会自动创建这些目录。

需要特别注意的是，只有当环境变量`GOPATH`中只包含一个工作区的目录路径时，go install命令才会把命令源码安装到当前工作区的bin目录下。若环境变量GOPATH中包含多个工作区的目录路径，像这样执行go install命令就会失效，此时必须设置环境变量`GOBIN`。

#### GOPATH设置

为了能够构建这个工程，需要先把所需工程的根目录加入到环境变量`GOPATH`中。否则，即使处于同一工作目录(工作区)，代码之间也无法通过绝对代码包路径完成调用。

在实际开发环境中，工作目录往往有多个。这些工作目录的目录路径都需要添加至`GOPATH`。当有多个目录时，请注意分隔符，多个目录的时候Windows是`分号`，Linux系统是`冒号`，**当有多个GOPATH时，默认会将go get的内容放在第一个目录下。**

### 包

所有 Go 语言的程序都会组织成若干组文件，每组文件被称为一个包。这样每个包的代码都可以作为很小的复用单元，被其他项目引用。

一个包的源代码保存在一个或多个以.go为文件后缀名的源文件中，通常一个包所在目录路径的后缀是包的导入路径。



#### 自定义包

对于一个较大的应用程序，我们应该将它的功能性分隔成逻辑的单元，分别在不同的包里实现。我们创建的的自定义包最好放在`GOPATH`的`src`目录下（或者GOPATH src的某个子目录）。

在Go语言中，代码包中的源码文件名可以是任意的。但是，这些任意名称的源码文件都必须以包声明语句作为文件中的第一行，每个包都对应一个独立的名字空间：

```go
package calc
```

包中成员以名称⾸字母⼤⼩写决定访问权限：

`public`: ⾸字母⼤写，可被包外访问

`privat`: ⾸字母⼩写，仅包内成员可以访问



> **注意：同一个目录下不能定义不同的package。**

#### main包

在 Go 语言里，命名为` main `的包具有特殊的含义。 Go 语言的编译程序会试图把这种名字的包编译为二进制可执行文件。所有用 Go 语言编译的可执行程序都必须有一个名叫 main 的包。**一个可执行程序有且仅有一个 main 包。**

当编译器发现某个包的名字为 main 时，它一定也会发现名为 main()的函数，否则不会创建可执行文件。 `main()函数是程序的入口`，所以，如果没有这个函数，程序就没有办法开始执行。程序编译时，会使用声明 main 包的代码所在的目录的目录名作为二进制可执行文件的文件名。

#### main函数和init函数

Go里面有两个保留的函数：`init函数`（能够应用于所有的package）和`main函数`（只能应用于package main）。这两个函数在定义时`不能有任何的参数和返回值`。虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数。

Go程序会`自动调用init()和main()`，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。

每个包可以包含任意多个 init 函数，这些函数都会在程序执行开始的时候被调用。所有被编译器发现的 init 函数都会安排在 main 函数之前执行。 init 函数用在设置包、初始化变量或者其他要在程序运行前优先完成的引导工作。

程序的初始化和执行都起始于main包。如果main包还导入了其它的包，那么就会在编译时将它们依次导入。

有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。

当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数。

下图详细地解释了整个执行过程：

[![执行过程](images\执行过程.jpg)



示例代码目录结构：

![目录结构](E:\smile\go\images\目录结构.jpg)

main.go示例代码如下：

```go

// main.go
package main

import (
    "fmt"
    "test"
)

func main() {
    
    fmt.Println("main.go main() is called")

    test.Test()
}
```

test.go示例代码如下：

```go
//test.go
package test

import "fmt"

func init() {
    fmt.Println("test.go init() is called")
}

func Test() {
    fmt.Println("test.go Test() is called")
}
```

main.go, test.go相同目录

```go
 //test.go
package main

import "fmt"

func test() {
	fmt.Println("this is a test func")
}


//main.go,
package main //必须

func main() {
	test()
}

```



#### 导入包

导入包需要使用关键字import，它会告诉编译器你想引用该位置的包内的代码。包的路径可以是相对路径，也可以是绝对路径。

```go
//方法1
import "calc"
import "fmt"

//方法2
import (
    "calc"
    "fmt"
)
```

标准库中的包会在安装 Go 的位置找到。 Go 开发者创建的包会在 `GOPATH `环境变量指定的目录里查找。GOPATH 指定的这些目录就是开发者的个人工作空间。

如果编译器查遍 GOPATH 也没有找到要导入的包，那么在试图对程序执行 run 或者 build的时候就会出错。

**注意：如果导入包之后，未调用其中的函数或者类型将会报出编译错误。**

##### 点操作

```go
import (
    //这个点操作的含义是这个包导入之后在你调用这个包的函数时，可以省略前缀的包名
    . "fmt"
)

func main() {
    Println("hello go")
}
```

##### 别名操作

在导⼊时，可指定包成员访问⽅式，⽐如对包重命名，以避免同名冲突：

```go
import (
    io "fmt" //fmt改为为io
)

func main() {
    io.Println("hello go") //通过io别名调用
}
```

##### _(匿名)操作

有时，用户可能需要导入一个包，但是不需要引用这个包的标识符。在这种情况，可以使用空白标识符`_`来重命名这个导入：

`_`操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的`init`函数。

```go
import (
    _ "fmt"
)
```

### 测试案例

#### 目录结构

![QQ截图20200827175859](images\QQ截图20200827175859.png)

#### 测试代码

/src/main.go

```go
package main //必须

import (
	"calc"
	"fmt"
)

func init() {
	fmt.Println("this is main init")
}

func main() {
	a := calc.Add(10, 20)
	fmt.Println("a = ", a)

	fmt.Println("r = ", calc.Minus(10, 5))
}

```

/src/calc/calc.go

```go
package calc

import "fmt"

func init() {
	fmt.Println("this is calc init")
}

//func add(a, b int) int {
func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}
```

#### GOPATH设置

略

#### 编译运行程序

![编译运行程序](images\编译运行程序.jpg)



#### go install的使用

设置环境变量`GOBIN`：

![设置环境变量GOBIN](images\设置环境变量GOBIN.png)

在源码目录，敲`go install`:

![在源码目录敲go install](E:\smile\go\images\在源码目录敲go install.png)





## 面向对象

### 概述

对于面向对象编程的支持Go 语言设计得非常简洁而优雅。因为， Go语言并没有沿袭传统面向对象编程中的诸多概念，比如继承(不支持继承，尽管匿名字段的内存布局和行为类似继承，但它并不是继承)、虚函数、构造函数和析构函数、隐藏的this指针等。

尽管Go语言中没有封装、继承、多态这些概念，但同样通过别的方式实现这些特性：

- 封装：通过方法实现
- 继承：通过匿名字段实现
- 多态：通过接口实现

### 匿名组合

#### 匿名字段

一般情况下，定义结构体的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，也就是`匿名字段`，也称为`嵌入字段`。

当匿名字段也是一个结构体的时候，那么这个结构体所拥有的全部字段都被隐式地引入了当前定义的这个结构体。

```go
//人
type Person struct {
    name string
    sex  byte
    age  int
}
//学生
type Student struct {
    Person // 匿名字段，那么默认Student就包含了Person的所有字段
    id     int
    addr   string
}
```

#### 初始化

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
}

func main() {
	//顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println("s1 = ", s1)

	//自动推导类型
	s2 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	//fmt.Println("s2 = ", s2)
    
	//%+v, 显示更详细
	fmt.Printf("s2 = %+v\n", s2)

	//指定成员初始化，没有初始化的常用自动赋值为0
	s3 := Student{id: 1}
	fmt.Printf("s3 = %+v\n", s3)

	s4 := Student{Person: Person{name: "mike"}, id: 1}
	fmt.Printf("s4 = %+v\n", s4)

	//s5 := Student{"mike", 'm', 18, 1, "bj"} //err
}
```

#### 成员的操作

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
}

func main() {
    
	s1 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	s1.name = "yoyo"
	s1.sex = 'f'
	s1.age = 22
	s1.id = 666
	s1.addr = "sz"

	s1.Person = Person{"go", 'm', 18}

	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
}
```

#### 同名字段

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
	name   string //和Person同名了
}

func main() {
	//声明（定义一个变量）
	var s Student

	//默认规则（纠结原则），如果能在本作用域找到此成员，就操作此成员
	//					如果没有找到，找到继承的字段
	s.name = "mike" //操作的是Student的name，还是Person的name?, 结论为Student的
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"

	//显式调用
	s.Person.name = "yoyo" //Person的name

	fmt.Printf("s = %+v\n", s)
}
```

#### 其它匿名字段

##### 非结构体类型

所有的内置类型和自定义类型都是可以作为匿名字段的：

```go
package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

type Student struct {
	Person //结构体匿名字段
	int    //基础类型的匿名字段
	mystr
}

func main() {
    
	s := Student{Person{"mike", 'm', 18}, 666, "hehehe"}
	fmt.Printf("s = %+v\n", s)

	s.Person = Person{"go", 'm', 22}

	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr)
	fmt.Println(s.Person, s.int, s.mystr)
}
```



##### 结构体指针类型

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

type Student struct {
	*Person //指针类型
	id      int
	addr    string
} 

func main() {
	s1 := Student{&Person{"mike", 'm', 18}, 666, "bj"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

	//先定义变量
	var s2 Student
	s2.Person = new(Person) //分配空间
	s2.name = "yoyo"
	s2.sex = 'm'
	s2.age = 18
	s2.id = 222
	s2.addr = "sz"
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.addr)
}
```



### 方法

#### 概述

在面向对象编程中，一个对象其实也就是一个简单的值或者一个变量，在这个对象中会包含一些函数，这种带有接收者的函数，我们称为方法(method)。 本质上，一个方法则是一个和特殊类型关联的函数。

一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象，而是借助方法来做这些事情。

在Go语言中，可以给任意自定义类型（包括内置类型，但不包括指针类型）添加相应的方法。

`⽅法总是绑定对象实例`，**并隐式将实例作为第⼀实参 (receiver)**，方法的语法如下：

```go
func (receiver ReceiverType) funcName(parameters) (results)
```

* 参数 receiver 可任意命名。如⽅法中未曾使⽤，可省略参数名。
* 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接⼝或指针。
* 不支持重载方法，也就是说，不能定义名字相同但是不同参数的方法。

```go
package main

import "fmt"

//实现2数相加
//面向过程
func Add01(a, b int) int {
	return a + b
}

//面向对象，方法：给某个类型绑定一个函数
type long int

//tmp叫接收者，接收者就是传递的一个参数
func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = Add01(1, 1) //普通函数调用方式
	fmt.Println("result = ", result)

	//定义一个变量
	var a long = 2
	//调用方法格式： 变量名.函数(所需参数)
	r := a.Add02(3)
	fmt.Println("r = ", r)

	//面向对象只是换了一种表现形式
}
```



#### 为类型添加方法

##### 基础类型作为接收者

```go
type MyInt int //自定义类型，给int改名为MyInt

//在函数定义时，在其名字之前放上一个变量，即是一个方法
func (a MyInt) Add(b MyInt) MyInt { //面向对象
    return a + b
}

//传统方式的定义
func Add(a, b MyInt) MyInt { //面向过程
    return a + b
}

func main() {
    var a MyInt = 1
    var b MyInt = 1

    //调用func (a MyInt) Add(b MyInt)
    fmt.Println("a.Add(b) = ", a.Add(b)) //a.Add(b) =  2

    //调用func Add(a, b MyInt)
    fmt.Println("Add(a, b) = ", Add(a, b)) //Add(a, b) =  2
}
```

通过上面的例子可以看出，面向对象只是换了一种语法形式来表达。方法是函数的语法糖，因为receiver其实就是方法所接收的第1个参数。

注意：虽然方法的名字一模一样，但是如果接收者不一样，那么方法就不一样。

##### 结构体作为接收者

方法里面可以访问接收者的字段，调用方法通过点( . )访问，就像struct里面访问字段一样：

```go
type Person struct {
    name string
    sex  byte
    age  int
}

func (p Person) PrintInfo() { //给Person添加方法
    fmt.Println(p.name, p.sex, p.age)
}

func main() {
    p := Person{"mike", 'm', 18} //初始化
    p.PrintInfo() //调用func (p Person) PrintInfo()
}
```



```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//带有接收者的函数叫方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

//通过一个函数，给成员赋值
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

type long int

//只要接收者类型不一样，这个方法就算同名，也是不同方法，不会出现重复定义函数的错误
func (tmp long) test() {

}

type char byte

func (tmp char) test() {

}

func (tmp *long) test02() {

}

type pointer *int

//pointer为接收者类型，它本身不能是指针类型
//func (tmp pointer) test() {

//}

func main() {
    
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()

}
```

#### 值语义和引用语义

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//修改成员变量的值

//接收者为普通变量，非指针，值语义，一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p = ", p)
	fmt.Printf("SetInfoValue &p = %p\n", &p)
}

//接收者为指针变量，引用传递
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a

	fmt.Printf("SetInfoPointer p = %p\n", p)
}

func main() {
    
	s1 := Person{"go", 'm', 22}
	fmt.Printf("&s1 = %p\n", &s1) //打印地址

	//值语义
	//s1.SetInfoValue("mike", 'm', 18)
	//fmt.Println("s1 = ", s1) //打印内容

	//引用语义
	(&s1).SetInfoPointer("mike", 'm', 18)
	fmt.Println("s1 = ", s1) //打印内容
}
```

### 方法集

类型的方法集是指可以被该类型的值调用的所有方法的集合。

用实例 实例 value 和 pointer 调用方法（含匿名字段）不受⽅法集约束，编译器编总是查找全部方法，并自动转换 receiver 实参。

#### 类型 *T 方法集

一个指向自定义类型的值的指针，它的方法集由该类型定义的所有方法组成，无论这些方法接受的是一个值还是一个指针。

如果在指针上调用一个接受值的方法，Go语言会聪明地将该指针`解引用`，并将指针所指的底层值作为方法的接收者。

类型` *T` ⽅法集包含全部 `receiver T + *T `⽅法：

```go
type Person struct {
    name string
    sex  byte
    age  int
}

//指针作为接收者，引用语义
func (p *Person) SetInfoPointer() {
    (*p).name = "yoyo"
    p.sex = 'f'
    p.age = 22
}

//值作为接收者，值语义
func (p Person) SetInfoValue() {
    p.name = "xxx"
    p.sex = 'm'
    p.age = 33
}

func main() {
    //p 为指针类型
    var p *Person = &Person{"mike", 'm', 18}
    p.SetInfoPointer() //func (p) SetInfoPointer()

    p.SetInfoValue()    //func (*p) SetInfoValue()
    (*p).SetInfoValue() //func (*p) SetInfoValue()
}
```

指针变量的方法集

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
    
	//结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 18}
	//p.SetInfoPointer() //func (p *Person) SetInfoPointer()
	(*p).SetInfoPointer() //把(*p)转换层p后再调用，等价于上面

	//内部做的转换， 先把指针p， 转成*p后再调用
	//(*p).SetInfoValue()
	//p.SetInfoValue()

}
```

#### 类型 T 方法集

一个自定义类型值的方法集 则由为该类型定义的接收者类型为值类型的方法组成，但是不包含那些接收者类型为指针的方法。

但这种限制通常并不像这里所说的那样，因为如果我们只有一个值，仍然可以调用一个接收者为指针类型的方法，这可以借助于Go语言传值的地址能力实现。

```go
type Person struct {
    name string
    sex  byte
    age  int
}

//指针作为接收者，引用语义
func (p *Person) SetInfoPointer() {
    (*p).name = "yoyo"
    p.sex = 'f'
    p.age = 22
}

//值作为接收者，值语义
func (p Person) SetInfoValue() {
    p.name = "xxx"
    p.sex = 'm'
    p.age = 33
}

func main() {
    //p 为普通值类型
    var p Person = Person{"mike", 'm', 18}
    (&p).SetInfoPointer() //func (&p) SetInfoPointer()
    p.SetInfoPointer()    //func (&p) SetInfoPointer()
    
    p.SetInfoValue()      //func (p) SetInfoValue()
    (&p).SetInfoValue()   //func (*&p) SetInfoValue()
}
```

普通变量的方法集

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	p := Person{"mike", 'm', 18}
	p.SetInfoPointer() //func (p *Person) SetInfoPointer()
	//内部，先把p, 转为为&p再调用， (&p).SetInfoPointer()

	p.SetInfoValue()
}
```

###  匿名字段

#### 方法的继承

如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也能调用该方法。

```go
type Person struct {
    name string
    sex  byte
    age  int
}

//Person定义了方法
func (p *Person) PrintInfo() {
    fmt.Printf("%s,%c,%d\n", p.name, p.sex, p.age)
}

type Student struct {
    Person // 匿名字段，那么Student包含了Person的所有字段
    id     int
    addr   string
}

func main() {
    p := Person{"mike", 'm', 18}
    p.PrintInfo()

    s := Student{Person{"yoyo", 'f', 20}, 2, "sz"}
    s.PrintInfo()
}
```

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//Person类型，实现了一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承Person字段，成员和方法都继承了
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "bj"}
	s.PrintInfo()
}
```



#### 方法的重写

```go
type Person struct {
    name string
    sex  byte
    age  int
}

//Person定义了方法
func (p *Person) PrintInfo() {
    fmt.Printf("Person: %s,%c,%d\n", p.name, p.sex, p.age)
}

type Student struct {
    Person // 匿名字段，那么Student包含了Person的所有字段
    id     int
    addr   string
}

//Student定义了方法
func (s *Student) PrintInfo() {
    fmt.Printf("Student：%s,%c,%d\n", s.name, s.sex, s.age)
}

func main() {
    
    p := Person{"mike", 'm', 18}
    p.PrintInfo() //Person: mike,m,18

    s := Student{Person{"yoyo", 'f', 20}, 2, "sz"}
    s.PrintInfo()        //Student：yoyo,f,20
    s.Person.PrintInfo() //Person: mike,m,18
}
```

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//Person类型，实现了一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承Person字段，成员和方法都继承了
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

//Student也实现了一个方法，这个方法和Person方法同名，这种方法叫重写
func (tmp *Student) PrintInfo() {
	fmt.Println("Student: tmp = ", tmp)
}

func main() {
    
	s := Student{Person{"mike", 'm', 18}, 666, "bj"}
	//就近原则：先找本作用域的方法，找不到再用继承的方法
	s.PrintInfo() //到底调用的是Person， 还是Student， 结论是Student

	//显式调用继承的方法
	s.Person.PrintInfo()
}
```



### 表达式

类似于我们可以对函数进行赋值和传递一样，方法也可以进行赋值和传递。

根据调用者不同，方法分为两种表现形式：方法值和方法表达式。两者都可像普通函数那样赋值和传参，区别在于方法值绑定实例，⽽方法表达式则须显式传参。

#### 方法值

```go
type Person struct {
    name string
    sex  byte
    age  int
}

func (p *Person) PrintInfoPointer() {
    fmt.Printf("%p, %v\n", p, p)
}

func (p Person) PrintInfoValue() {
    fmt.Printf("%p, %v\n", &p, p)
}

func main() {
    
    p := Person{"mike", 'm', 18}
    p.PrintInfoPointer() //0xc0420023e0, &{mike 109 18}

    pFunc1 := p.PrintInfoPointer //方法值，隐式传递 receiver
    pFunc1()                     //0xc0420023e0, &{mike 109 18}

    pFunc2 := p.PrintInfoValue
    pFunc2() //0xc042048420, {mike 109 18}
}
```



```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, p)
}

func main() {
    
	p := Person{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	p.SetInfoPointer() //传统调用方式

	//保存方式入口地址
	pFunc := p.SetInfoPointer //这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者
	pFunc()                   //等价于 p.SetInfoPointer()

	vFunc := p.SetInfoValue
	vFunc() //等价于 p.SetInfoValue()

}
```



#### 方法表达式

```go
type Person struct {
    name string
    sex  byte
    age  int
}

func (p *Person) PrintInfoPointer() {
    fmt.Printf("%p, %v\n", p, p)
}

func (p Person) PrintInfoValue() {
    fmt.Printf("%p, %v\n", &p, p)
}

func main() {
    
    p := Person{"mike", 'm', 18}
    p.PrintInfoPointer() //0xc0420023e0, &{mike 109 18}

    //方法表达式， 须显式传参
    //func pFunc1(p *Person))
    pFunc1 := (*Person).PrintInfoPointer
    pFunc1(&p) //0xc0420023e0, &{mike 109 18}

    pFunc2 := Person.PrintInfoValue
    pFunc2(p) //0xc042002460, {mike 109 18}
}
```

```go
package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, p)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	//方法值   f := p.SetInfoPointer //隐藏了接收者
	//方法表达式
	f := (*Person).SetInfoPointer
	f(&p) //显式把接收者传递过去 ====》 p.SetInfoPointer()

	f2 := (Person).SetInfoValue
	f2(p) //显式把接收者传递过去 ====》 p.SetInfoValue()
}
```

### 接口

#### 概述

在Go语言中，接口(interface)是一个自定义类型，接口类型具体描述了一系列方法的集合。

接口类型是一种抽象的类型，它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合，它们只会展示出它们自己的方法。因此接口类型不能将其实例化。

Go通过接口实现了鸭子类型(duck-typing)：“当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子”。我们并不关心对象是什么类型，到底是不是鸭子，只关心行为。

#### 接口的使用

```go
package main

import "fmt"

//定义接口类型
type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

//Student实现了此方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

//Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", tmp.addr, tmp.group)
}

type MyStr string

//MyStr实现了此方法
func (tmp *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *tmp)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同表现，多态
func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	s := &Student{"mike", 666}
	t := &Teacher{"bj", "go"}
	var str MyStr = "hello mike"

	//调用同一函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	//第一个返回下标，第二个返回下标所对应的值
	for _, i := range x {
		i.sayhi()
	}

}

func main01() {
    
	//定义接口类型的变量
	var i Humaner

	//只是实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值
	s := &Student{"mike", 666}
	i = s
	i.sayhi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayhi()

	var str MyStr = "hello mike"
	i = &str
	i.sayhi()
}
```



##### 接口定义

```go
type Humaner interface {
    SayHi()
}
```

* 接⼝命名习惯以 er 结尾
* 接口只有方法声明，没有实现，没有数据字段
* 接口可以匿名嵌入其它接口，或嵌入到结构中

##### 接口实现

接口是用来定义行为的类型。这些被定义的行为不由接口直接实现，而是通过方法由用户定义的类型实现，一个实现了这些方法的具体类型是这个接口类型的实例。

如果用户定义的类型实现了某个接口类型声明的一组方法，那么这个用户定义的类型的值就可以赋给这个接口类型的值。这个赋值会把用户定义的类型的值存入接口类型的值。

```go
type Humaner interface {
    SayHi()
}

type Student struct { //学生
    name  string
    score float64
}

//Student实现SayHi()方法
func (s *Student) SayHi() {
    fmt.Printf("Student[%s, %f] say hi!!\n", s.name, s.score)
}

type Teacher struct { //老师
    name  string
    group string
}

//Teacher实现SayHi()方法
func (t *Teacher) SayHi() {
    fmt.Printf("Teacher[%s, %s] say hi!!\n", t.name, t.group)
}

type MyStr string

//MyStr实现SayHi()方法
func (str MyStr) SayHi() {
    fmt.Printf("MyStr[%s] say hi!!\n", str)
}

//普通函数，参数为Humaner类型的变量i
func WhoSayHi(i Humaner) {
    i.SayHi()
}

func main() {
    s := &Student{"mike", 88.88}
    t := &Teacher{"yoyo", "Go语言"}
    var tmp MyStr = "测试"

    s.SayHi()   //Student[mike, 88.880000] say hi!!
    t.SayHi()   //Teacher[yoyo, Go语言] say hi!!
    tmp.SayHi() //MyStr[测试] say hi!!

    //多态，调用同一接口，不同表现
    WhoSayHi(s)   //Student[mike, 88.880000] say hi!!
    WhoSayHi(t)   //Teacher[yoyo, Go语言] say hi!!
    WhoSayHi(tmp) //MyStr[测试] say hi!!

    x := make([]Humaner, 3)
    //这三个都是不同类型的元素，但是他们实现了interface同一个接口
    x[0], x[1], x[2] = s, t, tmp
    for _, value := range x {
        value.SayHi()
    }
    /*
        Student[mike, 88.880000] say hi!!
        Teacher[yoyo, Go语言] say hi!!
        MyStr[测试] say hi!!
    */
}
```

通过上面的代码，你会发现接口就是一组抽象方法的集合，它必须由其他非接口类型实现，而不能自我实现。





#### 接口组合

```go
package main

import "fmt"

type Humaner interface { //子集
	sayhi()
}

type Personer interface { //超集
	Humaner //匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

//Student实现了sayhi()
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student在唱着：", lrc)
}

func main() {
	//定义一个接口类型的变量
	var i Personer
	s := &Student{"mike", 666}
	i = s

	i.sayhi() //继承过来的方法
	i.sing("学生哥")
}
```



##### 接口嵌入

如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的方法。

```go
type Humaner interface {
    SayHi()
}

type Personer interface {
    Humaner //这里想写了SayHi()一样
    Sing(lyrics string)
}

type Student struct { //学生
    name  string
    score float64
}

//Student实现SayHi()方法
func (s *Student) SayHi() {
    fmt.Printf("Student[%s, %f] say hi!!\n", s.name, s.score)
}

//Student实现Sing()方法
func (s *Student) Sing(lyrics string) {
    fmt.Printf("Student sing[%s]!!\n", lyrics)
}

func main() {
    s := &Student{"mike", 88.88}

    var i2 Personer
    i2 = s
    i2.SayHi()     //Student[mike, 88.880000] say hi!!
    i2.Sing("学生哥") //Student sing[学生哥]!!
}
```



##### 接口转换

超集接⼝对象可转换为⼦集接⼝，反之出错：

```go
type Humaner interface {
    SayHi()
}

type Personer interface {
    Humaner //这里像写了SayHi()一样
    Sing(lyrics string)
}

type Student struct { //学生
    name  string
    score float64
}

//Student实现SayHi()方法
func (s *Student) SayHi() {
    fmt.Printf("Student[%s, %f] say hi!!\n", s.name, s.score)
}

//Student实现Sing()方法
func (s *Student) Sing(lyrics string) {
    fmt.Printf("Student sing[%s]!!\n", lyrics)
}

func main() {
    
    //var i1 Humaner = &Student{"mike", 88.88}
    //var i2 Personer = i1 //err

    //Personer为超集，Humaner为子集
    var i1 Personer = &Student{"mike", 88.88}
    var i2 Humaner = i1
    i2.SayHi() //Student[mike, 88.880000] say hi!!
}
```

```go
package main

import "fmt"

type Humaner interface { //子集
	sayhi()
}

type Personer interface { //超集
	Humaner //匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

//Student实现了sayhi()
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student在唱着：", lrc)
}

func main() {
    
	//超集可以转换为子集，反过来不可以
	var iPro Personer //超集
	iPro = &Student{"mike", 666}

	var i Humaner //子集

	//iPro = i //err
	i = iPro //可以，超集可以转换为子集
	i.sayhi()
}

```



#### 空接口

空接口(interface{})不包含任何的方法，正因为如此，`所有的类型都实现了空接口`，因此空接口可以存储任意类型的数值。它有点类似于C语言的void *类型。

```go
var v1 interface{} = 1     // 将int类型赋值给interface{}
var v2 interface{} = "abc" // 将string类型赋值给interface{}
var v3 interface{} = &v2   // 将*interface{}类型赋值给interface{}
var v4 interface{} = struct{ X int }{1}
var v5 interface{} = &struct{ X int }{1}
```

当函数可以接受任意的对象实例时，我们会将其声明为interface{}，最典型的例子是标准库fmt中PrintXXX系列的函数，例如：

```go
func Printf(fmt string, args ...interface{})
func Println(args ...interface{})
```

```go
package main

import "fmt"

func xxx(arg ...interface{}) {

}

func main() {
	//空接口万能类型，保存任意类型的值
	var i interface{} = 1
	fmt.Println("i = ", i)

	i = "abc"
	fmt.Println("i = ", i)
}
```



#### 类型查询

我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：

* comma-ok断言
* switch测试

##### comma-ok断言

Go语言里面有一个语法，可以直接判断是否是该类型的变量：` value, ok = element.(T)`，这里value就是变量的值，`ok`是一个bool类型，element是interface变量，T是断言的类型。

如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。

示例代码：

```go
type Element interface{}

type Person struct {
    name string
    age  int
}

func main() {
    
    list := make([]Element, 3)
    list[0] = 1       // an int
    list[1] = "Hello" // a string
    list[2] = Person{"mike", 18}

    for index, element := range list {
        if value, ok := element.(int); ok {
            fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
        } else if value, ok := element.(string); ok {
            fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
        } else if value, ok := element.(Person); ok {
            fmt.Printf("list[%d] is a Person and its value is [%s, %d]\n", index, value.name, value.age)
        } else {
            fmt.Printf("list[%d] is of a different type\n", index)
        }
    }

    /*  打印结果：
    list[0] is an int and its value is 1
    list[1] is a string and its value is Hello
    list[2] is a Person and its value is [mike, 18]
    */
}
```

```go
package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1                    //int
	i[1] = "hello go"           //string
	i[2] = Student{"mike", 666} //Student

	//类型查询，类型断言
	//第一个返回下标，第二个返回下标对应的值， data分别是i[0], i[1], i[2]
	for index, data := range i {
        
		//第一个返回的是值，第二个返回判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student, 内容为name = %s, id = %d\n", index, value.name, value.id)
		}
	}

}
```



##### switch测试

```go
type Element interface{}

type Person struct {
    name string
    age  int
}

func main() {
    
    list := make([]Element, 3)
    list[0] = 1       //an int
    list[1] = "Hello" //a string
    list[2] = Person{"mike", 18}

    for index, element := range list {
        switch value := element.(type) {
        case int:
            fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
        case string:
            fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
        case Person:
            fmt.Printf("list[%d] is a Person and its value is [%s, %d]\n", index, value.name, value.age)
        default:
            fmt.Println("list[%d] is of a different type", index)
        }
    }
}
```



```go
package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
    
	i := make([]interface{}, 3)
	i[0] = 1                    //int
	i[1] = "hello go"           //string
	i[2] = Student{"mike", 666} //Student

	//类型查询，类型断言
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d] 类型为Student, 内容为name = %s, id = %d\n", index, value.name, value.id)
		}

	}
}
```



### 封装





### 继承





### 多态





## 错误处理

### error接口

Go语言引入了一个关于错误处理的标准模式，即`error`接口，它是Go语言内建的接口类型，该接口的定义如下：

```go
type error interface {
    Error() string
}
```

Go语言的标准库代码包errors为用户提供如下方法：

```go
package errors

type errorString struct { 
    text string 
}

func New(text string) error { 
    return &errorString{text} 
}

func (e *errorString) Error() string { 
    return e.text 
}
```

另一个可以生成error类型值的方法是调用fmt包中的Errorf函数：

```go
package fmt
import "errors"

func Errorf(format string, args ...interface{}) error {
    return errors.New(Sprintf(format, args...))
}
```

示例代码：

```go
package main

import "fmt"
import "errors"

func main() {
	//var err1 error = fmt.Errorf("%s", "this is normol err")
	err1 := fmt.Errorf("%s", "this is normal err1")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is normal err2")
	fmt.Println("err2 = ", err2)
}

```

函数通常在最后的返回值中返回错误信息：

```go
package main

import "fmt"
import "errors"

func MyDiv(a, b int) (result int, err error) {

	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}

	return
}

func main() {
	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("reslut = ", result)
	}

}
```

### 自定义错误

```go
package main

import (
    "fmt"
)

// 自定义错误信息结构
type DIV_ERR struct {   
    etype int  // 错误类型   
    v1 int     // 记录下出错时的除数、被除数   
    v2 int
}

// 实现接口方法 error.Error()
func (div_err DIV_ERR) Error() string {   
    if 0==div_err.etype {      
        return "除零错误"   
    }else{   
        return "其他未知错误"  
    }
}

// 除法
func div(a int, b int) (int,*DIV_ERR) {  
    if b == 0 {     
        // 返回错误信息    
        return 0, &DIV_ERR{0,a,b}  
    } else {   
        // 返回正确的商 
        return a / b, nil   
    }
}

func main() { 
    // 正确调用  
    v,r := div(100,2) 
    if nil!=r{   
        fmt.Println("(1)fail:",r)  
    }else{   
        fmt.Println("(1)succeed:",v) 
    }   
    // 错误调用
    v,r = div(100,0) 
    if nil!=r{   
        fmt.Println("(2)fail:",r)  
    }else{  
        fmt.Println("(2)succeed:",v) 
    }
}
```



### panic

panic 与 recover,一个用于主动抛出错误，一个用于捕获panic抛出的错误。

在通常情况下，向程序使用方报告错误状态的方式可以是返回一个额外的`error`类型值。

但是，当遇到不可恢复的错误状态的时候，如数组访问越界、空指针引用等，这些运行时错误会引起`painc异常`。这时，上述错误处理方式显然就不适合了。反过来讲，在一般情况下，我们不应通过调用panic函数来报告普通的错误，而应该只把它作为报告致命错误的一种方式。当某些不应该发生的场景发生时，我们就应该调用panic。

一般而言，`当panic异常发生时，程序会中断运行`，并立即执行在该`goroutine`（可以先理解成线程，在中被延迟的函数（defer 机制）。随后，程序崩溃并输出日志信息。日志信息包括panic value和函数调用的堆栈跟踪信息。

不是所有的panic异常都来自运行时，直接调用内置的panic函数也会引发panic异常；panic函数接受任何值作为参数。

```go
func panic(v interface{})
```

调用panic函数引发的panic异常：

```go
package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}

func testb() {
	//fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	//显式调用panic函数，导致程序中断
	panic("this is a panic test")
}

func testc() {
	fmt.Println("cccccccccccccccccc")
}

func main() {
	testa()
	testb()
	testc()
}
```

内置的panic函数引发的panic异常：

```go
package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}

func testb(x int) {
	var a [10]int
	a[x] = 111 //当x为20时候，导致数组越界，产生一个panic，导致程序崩溃
}

func testc() {
	fmt.Println("cccccccccccccccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
```

- 引发`panic`有两种情况，一是程序主动调用，二是程序产生运行时错误，由运行时检测并退出。
- 发生`panic`后，程序会从调用`panic`的函数位置或发生`panic`的地方立即返回，逐层向上执行函数的`defer`语句，然后逐层打印函数调用堆栈，直到被`recover`捕获或运行到最外层函数。
- `panic`不但可以在函数正常流程中抛出，在`defer`逻辑里也可以再次调用`panic`或抛出`panic`。`defer`里面的`panic`能够被后续执行的`defer`捕获。
- `recover`用来捕获`panic`，阻止`panic`继续向上传递。`recover()`和`defer`一起使用，但是`defer`只有在后面的函数体内直接被掉用才能捕获`panic`来终止异常，否则返回`nil`，异常继续向外传递。

```go
//以下捕获失败
defer recover()
defer fmt.Prinntln(recover)
defer func(){
    func(){
        recover() //无效，嵌套两层
    }()
}()

//以下捕获有效
defer func(){
    recover()
}()

func except(){
    recover()
}
func test(){
    defer except()
    panic("runtime error")
}
```



### recover

运行时panic异常一旦被引发就会导致程序崩溃。这当然不是我们愿意看到的，因为谁也不能保证程序不会发生任何运行时错误。

不过，Go语言为我们提供了专用于“拦截”运行时`panic的内建函数——recover`。它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。

```go
func recover() interface{}
```

>  注意：`recover`只有在`defer`调用的函数中有效。

如果调用了内置函数`recover`，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回`nil`。

示例代码：

```go
package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}

func testb(x int) {
    
	//设置recover
	defer func() {
		//recover() //可以打印panic的错误信息
		//fmt.Println(recover())
		if err := recover(); err != nil { //产生了panic异常
			fmt.Println(err)
		}

	}() //别忘了(), 调用此匿名函数

	var a [10]int
	a[x] = 111 //当x为20时候，导致数组越界，产生一个panic，导致程序崩溃
}

func testc() {
	fmt.Println("cccccccccccccccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
```

延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后⼀个错误可被捕获：

```go
func test() {
    defer func() {
        fmt.Println(recover())
    }()

    defer func() {
        panic("defer panic")
    }()

    panic("test panic")
}

func main() {
    test()
    //运行结果：defer panic
}
```



## 编译和工具链

### 工具

Go语言的工具箱集合了一系列的功能的命令集。它可以看作是一个包管理器（类似于Linux中的apt和rpm工具），用于包的查询、计算的包依赖关系、从远程版本控制系统和下载它们等任务。它也是一个构建系统，计算文件的依赖关系，然后调用编译器、汇编器和连接器构建程序，虽然它故意被设计成没有标准的make命令那么复杂。它也是一个单元测试和基准测试的驱动程序，我们将在第11章讨论测试话题。
Go语言工具箱的命令有着类似“瑞士军刀”的风格，带着一打子的子命令，有一些我们经常用到，例如get、run、build和fmt等。你可以运行go或go help命令查看内置的帮助文档，为了查询方便，我们列出了最常用的命令：

```go
$ go
...
    build            compile packages and dependencies
    clean            remove object files
    doc              show documentation for package or symbol
    env              print Go environment information
    fmt              run gofmt on package sources
    get              download and install packages and dependencies
    install          compile and install packages and dependencies
    list             list packages
    run              compile and run Go program
    test             test packages
    version          print Go version
    vet              run go tool vet on packages

Use "go help [command]" for more information about a command.
...
```

为了达到零配置的设计目标，Go语言的工具箱很多地方都依赖各种约定。例如，根据给定的源文件的名称，Go语言的工具可以找到源文件对应的包，因为每个目录只包含了单一的包，并且到的导入路径和工作区的目录结构是对应的。给定一个包的导入路径，Go语言的工具可以找到对应的目录中没个实体对应的源文件。它还可以根据导入路径找到存储代码仓库的远程服务器的URL。

###  go build命令

Go语言的编译速度非常快。Go 1.9 版本后默认利用Go语言的并发特性进行函数粒度的并发编译。

Go语言的程序编写基本以源码方式，无论是自己的代码还是第三方代码，并且以 GOPATH 作为工作目录和一套完整的工程目录规则

Go语言中使用` go build` 命令主要用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。

go build 有很多种编译方法，如无参数编译、文件列表编译、指定包编译等，使用这些方法都可以输出可执行文件。

#### go build 无参数编译

​	代码相对于 GOPATH 的目录关系如下：

```go
.
└── src
    └── chapter11
        └── gobuild
            ├── lib.go
            └── main.go
```

main.go 代码如下：

```
package main
import (
    "fmt"
)
func main() {
    // 同包的函数
    pkgFunc()
    fmt.Println("hello world")
}
```

lib.go 代码如下：

```go
package main
import "fmt"
func pkgFunc() {
    fmt.Println("call pkgFunc")
}
```

如果源码中没有依赖` GOPATH` 的包引用，那么这些源码可以使用无参数 go build。格式如下：

```go
go build
```

在代码所在目录（`./src/chapter11/gobuild`）下使用 go build 命令，如下所示：

```go
$ cd src/chapter11/gobuild/  #转到本例源码目录下。
$ go build #go build 在编译开始时，会搜索当前目录的 go 源码。
$ ls
gobuild lib.go main.go
$ ./gobuild   #，运行当前目录的可执行文件 go build。
call pkgFunc
hello world
```

#### go build+文件列表

编译同目录的多个源码文件时，可以在 go build 的后面提供多个文件名，go build 会编译这些源码，输出可执行文件，“go build+文件列表”的格式如下：

```go
go build file1.go file2.go……
```

在代码代码所在目录（./src/chapter11/gobuild）中使用 go build，在 go build 后添加要编译的源码文件名，代码如下：

```go
$ go build main.go lib.go //go build 后添加文件列表，选中需要编译的 Go 源码。
$ ls
lib.go  main  main.go  //列出完成编译后的当前目录的文件。这次的可执行文件名变成了 main。
$ ./main //执行 main 文件，得到期望输出。
call pkgFunc
hello world

$ go build lib.go main.go //，尝试调整文件列表的顺序，将 lib.go 放在列表的首位。
$ ls
lib  lib.go  main  main.go //，编译结果中出现了 lib 可执行文件。
```

>  提示

使用“`go build+文件列表`”方式编译时，可执行文件默认选择文件列表中第一个源码文件作为可执行文件名输出。

如果需要指定输出可执行文件名，可以使用`-o`参数，参见下面的例子：

```go
$ go build -o myexec main.go lib.go
$ ls
lib.go  main.go  myexec
$ ./myexec
call pkgFunc
hello world
```

上面代码中，在 go build 和文件列表之间插入了`-o myexec`参数，表示指定输出文件名为 myexec。

> 注意

使用“go build+文件列表”编译方式编译时，文件列表中的每个文件必须是同一个包的 Go 源码。也就是说，不能像 C++ 一样将所有工程的 Go 源码使用文件列表方式进行编译。编译复杂工程时需要用“指定包编译”的方式。

“go build+文件列表”方式更适合使用Go语言编写的只有少量文件的工具。

#### go build+包

“go build+包”在设置 GOPATH 后，可以直接根据包名进行编译，即便包内文件被增（加）删（除）也不影响编译指令。

1. 代码位置及源码

本小节需要用到的代码具体位置是`./src/chapter11/goinstall`。

相对于GOPATH的目录关系如下：

```go
.
└── src
    └── chapter11
        └──goinstall
            ├── main.go
            └── mypkg
                └── mypkg.go
```

main.go代码如下：

```go
package main
import (
    "chapter11/goinstall/mypkg"
    "fmt"
)
func main() {
    mypkg.CustomPkgFunc()
    fmt.Println("hello world")
}
```

mypkg.go代码如下：

```go
package mypkg
import "fmt"
func CustomPkgFunc() {
    fmt.Println("call CustomPkgFunc")
}
```

2. 按包编译命令

执行以下命令将按包方式编译 goinstall 代码：

```go
$ export GOPATH=/home/davy/golangbook/code //设置环境变量 GOPATH
$ go build -o main chapter11/goinstall
$ ./goinstall
call CustomPkgFunc
hello world
```

> `-o`执行指定输出文件为 main，后面接要编译的包名。包名是相对于 GOPATH 下的 src 目录开始的。

#### go build 编译时的附加参数

go build 还有一些附加参数，可以显示更多的编译信息和更多的操作，详见下表所示。



| 附加参数 | 备  注                                      |
| -------- | ------------------------------------------- |
| -v       | 编译时显示包名                              |
| -p n     | 开启并发编译，默认情况下该值为 CPU 逻辑核数 |
| -a       | 强制重新构建                                |
| -n       | 打印编译时会用到的所有命令，但不真正执行    |
| -x       | 打印编译时会用到的所有命令                  |
| -race    | 开启竞态检测                                |


表中的附加参数按使用频率排列，读者可以根据需要选择使用。

### go clean命令

Go语言中`go clean`命令可以移除当前源码包和关联源码包里面编译生成的文件，这些文件包括以下几种：

- 执行`go build`命令时在当前目录下生成的与包名或者 Go 源码文件同名的可执行文件。在 Windows 下，则是与包名或者 Go 源码文件同名且带有“.exe”后缀的文件。
- 执行`go test`命令并加入`-c`标记时在当前目录下生成的以包名加“.test”后缀为名的文件。在 Windows 下，则是以包名加“.test.exe”后缀的文件。
- 执行`go install`命令安装当前代码包时产生的结果文件。如果当前代码包中只包含库源码文件，则结果文件指的就是在工作区 pkg 目录下相应的归档文件。如果当前代码包中只包含一个命令源码文件，则结果文件指的就是在工作区 bin 目录下的可执行文件。
- 在编译 Go 或 C 源码文件时遗留在相应目录中的文件或目录 。包括：“_obj”和“_test”目录，名称为“_testmain.go”、“test.out”、“build.out”或“a.out”的文件，名称以“.5”、“.6”、“.8”、“.a”、“.o”或“.so”为后缀的文件。这些目录和文件是在执行`go build`命令时生成在临时目录中的。

`go clean`命令就像 [Java](http://c.biancheng.net/java/) 中的`maven clean`命令一样，会清除掉编译过程中产生的一些文件。在 Java 中通常是 .class 文件，而在Go语言中通常是上面我们所列举的那些文件。

```go
go clean -i -n
```

通过上面的示例可以看出，`go clean`命令还可以指定一些参数。对应的参数的含义如下所示：

- -i 清除关联的安装的包和可运行文件，也就是通过`go install`安装的文件；
- -n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的；
- -r 循环的清除在 import 中引入的包；
- -x 打印出来执行的详细命令，其实就是 -n 打印
- 的执行版本；
- -cache 删除所有`go build`命令的缓存
- -testcache 删除当前包所有的测试结果


实际开发中`go clean`命令使用的可能不是很多，一般都是利用`go clean`命令清除编译文件，然后再将源码递交到 github 上，方便对于源码的管理。

下面我们以本地的一个项目为例，演示一下`go clean`命令：

```go
go clean -n
cd D:\code
rm -f code code.exe code.test code.test.exe main main.exe
```

在命令中使用`-n`标记可以将命令的执行过程打印出来，但不会正真执行。如果既要打印命令的执行过程同时又执行命令的话可以使用`-x`标记，如下所示：

```
go clean -x
cd D:\code
rm -f code code.exe code.test code.test.exe main main.exe
```



### go run命令

Python或者 Lua 语言可以在不输出二进制的情况下，将代码使用虚拟机直接执行。Go语言虽然不使用虚拟机，但可使用`go run`指令达到同样的效果。

`go run`命令会编译源码，并且直接执行源码的 main() 函数，不会在当前目录留下可执行文件。

下面我们准备一个 main.go 的文件来观察`go run`的运行结果，源码如下：

```go
package main
import (   
    "fmt"
    "os"
)
func main() {
    fmt.Println("args:", os.Args)
}
```

这段代码的功能是将输入的参数打印出来。使用`go run`运行这个源码文件，命令如下：

```go
$ go run main.go --filename xxx.go
args: [/tmp/go-build006874658/command-line-arguments/_obj/exe/main--filename xxx.go]
```

`go run`不会在运行目录下生成任何文件，可执行文件被放在临时文件中被执行，工作目录被设置为当前目录。在`go run`的后部可以添加参数，这部分参数会作为代码可以接受的命令行输入提供给程序。

`go run`不能使用“go run+包”的方式进行编译，如需快速编译运行包，需要使用如下步骤来代替：

1. 使用`go build`生成可执行文件。
2. 运行可执行文件。



### go fmt命令

#### gofmt 介绍

Go语言的开发团队制定了统一的官方代码风格，并且推出了 gofmt 工具（gofmt 或 go fmt）来帮助开发者格式化他们的代码到统一的风格。

gofmt 是一个 cli 程序，会优先读取标准输入，如果传入了文件路径的话，会格式化这个文件，如果传入一个目录，会格式化目录中所有 .go 文件，如果不传参数，会格式化当前目录下的所有 .go 文件。

gofmt 默认不对代码进行简化，使用`-s`参数可以开启简化代码功能，具体来说会进行如下的转换：

##### 1) 去除数组、切片、Map 初始化时不必要的类型声明

如下形式的切片表达式：

```
[]T{T{}, T{}}
```

简化后的代码为：

```
[]T{{}, {}}
```



##### 2) 去除数组切片操作时不必要的索引指定

如下形式的切片表达式：

```
s[a:len(s)]
```

简化后的代码为：

```
s[a:]
```



##### 3) 去除循环时非必要的变量赋值

如下形式的循环：

```go
for x, _ = range v {...}
```

简化后的代码为：

```go
for x = range v {...}
```

如下形式的循环：

```go
for _ = range v {...}
```

简化后的代码为：

```go
for range v {...}
```

gofmt 命令参数如下表所示：



| 标记名称    | 标记描述                                                     |
| ----------- | ------------------------------------------------------------ |
| -l          | 仅把那些不符合格式化规范的、需要被命令程序改写的源码文件的绝对路径打印到标准输出。而不是把改写后的全部内容都打印到标准输出。 |
| -w          | 把改写后的内容直接写入到文件中，而不是作为结果打印到标准输出。 |
| -r          | 添加形如“a[b:len(a)] -> a[b:]”的重写规则。如果我们需要自定义某些额外的格式化规则，就需要用到它。 |
| -s          | 简化文件中的代码。                                           |
| -d          | 只把改写前后内容的对比信息作为结果打印到标准输出。而不是把改写后的全部内容都打印到标准输出。 命令程序将使用 diff 命令对内容进行比对。在 Windows 操作系统下可能没有 diff 命令，需要另行安装。 |
| -e          | 打印所有的语法错误到标准输出。如果不使用此标记，则只会打印每行的第 1 个错误且只打印前 10 个错误。 |
| -comments   | 是否保留源码文件中的注释。在默认情况下，此标记会被隐式的使用，并且值为 true。 |
| -tabwidth   | 此标记用于设置代码中缩进所使用的空格数量，默认值为 8。要使此标记生效，需要使用“-tabs”标记并把值设置为 false。 |
| -tabs       | 是否使用 tab（'\t'）来代替空格表示缩进。在默认情况下，此标记会被隐式的使用，并且值为 true。 |
| -cpuprofile | 是否开启 CPU 使用情况记录，并将记录内容保存在此标记值所指的文件中。 |


可以看到 gofmt 命令还支持自定义的重写规则，使用`-r`参数，按照 pattern -> replacement 的格式传入规则。

【示例】有如下内容的 Golang 程序，存储在 main.go 文件中。

```go
package main
import "fmt"
func main() {
    a := 1
    b := 2
    c := a + b
    fmt.Println(c)
}
```

用以下规则来格式化上面的代码。

```go
gofmt -w -r "a + b -> b + a" main.go
```

格式化的结果如下。

```
package main
import "fmt"
func main() {
    a := 1
    b := 2
    c := b + a
    fmt.Println(c)
}
```

> 注意：gofmt 使用 tab 来表示缩进，并且对行宽度无限制，如果手动对代码进行了换行，gofmt 不会强制把代码格式化回一行。

#### go fmt 和 gofmt

gofmt 是一个独立的 cli 程序，而Go语言中还有一个`go fmt`命令，`go fmt`命令是 gofmt 的简单封装。

```go
go help fmt
usage: go fmt [-n] [-x] [packages]

Fmt runs the command 'gofmt -l -w' on the packages named
by the import paths. It prints the names of the files that are modified.

For more about gofmt, see 'go doc cmd/gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

To run gofmt with specific options, run gofmt itself.

See also: go fix, go vet.
```

`go fmt`命令本身只有两个可选参数`-n`和`-x`：

- `-n`仅打印出内部要执行的`go fmt`的命令；
- `-x`命令既打印出`go fmt`命令又执行它，如果需要更细化的配置，需要直接执行 gofmt 命令。

`go fmt`在调用 gofmt 时添加了`-l -w`参数，相当于执行了`gofmt -l -w`。



###  go install命令

go install 命令的功能和前面一节《[go build命令](http://c.biancheng.net/view/120.html)》中介绍的 go build 命令类似，附加参数绝大多数都可以与 go build 通用。go install 只是将编译的中间文件放在 GOPATH 的 pkg 目录下，以及固定地将编译结果放在 GOPATH 的 bin 目录下。

这个命令在内部实际上分成了两步操作：第一步是生成结果文件（可执行文件或者 .a 包），第二步会把编译好的结果移到 $GOPATH/pkg 或者 $GOPATH/bin。

使用 go install 来执行代码，参考下面的 shell：

```shell
$ export GOPATH=/home/davy/golangbook/code
$ go install chapter11/goinstall
```

编译完成后的目录结构如下：

```go
├── bin
│   └── goinstall
├── pkg
│   └── linux_amd64
│       └── chapter11
│           └── goinstall
│               └── mypkg.a
└── src
    └── chapter11
        ├── gobuild
        │   ├── lib.go
        │   └── main.go
        └── goinstall
            ├── main.go
            └── mypkg
                └── mypkg.go	
```

go install 的编译过程有如下规律：

- go install 是建立在 GOPATH 上的，无法在独立的目录里使用 go install。
- GOPATH 下的 bin 目录放置的是使用 go install 生成的可执行文件，可执行文件的名称来自于编译时的包名。
- go install 输出目录始终为 GOPATH 下的 bin 目录，无法使用`-o`附加参数进行自定义。
- GOPATH 下的 pkg 目录放置的是编译期间的中间文件。



### go get命令

`go get `命令可以借助代码管理工具通过远程拉取或更新代码包及其依赖包，并自动完成编译和安装。整个过程就像安装一个 App 一样简单。

这个命令可以动态获取远程代码包，目前支持的有 BitBucket、GitHub、Google Code 和 Launchpad。在使用 go get 命令前，需要安装与远程包匹配的代码管理工具，如 Git、SVN、HG 等，参数中需要提供一个包名。

这个命令在内部实际上分成了两步操作：第一步是下载源码包，第二步是执行 go install。下载源码包的 go 工具会自动根据不同的域名调用不同的源码工具，对应关系如下：

```go
BitBucket (Mercurial Git)
GitHub (Git)
Google Code Project Hosting (Git, Mercurial, Subversion)
Launchpad (Bazaar)
```

所以为了 go get 命令能正常工作，你必须确保安装了合适的源码管理工具，并同时把这些命令加入你的 PATH 中。其实 go get 支持自定义域名的功能。

参数介绍：

- -d 只下载不安装
- -f 只有在你包含了 -u 参数的时候才有效，不让 -u 去验证 import 中的每一个都已经获取了，这对于本地 fork 的包特别有用
- -fix 在获取源码之后先运行 fix，然后再去做其他的事情
- -t 同时也下载需要为运行测试所需要的包
- -u 强制使用网络去更新包和它的依赖包
- -v 显示执行的命令

#### go get+ 远程包

默认情况下，go get 可以直接使用。例如，想获取 go 的源码并编译，使用下面的命令行即可：

```go
$ go get github.com/davyxu/cellnet
```

获取前，请确保 GOPATH 已经设置。Go 1.8 版本之后，GOPATH 默认在用户目录的 go 文件夹下。

cellnet 只是一个网络库，并没有可执行文件，因此在 go get 操作成功后 GOPATH 下的 bin 目录下不会有任何编译好的二进制文件。

需要测试获取并编译二进制的，可以尝试下面的这个命令。当获取完成后，就会自动在 GOPATH 的 bin 目录下生成编译好的二进制文件。

```go
$ go get github.com/davyxu/tabtoy
```



#### go get 使用时的附加参数

使用 go get 时可以配合附加参数显示更多的信息及实现特殊的下载和安装操作，详见下表所示。



| 附加参数  | 备  注                                 |
| --------- | -------------------------------------- |
| -v        | 显示操作流程的日志及信息，方便检查错误 |
| -u        | 下载丢失的包，但不会更新已经存在的包   |
| -d        | 只下载，不安装                         |
| -insecure | 允许使用不安全的 HTTP 方式进行下载操作 |

### go generate命令

`go generate`命令是在Go语言 1.4 版本里面新添加的一个命令，当运行该命令时，它将扫描与当前包相关的源代码文件，找出所有包含`//go:generate`的特殊注释，提取并执行该特殊注释后面的命令。

使用`go generate`命令时有以下几点需要注意：

- 该特殊注释必须在 .go 源码文件中；
- 每个源码文件可以包含多个 generate 特殊注释；
- 运行`go generate`命令时，才会执行特殊注释后面的命令；
- 当`go generate`命令执行出错时，将终止程序的运行；
- 特殊注释必须以`//go:generate`开头，双斜线后面没有空格。


在下面这些场景下，我们会使用`go generate`命令：

- yacc：从 .y 文件生成 .go 文件；
- protobufs：从 protocol buffer 定义文件（.proto）生成 .pb.go 文件；
- Unicode：从 UnicodeData.txt 生成 Unicode 表；
- HTML：将 HTML 文件嵌入到 go 源码；
- bindata：将形如 JPEG 这样的文件转成 go 代码中的字节数组。


再比如：

- string 方法：为类似枚举常量这样的类型生成 String() 方法；
- 宏：为既定的泛型包生成特定的实现，比如用于 ints 的 sort.Ints。


`go generate`命令格式如下所示：

```go
go generate [-run regexp] [-n] [-v] [-x] [command] [build flags] [file.go... | packages]
```

参数说明如下：

- -run 正则表达式匹配命令行，仅执行匹配的命令；
- -v 输出被处理的包名和源文件名；
- -n 显示不执行命令；
- -x 显示并执行命令；
- command 可以是在环境变量 PATH 中的任何命令。


执行`go generate`命令时，也可以使用一些环境变量，如下所示:

- $GOARCH 体系架构（arm、amd64 等）；
- $GOOS 当前的 OS 环境（linux、windows 等）；
- $GOFILE 当前处理中的文件名；
- $GOLINE 当前命令在文件中的行号；
- $GOPACKAGE 当前处理文件的包名；
- $DOLLAR 固定的`$`，不清楚具体用途。


【示例 1】假设我们有一个 main.go 文件，内容如下：

```go
package main
import "fmt"
//go:generate go run main.go
//go:generate go version
func main() {
    fmt.Println("http://c.biancheng.net/golang/")
}
```

执行`go generate -x`命令，输出结果如下：

```go
go generate -x
go run main.go
http://c.biancheng.net/golang/
go version
go version go1.13.6 windows/amd64
```

通过运行结果可以看出`//go:generate`之后的命令成功运行了，命令中使用的`-x`参数是为了将执行的具体命令同时打印出来。

下面通过 stringer 工具来演示一下`go generate`命令的使用。

stringer 并不是Go语言自带的工具，需要手动安装。我们可以通过下面的命令来安装 stringer 工具。

```go
go get golang.org/x/tools/cmd/stringer
```

上面的命令需要翻墙。条件不允许的话也可以通过 Github 上的镜像来安装，安装方法如下：

```go
git clone https://github.com/golang/tools/ $GOPATH/src/golang.org/x/tools
go install golang.org/x/tools/cmd/stringer
```

安装好的 stringer 工具位于 GOPATH/bin 目录下，想要正常使用它，需要先将 GOPATH/bin 目录添加到系统的环境变量 PATH 中。

【示例 2】使用 stringer 工具实现 String() 方法：

首先，在项目目录下新建一个 painkiller 文件夹，并在该文件夹中创建 painkiller.go 文件，文件内容如下：

```go
//go:generate stringer -type=Pill
package painkiller
type Pill int
const (
    Placebo Pill = iota
    Aspirin
    Ibuprofen
    Paracetamol
    Acetaminophen = Paracetamol
)
```

然后，在 painkiller.go 文件所在的目录下运行`go generate`命令。

执行成功后没有任何提示信息，但会在当前目录下面生成一个 pill_string.go 文件，文件中实现了我们需要的 String() 方法，文件内容如下：

```go
// Code generated by "stringer -type=Pill"; DO NOT EDIT.
package painkiller
import "strconv"
func _() {
    // An "invalid array index" compiler error signifies that the constant values have changed.
    // Re-run the stringer command to generate them again.
    var x [1]struct{}
    _ = x[Placebo-0]
    _ = x[Aspirin-1]
    _ = x[Ibuprofen-2]
    _ = x[Paracetamol-3]
}
const _Pill_name = "PlaceboAspirinIbuprofenParacetamol"
var _Pill_index = [...]uint8{0, 7, 14, 23, 34}
func (i Pill) String() string {
    if i < 0 || i >= Pill(len(_Pill_index)-1) {
        return "Pill(" + strconv.FormatInt(int64(i), 10) + ")"
    }
    return _Pill_name[_Pill_index[i]:_Pill_index[i+1]]
}
```



### go test命令

Go语言拥有一套单元测试和性能测试系统，仅需要添加很少的代码就可以快速测试一段需求代码。

go test 命令，会自动读取源码目录下面名为 *_test.go 的文件，生成并运行测试用的可执行文件。输出的信息类似下面所示的样子：

```go
ok archive/tar 0.011s
FAIL archive/zip 0.022s
ok compress/gzip 0.033s
...
```

性能测试系统可以给出代码的性能数据，帮助测试者分析性能问题。

**提示**

单元测试（unit testing），是指对软件中的最小可测试单元进行检查和验证。对于单元测试中单元的含义，一般要根据实际情况去判定其具体含义，如C语言中单元指一个函数，[Java](http://c.biancheng.net/java/) 里单元指一个类，图形化的软件中可以指一个窗口或一个菜单等。总的来说，单元就是人为规定的最小的被测功能模块。

单元测试是在软件开发过程中要进行的最低级别的测试活动，软件的独立单元将在与程序的其他部分相隔离的情况下进行测试。

#### 单元测试—测试和验证代码的框架

要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时需要让文件必须以`_test`结尾。默认的情况下，`go test`命令不需要任何的参数，它会自动把你源码包下面所有 test 文件测试完毕，当然你也可以带上参数。

这里介绍几个常用的参数：

- -bench regexp 执行相应的 benchmarks，例如 -bench=.；
- -cover 开启测试覆盖率；
- -run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数；
- -v 显示测试的详细命令。


单元测试源码文件可以由多个测试用例组成，每个测试用例函数需要以`Test`为前缀，例如：

func TestXXX( t *testing.T )

- 测试用例文件不会参与正常源码编译，不会被包含到可执行文件中。
- 测试用例文件使用`go test`指令来执行，没有也不需要 main() 作为函数入口。所有在以`_test`结尾的源码内以`Test`开头的函数会自动被执行。
- 测试用例可以不传入 *testing.T 参数。

helloworld 的测试代码（具体位置是`./src/chapter11/gotest/helloworld_test.go`）：

```
package code11_3
import "testing"
func TestHelloWorld(t *testing.T) {
    t.Log("hello world")
}
```

- 单元测试文件 (*_test.go) 里的测试入口必须以 Test 开始，参数为 *testing.T 的函数。一个单元测试文件可以有多个测试入口。
- 使用 testing 包的 T 结构提供的 Log() 方法打印字符串。

##### 1) 单元测试命令行

单元测试使用 go test 命令启动，例如：

```go
$ go test helloworld_test.go
ok          command-line-arguments        0.003s
$ go test -v helloworld_test.go
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
        helloworld_test.go:8: hello world
PASS
ok          command-line-arguments        0.004s
```

代码说明如下：

- 第 1 行，在 go test 后跟 helloworld_test.go 文件，表示测试这个文件里的所有测试用例。
- 第 2 行，显示测试结果，ok 表示测试通过，command-line-arguments 是测试用例需要用到的一个包名，0.003s 表示测试花费的时间。
- 第 3 行，显示在附加参数中添加了`-v`，可以让测试时显示详细的流程。
- 第 4 行，表示开始运行名叫 TestHelloWorld 的测试用例。
- 第 5 行，表示已经运行完 TestHelloWorld 的测试用例，PASS 表示测试成功。
- 第 6 行打印字符串 hello world。

##### 2) 运行指定单元测试用例

`go test`指定文件时默认执行文件内的所有测试用例。可以使用`-run`参数选择需要的测试用例单独执行，参考下面的代码。

一个文件包含多个测试用例（具体位置是`./src/chapter11/gotest/select_test.go`）

```
package code11_3
import "testing"
func TestA(t *testing.T) {
    t.Log("A")
}
func TestAK(t *testing.T) {
    t.Log("AK")
}
func TestB(t *testing.T) {
    t.Log("B")
}
func TestC(t *testing.T) {
    t.Log("C")
}
```

这里指定 TestA 进行测试：

```go
$ go test -v -run TestA select_test.go
=== RUN   TestA
--- PASS: TestA (0.00s)
        select_test.go:6: A
=== RUN   TestAK
--- PASS: TestAK (0.00s)
        select_test.go:10: AK
PASS
ok          command-line-arguments        0.003s
```

TestA 和 TestAK 的测试用例都被执行，原因是`-run`跟随的测试用例的名称支持正则表达式，使用`-run TestA$`即可只执行 TestA 测试用例。

##### 3) 标记单元测试结果

当需要终止当前测试用例时，可以使用 FailNow，参考下面的代码。

测试结果标记（具体位置是`./src/chapter11/gotest/fail_test.go`）

```go
func TestFailNow(t *testing.T) {
    t.FailNow()
}
```

还有一种只标记错误不终止测试的方法，代码如下：

```go
func TestFail(t *testing.T) {
    fmt.Println("before fail")
    t.Fail()
    fmt.Println("after fail")
}
```

测试结果如下：

```go
=== RUN   TestFail
before fail
after fail
--- FAIL: TestFail (0.00s)
FAIL
exit status 1
FAIL        command-line-arguments        0.002s
```

从日志中看出，第 5 行调用 Fail() 后测试结果标记为失败，但是第 7 行依然被程序执行了。

##### 4) 单元测试日志

每个测试用例可能并发执行，使用 testing.T 提供的日志输出可以保证日志跟随这个测试上下文一起打印输出。testing.T 提供了几种日志输出方法，详见下表所示。



| 方  法 | 备  注                           |
| ------ | -------------------------------- |
| Log    | 打印日志，同时结束测试           |
| Logf   | 格式化打印日志，同时结束测试     |
| Error  | 打印错误日志，同时结束测试       |
| Errorf | 格式化打印错误日志，同时结束测试 |
| Fatal  | 打印致命日志，同时结束测试       |
| Fatalf | 格式化打印致命日志，同时结束测试 |


开发者可以根据实际需要选择合适的日志。

#### 基准测试—获得代码内存占用和运行效率的性能数据

基准测试可以测试一段程序的运行性能及耗费 CPU 的程度。Go语言中提供了基准测试框架，使用方法类似于单元测试，使用者无须准备高精度的计时器和各种分析工具，基准测试本身即可以打印出非常标准的测试报告。

##### 1) 基础测试基本使用

下面通过一个例子来了解基准测试的基本使用方法。

基准测试（具体位置是`./src/chapter11/gotest/benchmark_test.go`）

```go
package code11_3
import "testing"
func Benchmark_Add(b *testing.B) {
    var n int
    for i := 0; i < b.N; i++ {
        n++
    }
}
```

这段代码使用基准测试框架测试加法性能。第 7 行中的 b.N 由基准测试框架提供。测试代码需要保证函数可重入性及无状态，也就是说，测试代码不使用全局变量等带有记忆性质的数据结构。避免多次运行同一段代码时的环境不一致，不能假设 N 值范围。

使用如下命令行开启基准测试：

```go
$ go test -v -bench=. benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Add-4           20000000         0.33 ns/op
PASS
ok          command-line-arguments        0.700s
```

代码说明如下：

- 第 1 行的`-bench=.`表示运行 benchmark_test.go 文件里的所有基准测试，和单元测试中的`-run`类似。
- 第 4 行中显示基准测试名称，2000000000 表示测试的次数，也就是 testing.B 结构中提供给程序使用的 N。“0.33 ns/op”表示每一个操作耗费多少时间（纳秒）。


注意：Windows 下使用 go test 命令行时，`-bench=.`应写为`-bench="."`。

##### 2) 基准测试原理

基准测试框架对一个测试用例的默认测试时间是 1 秒。开始测试时，当以 Benchmark 开头的基准测试用例函数返回时还不到 1 秒，那么 testing.B 中的 N 值将按 1、2、5、10、20、50……递增，同时以递增后的值重新调用基准测试用例函数。

##### 3) 自定义测试时间

通过`-benchtime`参数可以自定义测试时间，例如：

```go
$ go test -v -bench=. -benchtime=5s benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Add-4           10000000000                 0.33 ns/op
PASS
ok          command-line-arguments        3.380s
```

##### 4) 测试内存

基准测试可以对一段代码可能存在的内存分配进行统计，下面是一段使用字符串格式化的函数，内部会进行一些分配操作。

```go
func Benchmark_Alloc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("%d", i)
    }
}
```

在命令行中添加-benchmem参数以显示内存分配情况，参见下面的指令：

```go
$ go test -v -bench=Alloc -benchmem benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Alloc-4 20000000 109 ns/op 16 B/op 2 allocs/op
PASS
ok          command-line-arguments        2.311s
```

代码说明如下：

- 第 1 行的代码中`-bench`后添加了 Alloc，指定只测试 Benchmark_Alloc() 函数。
- 第 4 行代码的“16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配。


开发者根据这些信息可以迅速找到可能的分配点，进行优化和调整。

##### 5) 控制计时器

有些测试需要一定的启动和初始化时间，如果从 Benchmark() 函数开始计时会很大程度上影响测试结果的精准性。testing.B 提供了一系列的方法可以方便地控制计时器，从而让计时器只在需要的区间进行测试。我们通过下面的代码来了解计时器的控制。

基准测试中的计时器控制（具体位置是`./src/chapter11/gotest/benchmark_test.go`）：

```go
func Benchmark_Add_TimerControl(b *testing.B) {
    // 重置计时器
    b.ResetTimer()
    // 停止计时器
    b.StopTimer()
    // 开始计时器
    b.StartTimer()
    var n int
    for i := 0; i < b.N; i++ {
        n++
    }
}
```

从 Benchmark() 函数开始，Timer 就开始计数。StopTimer() 可以停止这个计数过程，做一些耗时的操作，通过 StartTimer() 重新开始计时。ResetTimer() 可以重置计数器的数据。

计数器内部不仅包含耗时数据，还包括内存分配的数据。

