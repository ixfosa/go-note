## 概述

Go语言提供了一种机制在运行时更新和检查变量的值、调用变量的方法和变量支持的内在操作，但是在编译时并不知道这些变量的具体类型，这种机制被称为反射。反射也可以让我们将类型本身作为第一类的值类型处理。

反射是指在**程序运行期对程序本身进行访问和修改的能力**，程序在编译时变量被转换为内存地址，变量名不会被编译器写入到可执行部分，在运行程序时程序无法获取自身的信息。

支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

Go语言程序的反射系统无法获取到一个可执行文件空间中或者是一个包中的所有类型信息，需要配合使用标准库中对应的`词法`、`语法解析器`和`抽象语法树`（AST）对源码进行扫描后获得这些信息。

Go语言提供了 `reflect` 包来访问程序的反射信息。

**reflect 包**

Go语言中的反射是由 reflect 包提供支持的，它定义了两个重要的类型 `Type `和 `Value` 任意接口值在反射中都可以理解为由` reflect.Type `和 `reflect.Value` 两部分组成，并且 reflect 包提供了` reflect.TypeOf` 和 `reflect.ValueOf `两个函数来获取任意对象的 Value 和 Type。

`reflect `实现了运行时的反射能力，能够让 Golang 的程序操作不同类型的对象，我们可以使用包中的函数 `TypeOf `从静态类型` interface{} `中获取动态类型信息并通过` ValueOf `获取数据的`运行时`表示，通过这两个函数和包中的其他工具我们就可以得到更强大的表达能力。

在具体介绍反射包的实现原理之前，我们先要对 Go 语言的反射有一些比较简单的理解，首先 reflect 中有两对非常重要的函数和类型，我们在上面已经介绍过其中的两个函数 TypeOf 和 ValueOf，另外两个类型是 `Type` 和` Value`，它们与函数是一一对应的关系：

![Go 语言反射的实现原理](E:/smile/go/images/Go 语言反射的实现原理.png)



## 反射类型与种类

在使用反射时，需要首先理解`类型`（Type）和 `种类`（Kind）的区别。编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。例如需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。

### 概述

类型 `Type` 是 Golang 反射包中定义的一个接口，我们可以使用 `TypeOf` 函数获取任意值的变量的的类型，我们能从这个接口中看到非常多有趣的方法，`MethodByName` 可以获取当前类型对应方法的引用、`Implements` 可以判断当前类型是否实现了某个接口：

```go
type Type interface {
        Align() int
        FieldAlign() int
        Method(int) Method
        MethodByName(string) (Method, bool)
        NumMethod() int
        Name() string
        PkgPath() string
        Size() uintptr
        String() string
        Kind() Kind
        Implements(u Type) bool
        ...
}
```

反射包中 `Value` 的类型却与 `Type` 不同，`Type` 是一个接口类型，但是 `Value` 在 [`reflect`](https://golang.org/pkg/reflect/) 包中的定义是一个结构体，这个结构体没有任何对外暴露的成员变量，但是却提供了很多方法让我们获取或者写入 `Value` 结构体中存储的数据：

```go
type Value struct {
        // contains filtered or unexported fields
}

func (v Value) Addr() Value
func (v Value) Bool() bool
func (v Value) Bytes() []byte
func (v Value) Float() float64
...

```

反射包中的所有方法基本都是围绕着 `Type `和` Value` 这两个对外暴露的类型设计的，我们通过 `TypeOf`、`ValueOf `方法就可以将一个普通的变量转换成『反射』包中提供的 Type 和 Value，使用反射提供的方法对这些类型进行复杂的操作。

在Go语言程序中，使用` reflect.TypeOf() `函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息，下面通过示例来理解获取类型对象的过程：

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    
    var a int
    
    //通过 reflect.TypeOf() 取得变量 a 的类型对象 typeOfA，类型为 reflect.Type()。
    typeOfA := reflect.TypeOf(a)
    
    //通过 typeOfA 类型对象的成员函数，可以分别获取到 typeOfA 变量的类型名为 int，种类（Kind）为 int。
    fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

//运行结果如下：
int  int
```

### 反射种类（Kind）的定义

Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。

种类（Kind）指的是对象归属的品种，在 reflect 包中有如下定义：

```go
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
```

Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

### 类型对象中获取类型名称和种类

Go语言中的类型名称对应的反射获取方法是` reflect.Type `中的 `Name()` 方法，返回表示`类型名称的字符串`；类型归属的`种类`（Kind）使用的是 reflect.Type 中的` Kind() `方法，返回 reflect.Kind 类型的常量。

下面的代码中会对常量和结构体进行类型信息获取。

```go
package main

import (
    "fmt"
    "reflect"
)

// 定义一个Enum类型
type Enum int

const (
    Zero Enum = 0
)

func main() {
    
    // 声明一个空结构体
    type cat struct {
    }
    
    // 将 cat 实例化，并且使用 reflect.TypeOf() 获取被实例化后的 cat 的反射类型对象。
    typeOfCat := reflect.TypeOf(cat{})
    
    //输出 cat 的类型名称和种类，类型名称就是 cat，而 cat 属于一种结构体种类，因此种类为 struct。
    fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
    
    // 获取Zero常量的反射类型对象
    typeOfA := reflect.TypeOf(Zero)
    
    // 显示反射类型对象的名称和种类
    fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

//运行结果如下：
cat struct
Enum int
```

### 指针与指针指向的元素

Go语言程序中对指针获取反射对象时，可以通过` reflect.Elem() `方法获取这个指针指向的元素类型，这个获取过程被称为取元素，等效于对指针类型变量做了一个`*`操作，代码如下：

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    
    // 声明一个空结构体
    type cat struct {
    }
    
    // 创建了 cat 结构体的实例，ins 是一个 *cat 类型的指针变量。
    ins := &cat{}
    
    // 对指针变量获取反射类型信息。
    typeOfCat := reflect.TypeOf(ins)
    
    //输出指针变量的类型名称和种类。
    //反射中对所有指针变量的种类都是 Ptr，但需要注意的是，指针变量的类型名称是空，不是 *cat。
    fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
    
    //取指针类型的元素类型
    //也就是 cat 类型。这个操作不可逆，不可以通过一个非指针类型获取它的指针类型。
    typeOfCat = typeOfCat.Elem()
    
    // 显示反射类型对象的名称和种类
    fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}

//运行结果如下：
name:'' kind:'ptr'
element name: 'cat', element kind: 'struct'
```



## 反射法则

### Go语言中的类型

Go语言是一门静态类型的语言，每个变量都有一个静态类型，类型在编译的时候确定下来。

```go
type MyInt int

var i int
var j MyInt
```

变量 i 的类型是 int，变量 j 的类型是 MyInt，虽然它们有着相同的基本类型，但静态类型却不一样，在没有类型转换的情况下，它们之间无法互相赋值。

接口是一个重要的类型，它意味着一个确定的方法集合，一个接口变量可以存储任何实现了接口的方法的具体值（除了接口本身），例如 io.Reader 和 io.Writer：

```go
// Reader is the interface that wraps the basic Read method.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

如果一个类型声明实现了 Reader（或 Writer）方法，那么它便实现了 io.Reader（或 io.Writer），这意味着一个 io.Reader 的变量可以持有任何一个实现了 Read 方法的的类型的值。

```go
var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
// and so on
```

必须要弄清楚的一点是，不管变量 r 中的具体值是什么，r 的类型永远是 io.Reader，由于Go语言是静态类型的，r 的静态类型就是 io.Reader。

在接口类型中有一个极为重要的例子——空接口：

```go
interface{}
```

它表示了一个空的方法集，一切值都可以满足它，因为它们都有零值或方法。

有人说Go语言的接口是动态类型，这是错误的，它们都是静态类型，虽然在运行时中，接口变量存储的值也许会变，但接口变量的类型是不会变的。我们必须精确地了解这些，因为反射与接口是密切相关的。

运行时反射是程序在运行期间检查其自身结构的一种方式，它是元编程 的一种，但是它带来的灵活性也是一把双刃剑，过量的使用反射会使我们的程序逻辑变得难以理解并且运行缓慢，我们在这一节中就会介绍 Go 语言反射的三大法则，这能够帮助我们更好地理解反射的作用。

1. 从接口值可反射出反射对象；
2. 从反射对象可反射出接口值；
3. 要修改反射对象，其值必须可设置；

### 第一法则

反射的第一条法则就是，我们能够将 Go 语言中的接口类型变量转换成反射对象，上面提到的reflect.TypeOf 和 reflect.ValueOf 就是完成这个转换的两个最重要方法，如果我们认为 Go 语言中的类型和反射类型是两个不同『世界』的话，那么这两个方法就是连接这两个世界的桥梁。

> 注：这里反射类型指 reflect.Type 和 reflect.Value。

![第一法则](E:/smile/go/images/第一法则.png)

我们通过以下例子简单介绍这两个方法的作用，其中 `TypeOf` 获取了变量 `author` 的类型也就是 `string` 而 `ValueOf` 获取了变量的值 `draven`，如果我们知道了一个变量的类型和值，那么也就意味着我们知道了关于这个变量的全部信息。

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    author := "draven"
    fmt.Println("TypeOf author:", reflect.TypeOf(author))
    fmt.Println("ValueOf author:", reflect.ValueOf(author))
}

$ go run main.go
TypeOf author: string
ValueOf author: draven
```

从变量的类型上我们可以获当前类型能够执行的方法 `Method` 以及当前类型实现的接口等信息；

- 对于结构体，可以获取字段的数量并通过下标和字段名获取字段 `StructField`；
- 对于哈希表，可以获取哈希表的 `Key` 类型；
- 对于函数或方法，可以获得入参和返回值的类型；
- …

总而言之，使用 `TypeOf` 和 `ValueOf` 能够将 Go 语言中的变量转换成反射对象，在这时我们能够获得几乎一切跟当前类型相关数据和操作，然后就可以用这些运行时获取的结构动态的执行一些方法。

> 为什么是从**接口**到反射对象，如果直接调用 `reflect.ValueOf(1)`，看起来是从基本类型 `int` 到反射类型，但是 `TypeOf` 和 `ValueOf` 两个方法的入参其实是 `interface{}` 类型。
> Go 语言的函数调用都是值传递的，变量会在方法调用前进行类型转换，也就是 `int` 类型的基本变量会被转换成 `interface{}` 类型，这也就是第一条法则介绍的是从接口到反射对象。



类型 reflect.Value 有一个方法` Type()`，它会返回一个 reflect.Type 类型的对象。

Type 和 Value 都有一个名为` Kind `的方法，它会返回一个常量，表示底层数据的类型，常见值有：Uint、Float64、Slice 等。

Value 类型也有一些类似于 Int、Float 的方法，用来提取底层的数据：

- Int 方法用来提取 int64
- Float 方法用来提取 float64，示例代码如下：

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    v := reflect.ValueOf(x)
    fmt.Println("type:", v.Type())
    fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
    fmt.Println("value:", v.Float())
}

//运行结果如下：
type: float64
kind is float64: true
value: 3.4
```

还有一些用来修改数据的方法，比如 SetInt、SetFloat。在介绍它们之前，我们要先理解“可修改性”（settability）

首先是介绍下 Value 的 `getter` 和` setter` 方法，为了保证 API 的精简，这两个方法操作的是某一组类型范围最大的那个。比如，处理任何含符号整型数，都使用 int64，也就是说 Value 类型的 Int 方法返回值为 int64 类型，SetInt 方法接收的参数类型也是 int64 类型。实际使用时，可能需要转化为实际的类型：

```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    var x uint8 = 'x'
    v := reflect.ValueOf(x)
    fmt.Println("type:", v.Type())                            // uint8.
    fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
    x = uint8(v.Uint())            // v.Uint returns a uint64.
}

//运行结果如下：
type: uint8
kind is uint8: true
```

其次，反射对象的 Kind 方法描述的是基础类型，而不是静态类型。如果一个反射对象包含了用户定义类型的值，如下所示：

```go
type MyInt int
var x MyInt = 7
v := reflect.ValueOf(x)
```

上面的代码中，虽然变量 v 的静态类型是 MyInt，而不是 int，但 Kind 方法仍然会返回 reflect.Int。换句话说 Kind 方法不会像 Type 方法一样区分 MyInt 和 int。

### 第二法则

我们既然能够将接口类型的变量转换成反射对象类型，那么也需要一些其他方法将反射对象还原成成接口类型的变量， `reflect `中的` Interface`方法就能完成这项工作：

![第二法则](E:/smile/go/images/第二法则.png)

然而调用 `Interface` 方法我们也只能获得 `interface{}` 类型的接口变量，如果想要将其还原成原本的类型还需要经过一次强制的类型转换，如下所示：

```go
v := reflect.ValueOf(1)
v.Interface{}.(int)
```

这个过程就像从接口值到反射对象的镜面过程一样，从接口值到反射对象需要经过从基本类型到接口类型的类型转换和从接口类型到反射对象类型的转换，反过来的话，所有的反射对象也都需要先转换成接口类型，再通过强制类型转换变成原始类型：

![第二法则2](E:/smile/go/images/第二法则2.png)

当然不是所有的变量都需要类型转换这一过程，如果本身就是 `interface{}` 类型的，那么它其实并不需要经过类型转换，对于大多数的变量来说，类型转换这一过程很多时候都是隐式发生的，只有在我们需要将反射对象转换回基本类型时才需要做显示的转换操作。

### 第三法则

Go 语言反射的最后一条法则是与值是否可以被更改相关的，如果我们想要更新一个 `reflect.Value`，那么它持有的值一定是可以被更新的，假设我们有以下代码：

```go
func main() {
    i := 1
    v := reflect.ValueOf(i)
    v.SetInt(10)
    fmt.Println(i)
}

$ go run reflect.go
panic: reflect: reflect.flag.mustBeAssignable using unaddressable value

goroutine 1 [running]:
reflect.flag.mustBeAssignableSlow(0x82, 0x1014c0)
    /usr/local/go/src/reflect/value.go:247 +0x180
reflect.flag.mustBeAssignable(...)
    /usr/local/go/src/reflect/value.go:234
reflect.Value.SetInt(0x100dc0, 0x414020, 0x82, 0x1840, 0xa, 0x0)
    /usr/local/go/src/reflect/value.go:1606 +0x40
main.main()
    /tmp/sandbox590309925/prog.go:11 +0xe0
```

运行上述代码时会导致程序 panic 并报出 `reflect: reflect.flag.mustBeAssignable using unaddressable value` 错误，仔细想一下其实能够发现出错的原因，Go 语言的 [函数调用]都是传值的，所以我们得到的反射对象其实跟最开始的变量没有任何关系，没有任何变量持有复制出来的值，所以直接对它修改会导致崩溃。

想要修改原有的变量我们只能通过如下所示的方法，首先通过 `reflect.ValueOf` 获取变量指针，然后通过 `Elem` 方法获取指针指向的变量并调用 `SetInt` 方法更新变量的值：

```go
func main() {
    i := 1
    v := reflect.ValueOf(&i)
    v.Elem().SetInt(10)
    fmt.Println(i)
}

$ go run reflect.go
10
```

这种获取指针对应的 `reflect.Value` 并通过 `Elem` 方法迂回的方式就能够获取到可以被设置的变量，这一复杂的过程主要也是因为 Go 语言的函数调用都是值传递的，我们可以将上述代码理解成：

```go
func main() {
    i := 1
    v := &i
    *v = 10
}
```

如果不能直接操作 `i` 变量修改其持有的值，我们就只能获取 `i` 变量所在地址并使用 `*v` 修改所在地址中存储的整数。



## 实现原理

### 类型和值

Golang 的 `interface{}` 类型在语言内部都是通过 `emptyInterface` 这个结体来表示的，其中包含一个 `rtype` 字段用于表示变量的类型以及一个 `word` 字段指向内部封装的数据：

```go
type emptyInterface struct {
    typ  *rtype
    word unsafe.Pointer
}
```

用于获取变量类型的 `TypeOf` 函数就是将传入的 `i` 变量强制转换成 `emptyInterface` 类型并获取其中存储的类型信息 `rtype`：

```go
func TypeOf(i interface{}) Type {
    eface := *(*emptyInterface)(unsafe.Pointer(&i))
    return toType(eface.typ)
}

func toType(t *rtype) Type {
    if t == nil {
        return nil
    }
    return t
}
```

> `unsafe.Pointer`是一种特殊意义的指针，它可以包含任意类型的地址
>
> 关于`unsafe.Pointer`的4个规则。
>
> 1. 任何指针都可以转换为`unsafe.Pointer`
> 2. `unsafe.Pointer`可以转换为任何指针
> 3. `uintptr`可以转换为`unsafe.Pointer`
> 4. `unsafe.Pointer`可以转换为`uintptr`



`rtype` 就是一个实现了 `Type` 接口的接口体，我们能在 [`reflect`](https://golang.org/pkg/reflect/) 包中找到如下所示的 `Name` 方法帮助我们获取当前类型的名称等信息：

```go
func (t *rtype) String() string {
    s := t.nameOff(t.str).name()
    if t.tflag&tflagExtraStar != 0 {
        return s[1:]
    }
    return s
}
```

`TypeOf` 函数的实现原理其实并不复杂，它只是将一个 `interface{}` 变量转换成了内部的 `emptyInterface` 表示，然后从中获取相应的类型信息。

用于获取接口值 `Value` 的函数 `ValueOf` 实现也非常简单，在该函数中我们先调用了 `escapes` 函数保证当前值逃逸到堆上，然后通过 `unpackEface` 方法从接口中获取 `Value` 结构体：

```go
func ValueOf(i interface{}) Value {
    if i == nil {
        return Value{}
    }

    escapes(i)

    return unpackEface(i)
}

func unpackEface(i interface{}) Value {
    e := (*emptyInterface)(unsafe.Pointer(&i))
    t := e.typ
    if t == nil {
        return Value{}
    }
    f := flag(t.Kind())
    if ifaceIndir(t) {
        f |= flagIndir
    }
    return Value{t, e.word, f}
}
```

`unpackEface` 函数会将传入的接口 `interface{}` 转换成 `emptyInterface` 结构体然后将其中表示接口值类型、指针以及值的类型包装成 `Value` 结构体并返回。

`TypeOf` 和 `ValueOf` 两个方法的实现其实都非常简单，从一个 Go 语言的基本变量中获取反射对象以及类型的过程中，`TypeOf` 和 `ValueOf` 两个方法的执行过程并不是特别的复杂，我们还需要注意基本变量到接口值的转换过程：

```go
package main

import (
    "reflect"
)

func main() {
    i := 20
    _ = reflect.TypeOf(i)
}

$ go build -gcflags="-S -N" main.go
...
MOVQ    $20, ""..autotmp_20+56(SP) // autotmp = 20
LEAQ    type.int(SB), AX           // AX = type.int(SB)
MOVQ    AX, ""..autotmp_19+280(SP) // autotmp_19+280(SP) = type.int(SB)
LEAQ    ""..autotmp_20+56(SP), CX  // CX = 20
MOVQ    CX, ""..autotmp_19+288(SP) // autotmp_19+288(SP) = 20
...
```

我们使用 `-S -N` 编译指令编译了上述代码，从这段截取的汇编语言中我们可以发现，在函数调用之前其实发生了类型转换，我们将 `int` 类型的变量转换成了占用 16 字节 `autotmp_19+280(SP) ~ autotmp_19+288(SP)` 的 `interface{}` 结构体，两个 `LEAQ` 指令分别获取了类型的指针 `type.int(SB)` 以及变量 `i` 所在的地址。

总的来说，在 Go 语言的编译期间我们就完成了类型转换的工作，将变量的类型和值转换成了 `interface{}` 等待运行期间使用 [`reflect`](https://golang.org/pkg/reflect/) 包获取其中存储的信息。

### 更新变量

当我们想要更新一个 `reflect.Value` 时，就需要调用 `Set` 方法更新反射对象，该方法会调用 `mustBeAssignable` 和 `mustBeExported` 分别检查当前反射对象是否是可以被设置的和对外暴露的公开字段：

```go
func (v Value) Set(x Value) {
    v.mustBeAssignable()
    x.mustBeExported() // do not let unexported x leak
    var target unsafe.Pointer
    if v.kind() == Interface {
        target = v.ptr
    }
    x = x.assignTo("reflect.Set", v.typ, target)
    if x.flag&flagIndir != 0 {
        typedmemmove(v.typ, v.ptr, x.ptr)
    } else {
        *(*unsafe.Pointer)(v.ptr) = x.ptr
    }
}
Set 
```

`Set` 方法中会调用 `assignTo`，该方法会返回一个新的 `reflect.Value` 反射对象，我们可以将反射对象的指针直接拷贝到被设置的反射变量上：

```go
func (v Value) assignTo(context string, dst *rtype, target unsafe.Pointer) Value {
    if v.flag&flagMethod != 0 {
        v = makeMethodValue(context, v)
    }

    switch {
    case directlyAssignable(dst, v.typ):
        fl := v.flag&(flagAddr|flagIndir) | v.flag.ro()
        fl |= flag(dst.Kind())
        return Value{dst, v.ptr, fl}

    case implements(dst, v.typ):
        if target == nil {
            target = unsafe_New(dst)
        }
        if v.Kind() == Interface && v.IsNil() {
            return Value{dst, nil, flag(Interface)}
        }
        x := valueInterface(v, false)
        if dst.NumMethod() == 0 {
            *(*interface{})(target) = x
        } else {
            ifaceE2I(dst, x, target)
        }
        return Value{dst, target, flagIndir | flag(Interface)}
    }

    panic(context + ": value of type " + v.typ.String() + " is not assignable to type " + dst.String())
}
```

`assignTo` 会根据当前和被设置的反射对象类型创建一个新的 `Value` 结构体，当两个反射对象的类型是可以被直接替换时，就会直接将目标反射对象返回；如果当前反射对象是接口并且目标对象实现了接口，就会将目标对象简单包装成接口值，上述方法返回反射对象的 `ptr` 最终会覆盖当前反射对象中存储的值

### 实现协议

[`reflect`](https://golang.org/pkg/reflect/) 包还为我们提供了 `Implements` 方法用于判断某些类型是否遵循协议实现了全部的方法，在 Go 语言中想要获取结构体的类型还是比较容易的，但是想要获得接口的类型就需要比较黑魔法的方式：

```go
reflect.TypeOf((*<interface>)(nil)).Elem()
```

只有通过上述方式才能获得一个接口类型的反射对象，假设我们有以下代码，我们需要判断 `CustomError` 是否实现了 Go 语言标准库中的 `error` 协议：

```go
type CustomError struct{}

func (*CustomError) Error() string {
    return ""
}

func main() {
    typeOfError := reflect.TypeOf((*error)(nil)).Elem()
    customErrorPtr := reflect.TypeOf(&CustomError{})
    customError := reflect.TypeOf(CustomError{})

    fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
    fmt.Println(customError.Implements(typeOfError)) // #=> false
}
```

运行上述代码我们会发现 `CustomError` 类型并没有实现 `error` 接口，而 `*CustomError` 指针类型却实现了接口,可以使用结构体和指针两种不同的类型实现接口。

```go
 func (t *rtype) Implements(u Type) bool {
    if u == nil {
        panic("reflect: nil type passed to Type.Implements")
    }
    if u.Kind() != Interface {
        panic("reflect: non-interface type passed to Type.Implements")
    }
    return implements(u.(*rtype), t)
}
```

`Implements` 方法会检查传入的类型是不是接口，如果不是接口或者是空值就会直接 panic 中止当前程序，否则就会调用私有的函数 `implements` 判断类型之间是否有实现关系

```go
func implements(T, V *rtype) bool {
    t := (*interfaceType)(unsafe.Pointer(T))
    if len(t.methods) == 0 {
        return true
    }

    // ...

    v := V.uncommon()
    i := 0
    vmethods := v.methods()
    for j := 0; j < int(v.mcount); j++ {
        tm := &t.methods[i]
        tmName := t.nameOff(tm.name)
        vm := vmethods[j]
        vmName := V.nameOff(vm.name)
        if vmName.name() == tmName.name() && V.typeOff(vm.mtyp) == t.typeOff(tm.typ) {
            if i++; i >= len(t.methods) {
                return true
            }
        }
    }
    return false
}
```

如果接口中不包含任何方法，也就意味着这是一个空的 `interface{}`，任意的类型都可以实现该协议，所以就会直接返回 `true`。

![实现协议](E:/smile/go/images/实现协议.png)

在其他情况下，由于方法是按照一定顺序排列的，`implements` 中就会维护两个用于遍历接口和类型方法的索引 `i` 和 `j`，所以整个过程的实现复杂度是 `O(n+m)`，最多只会进行 `n + m` 次数的比较，不会出现次方级别的复杂度。

## 方法调用

作为一门静态语言，如果我们想要通过 [`reflect`](https://golang.org/pkg/reflect/) 包利用反射在运行期间执行方法并不是一件容易的事情，下面的代码就使用了反射来执行 `Add(0, 1)` 这一表达式：

```go
func Add(a, b int) int { return a + b }

func main() {
    v := reflect.ValueOf(Add)
    if v.Kind() != reflect.Func {
        return
    }
    t := v.Type()
    argv := make([]reflect.Value, t.NumIn())
    for i := range argv {
        if t.In(i).Kind() != reflect.Int {
            return
        }
        argv[i] = reflect.ValueOf(i)
    }
    result := v.Call(argv)
    if len(result) != 1 || result[0].Kind() != reflect.Int {
        return
    }
    fmt.Println(result[0].Int()) // #=> 1
}
```

1. 通过 `reflect.ValueOf` 获取函数 `Add` 对应的反射对象；
2. 根据反射对象 `NumIn` 方法返回的参数个数创建 `argv` 数组；
3. 多次调用 `reflect.Value` 逐一设置 `argv` 数组中的各个参数；
4. 调用反射对象 `Add` 的 `Call` 方法并传入参数列表；
5. 获取返回值数组、验证数组的长度以及类型并打印其中的数据；

使用反射来调用方法非常复杂，原本只需要一行代码就能完成的工作，现在需要 10 多行代码才能完成，但是这也是在静态语言中使用这种动态特性需要付出的成本，理解这个调用过程能够帮助我们深入理解 Go 语言函数和方法调用的原理。

```go
func (v Value) Call(in []Value) []Value {
    v.mustBe(Func)
    v.mustBeExported()
    return v.call("Call", in)
}
```

`Call` 作为反射包运行时调用方法的入口，通过两个 `MustBe` 方法保证了当前反射对象的类型和可见性，随后调用 `call` 方法完成运行时方法调用的过程，这个过程会被分成以下的几个部分：

1. 检查输入参数的合法性以及类型等信息；
2. 将传入的 `reflect.Value` 参数数组设置到栈上；
3. 通过函数指针和输入参数调用函数；
4. 从栈上获取函数的返回值；

我们将按照上面的顺序依次详细介绍使用 [`reflect`](https://golang.org/pkg/reflect/) 进行函数调用的几个过程。

### 参数检查

参数检查是通过反射调用方法的第一步，在参数检查期间我们会从反射对象中取出当前的函数指针 `unsafe.Pointer`，如果待执行的函数是方法，就会通过 `methodReceiver` 函数获取方法的接受者和函数指针。

```go
func (v Value) call(op string, in []Value) []Value {
    t := (*funcType)(unsafe.Pointer(v.typ))
    var (
        fn       unsafe.Pointer
        rcvr     Value
        rcvrtype *rtype
    )
    if v.flag&flagMethod != 0 {
        rcvr = v
        rcvrtype, t, fn = methodReceiver(op, v, int(v.flag)>>flagMethodShift)
    } else if v.flag&flagIndir != 0 {
        fn = *(*unsafe.Pointer)(v.ptr)
    } else {
        fn = v.ptr
    }

    n := t.NumIn()
    if len(in) < n {
        panic("reflect: Call with too few input arguments")
    }
    if len(in) > n {
        panic("reflect: Call with too many input arguments")
    }
    for i := 0; i < n; i++ {
        if xt, targ := in[i].Type(), t.In(i); !xt.AssignableTo(targ) {
            panic("reflect: " + op + " using " + xt.String() + " as type " + targ.String())
        }
    }

    nin := len(in)
    if nin != t.NumIn() {
        panic("reflect.Value.Call: wrong argument count")
    }
```

除此之外，在参数检查的过程中我们还会检查当前传入参数的个数以及所有参数的类型是否能被传入该函数中，任何参数不匹配的问题都会导致当前函数直接 panic 并中止整个程序。

### 准备参数

当我们已经对当前方法的参数验证完成之后，就会进入函数调用的下一个阶段，为函数调用准备参数，Go 语言的函数调用的惯例，所有的参数都会被依次放置到堆栈上。

```go
    nout := t.NumOut()
    frametype, _, retOffset, _, framePool := funcLayout(t, rcvrtype)

    var args unsafe.Pointer
    if nout == 0 {
        args = framePool.Get().(unsafe.Pointer)
    } else {
        args = unsafe_New(frametype)
    }
    off := uintptr(0)

    if rcvrtype != nil {
        storeRcvr(rcvr, args)
        off = ptrSize
    }
    for i, v := range in {
        targ := t.In(i).(*rtype)
        a := uintptr(targ.align)
        off = (off + a - 1) &^ (a - 1)
        n := targ.size
        if n == 0 {
            v.assignTo("reflect.Value.Call", targ, nil)
            continue
        }
        addr := add(args, off, "n > 0")
        v = v.assignTo("reflect.Value.Call", targ, addr)
        if v.flag&flagIndir != 0 {
            typedmemmove(targ, addr, v.ptr)
        } else {
            *(*unsafe.Pointer)(addr) = v.ptr
        }
        off += n
    }

```

1. 通过 `funcLayout` 函数计算当前函数需要的参数和返回值的堆栈布局，也就是每一个参数和返回值所占的空间大小；
2. 如果当前函数有返回值，需要为当前函数的参数和返回值分配一片内存空间 `args`；
3. 如果当前函数是方法，需要向将方法的接受者拷贝到 `args` 这片内存中；
4. 将所有函数的参数按照顺序依次拷贝到对应`args`内存中
5. 使用 `funcLayout` 返回的参数计算参数在内存中的位置；
6. 通过 `typedmemmove` 或者寻址的放置拷贝参数；
   准备参数的过程其实就是计算各个参数和返回值占用的内存空间，并将所有的参数都拷贝内存空间对应的位置上。

### 调用函数

准备好调用函数需要的全部参数之后，就会通过以下的表达式开始方法的调用了，我们会向该函数中传入栈类型、函数指针、参数和返回值的内存空间、栈的大小以及返回值的偏移量：

```go
call(frametype, fn, args, uint32(frametype.size), uint32(retOffset))
```

这个函数实际上并不存在，它会在编译期间被链接到 [`runtime.reflectcall`](https://github.com/golang/go/blob/a38a917aee626a9b9d5ce2b93964f586bf759ea0/src/runtime/asm_386.s#L489-L526) 这个用汇编实现的函数上，我们在这里并不会展开介绍该函数的具体实现，感兴趣的读者可以自行了解其实现原理。



### 处理返回值

当函数调用结束之后，我们就会开始处理函数的返回值了，如果函数没有任何返回值我们就会直接清空 `args` 中的全部内容来释放内存空间，不过如果当前函数有返回值就会进入另一个分支：

```go
    var ret []Value
    if nout == 0 {
        typedmemclr(frametype, args)
        framePool.Put(args)
    } else {
        typedmemclrpartial(frametype, args, 0, retOffset)

        ret = make([]Value, nout)
        off = retOffset
        for i := 0; i < nout; i++ {
            tv := t.Out(i)
            a := uintptr(tv.Align())
            off = (off + a - 1) &^ (a - 1)
            if tv.Size() != 0 {
                fl := flagIndir | flag(tv.Kind())
                ret[i] = Value{tv.common(), add(args, off, "tv.Size() != 0"), fl}
            } else {
                ret[i] = Zero(tv)
            }
            off += tv.Size()
        }
    }
    
    return ret
}
```

1. 将 `args` 中与输入参数有关的内存空间清空；
2. 创建一个 `nout` 长度的切片用于保存由反射对象构成的返回值数组；
3. 从函数对象中获取返回值的类型和内存大小，将 `args` 内存中的数据转换成 `reflect.Value` 类型的返回值；

由 `reflect.Value` 构成的 `ret` 数组最终就会被返回到上层，使用反射进行函数调用的过程也就结束了。

## 反射获取值信息

当我们将一个接口值传递给一个` reflect.ValueOf `函数调用时，此调用返回的是代表着此接口值的动态值的一个 reflect.Value 值。我们必须通过间接的途径获得一个代表一个接口值的 reflect.Value 值。

reflect.Value 类型有很多方法（https://golang.google.cn/pkg/reflect/）。我们可以调用这些方法来观察和操纵一个 reflect.Value 属主值表示的 Go 值。这些方法中的有些适用于所有种类类型的值，有些只适用于一种或几种类型的值。

通过不合适的 reflect.Value 属主值调用某个方法将在运行时产生一个恐慌。请阅读 reflect 代码库中各个方法的文档来获取如何正确地使用这些方法。

一个 reflect.Value 值的 CanSet 方法将返回此 reflect.Value 值代表的 Go 值是否可以被修改（可以被赋值）。如果一个 Go 值可以被修改，则我们可以调用对应的 reflect.Value 值的 Set 方法来修改此 Go 值。注意：reflect.ValueOf 函数直接返回的 reflect.Value 值都是不可修改的。

反射不仅可以获取值的类型信息，还可以动态地获取或者设置变量的值。Go语言中使用 reflect.Value 获取和设置变量的值。

### 使用反射值对象包装任意值

Go语言中，使用 reflect.ValueOf() 函数获得值的反射值对象（reflect.Value）。书写格式如下：

```go
value := reflect.ValueOf(rawValue)
```

reflect.ValueOf 返回 reflect.Value 类型，包含有 rawValue 的值信息。reflect.Value 与原值间可以通过值包装和值获取互相转化。reflect.Value 是一些反射操作的重要类型，如反射调用函数。

### 从反射值对象获取被包装的值

Go语言中可以通过 reflect.Value 重新获得原始值。

#### reflect.Value获取值的方法

可以通过下面几种方法从反射值对象 reflect.Value 中获取原值，如下表所示。

| 方法名                   | 说  明                                                       |
| ------------------------ | ------------------------------------------------------------ |
| Interface() interface {} | 将值以 interface{} 类型返回，可以通过类型断言转换为指定类型  |
| Int() int64              | 将值以 int 类型返回，所有有符号整型均可以此方式返回          |
| Uint() uint64            | 将值以 uint 类型返回，所有无符号整型均可以此方式返回         |
| Float() float64          | 将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回 |
| Bool() bool              | 将值以 bool 类型返回                                         |
| Bytes() []bytes          | 将值以字节数组 []bytes 类型返回                              |
| String() string          | 将值以字符串类型返回                                         |

#### reflect.Value中获取值的例子

下面代码中，将整型变量中的值使用 reflect.Value 获取反射值对象（reflect.Value）。再通过 reflect.Value 的 Interface() 方法获得 interface{} 类型的原值，通过 int 类型对应的 reflect.Value 的 Int() 方法获得整型值。

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {

    // 声明整型变量a并赋初值
    var a int = 1024
    
    // 获取变量a的反射值对象
    valueOfA := reflect.ValueOf(a)
    
    // 获取interface{}类型的值, 通过类型断言转换
    var getA int = valueOfA.Interface().(int)
    
    // 获取64位的值, 强制类型转换为int类型
    var getA2 int = int(valueOfA.Int())
    
    fmt.Println(getA, getA2)
}
代码输出如下：
1024 1024
```

## 反射访问结构体成员的值

反射值对象（reflect.Value）提供对结构体访问的方法，通过这些方法可以完成对结构体任意值的访问，如下表所示。

| 方  法                                         | 备  注                                                       |
| ---------------------------------------------- | ------------------------------------------------------------ |
| Field(i int) Value                             | 根据索引，返回索引对应的结构体成员字段的反射值对象。当值不是结构体或索引超界时发生宕机 |
| NumField() int                                 | 返回结构体成员字段数量。当值不是结构体或索引超界时发生宕机   |
| FieldByName(name string) Value                 | 根据给定字符串返回字符串对应的结构体字段。没有找到时返回零值，当值不是结构体或索引超界时发生宕机 |
| FieldByIndex(index []int) Value                | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的值。 没有找到时返回零值，当值不是结构体或索引超界时发生宕机 |
| FieldByNameFunc(match func(string) bool) Value | 根据匹配函数匹配需要的字段。找到时返回零值，当值不是结构体或索引超界时发生宕机 |


下面代码构造一个结构体包含不同类型的成员。通过 reflect.Value 提供的成员访问函数，可以获得结构体值的各种数据。

反射访问结构体成员值：

```go
package main

import (
    "fmt"
    "reflect"
)

// 定义结构体
type dummy struct {
    a int
    b string
    // 嵌入字段
    float32
    bool
    next *dummy
}

func main() {
    
    // 值包装结构体
    d := reflect.ValueOf(dummy{
            next: &dummy{},
    })
    
    // 获取字段数量
    fmt.Println("NumField", d.NumField()) 
    
    // 获取索引为2的字段(float32字段)
    floatField := d.Field(2)
    // 输出字段类型
    fmt.Println("Field", floatField.Type())
    
    // 根据名字查找字段
    fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())
    
    // 根据索引查找值中, next字段的int字段的值
    fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
    //[]int{4,0} 中的 4 表示，在 dummy 结构中索引值为 4 的成员，也就是 next。next 的类型为 dummy，也是一个结构体，因此使用 []int{4,0} 中的 0 继续在 next 值的基础上索引，结构为 dummy 中索引值为 0 的 a 字段，类型为 int。

}

//代码输出如下：
NumField 5
Field float32
FieldByName("b").Type string
FieldByIndex([]int{4, 0}).Type() int
```

## 判断反射值的空和有效性

| 方 法          | 说 明                                                        |
| -------------- | ------------------------------------------------------------ |
| IsNil() bool   | 返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的`v== nil`操作 |
| IsValid() bool | 判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等。 |

下面的例子将会对各种方式的空指针进行 `IsNil() `和` IsValid()` 的返回值判定检测。同时对结构体成员及方法查找 map 键值对的返回值进行 IsValid() 判定，参考下面的代码。

反射值对象的零值和有效性判断：

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    
    // *int的空指针
    var a *int
    fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
    
    // nil值
    fmt.Println("nil:", reflect.ValueOf(nil).IsValid())
    
    // *int类型的空指针
    fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
    
    // 实例化一个结构体
    s := struct{}{}
    // 尝试从结构体中查找一个不存在的字段
    fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
    // 尝试从结构体中查找一个不存在的方法
    fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())
    
    // 实例化一个map
    m := map[int]int{}
    // 尝试从map中查找一个不存在的键
    fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}

//代码输出如下：
var a *int: true
nil: false
(*int)(nil): false
不存在的结构体成员: false
不存在的结构体方法: false
不存在的键： false
```

## 通过反射修改变量的值

### 概述

Go语言中类似 x、x.f[1] 和 *p 形式的表达式都可以表示变量，但是其它如 x + 1 和 f(2) 则不是变量。`一个变量就是一个可寻址的内存空间`，里面存储了一个值，并且存储的值可以通过内存地址来更新。

对于 reflect.Values 也有类似的区别。有一些 reflect.Values 是可取地址的；其它一些则不可以。考虑以下的声明语句：

```go
x := 2 // value type variable?

a := reflect.ValueOf(2) // 2 int no

b := reflect.ValueOf(x) // 2 int no

c := reflect.ValueOf(&x) // &x *int no

d := c.Elem() // 2 int yes (x)
```

其中 a 对应的变量则不可取地址。因为 a 中的值仅仅是整数 2 的拷贝副本。b 中的值也同样不可取地址。c 中的值还是不可取地址，它只是一个指针 &x 的拷贝。实际上，所有通过 reflect.ValueOf(x) 返回的 reflect.Value 都是不可取地址的。但是对于 d，它是 c 的解引用方式生成的，指向另一个变量，因此是可取地址的。我们可以通过调用 `reflect.ValueOf(&x).Elem()`，来获取任意变量x对应的可取地址的 Value。

我们可以通过调用 reflect.Value 的 `CanAddr `方法来判断其是否可以被取地址：

```go
fmt.Println(a.CanAddr()) // "false"
fmt.Println(b.CanAddr()) // "false"
fmt.Println(c.CanAddr()) // "false"
fmt.Println(d.CanAddr()) // "true"
```

每当我们通过指针间接地获取的 reflect.Value 都是可取地址的，即使开始的是一个不可取地址的 Value。在反射机制中，所有关于是否支持取地址的规则都是类似的。例如，slice 的索引表达式 e[i]将隐式地包含一个指针，它就是可取地址的，即使开始的e表达式不支持也没有关系。

以此类推，reflect.ValueOf(e).Index(i) 对于的值也是可取地址的，即使原始的 reflect.ValueOf(e) 不支持也没有关系。

使用 reflect.Value 对包装的值进行修改时，需要遵循一些规则。如果没有按照规则进行代码设计和编写，轻则无法修改对象值，重则程序在运行时会发生宕机。

### 判定及获取元素的相关方法

使用 reflect.Value 取元素、取地址及修改值的属性方法请参考下表。

| 方法名         | 备  注                                                       |
| -------------- | ------------------------------------------------------------ |
| Elem() Value   | 取值指向的元素值，类似于语言层`*`操作。当值类型不是指针或接口时发生宕 机，空指针时返回 nil 的 Value |
| Addr() Value   | 对可寻址的值返回其地址，类似于语言层`&`操作。当值不可寻址时发生宕机 |
| CanAddr() bool | 表示值是否可寻址                                             |
| CanSet() bool  | 返回值能否被修改。要求值可寻址且是导出的字段                 |

### 值修改相关方法

使用 reflect.Value 修改值的相关方法如下表所示。

| Set(x Value)        | 将值设置为传入的反射值对象的值                               |
| ------------------- | ------------------------------------------------------------ |
| Setlnt(x int64)     | 使用 int64 设置值。当值的类型不是 int、int8、int16、 int32、int64 时会发生宕机 |
| SetUint(x uint64)   | 使用 uint64 设置值。当值的类型不是 uint、uint8、uint16、uint32、uint64 时会发生宕机 |
| SetFloat(x float64) | 使用 float64 设置值。当值的类型不是 float32、float64 时会发生宕机 |
| SetBool(x bool)     | 使用 bool 设置值。当值的类型不是 bod 时会发生宕机            |
| SetBytes(x []byte)  | 设置字节数组 []bytes值。当值的类型不是 []byte 时会发生宕机   |
| SetString(x string) | 设置字符串值。当值的类型不是 string 时会发生宕机             |


以上方法，在 reflect.Value 的 CanSet 返回 false 仍然修改值时会发生宕机。

在已知值的类型时，应尽量使用值对应类型的反射设置值。



### 值可修改条件之一：可被寻址

**通过反射修改变量值的前提条件之一：这个值必须可以被寻址。**简单地说就是这个变量必须能被修改。示例代码如下：

```go
package main

import (
    "reflect"
)

func main() {
    
    // 声明整型变量a并赋初值
    var a int = 1024
    
    // 获取变量a的反射值对象
    valueOfA := reflect.ValueOf(a)
    
    // 尝试将a修改为1(此处会发生崩溃)
    valueOfA.SetInt(1)
}
```

程序运行崩溃，打印错误：

```go
panic: reflect: reflect.Value.SetInt using unaddressable value
```

报错意思是：SetInt 正在使用一个不能被寻址的值。从 reflect.ValueOf 传入的是 a 的值，而不是 a 的地址，这个 reflect.Value 当然是不能被寻址的。将代码修改一下，重新运行：

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {

    // 声明整型变量a并赋初值
    var a int = 1024
    
    // 获取变量a的反射值对象(a的地址)
    valueOfA := reflect.ValueOf(&a)
    
    // 取出a地址的元素(a的值)
    valueOfA = valueOfA.Elem()
    
    // 修改a的值为1
    valueOfA.SetInt(1)
    
    // 打印a的值
    fmt.Println(valueOfA.Int())
}
//代码输出如下：
1
```

### 值可修改条件之一：被导出

结构体成员中，如果字段没有被导出，即便不使用反射也可以被访问，但不能通过反射修改，代码如下：	

```go
package main

import (
    "reflect"
)

func main() {		
    type dog struct {
            legCount int
    }
    
    // 获取dog实例的反射值对象
    valueOfDog := reflect.ValueOf(dog{})
    
    // 获取legCount字段的值
    vLegCount := valueOfDog.FieldByName("legCount")
    
    // 尝试设置legCount的值(这里会发生崩溃)
    vLegCount.SetInt(4)
}

程序发生崩溃，报错：
panic: reflect: reflect.Value.SetInt using value obtained using unexported field
```

报错的意思是：SetInt() 使用的值来自于一个未导出的字段。

为了能修改这个值，需要将该字段导出。将 dog 中的 legCount 的成员首字母大写，导出 LegCount 让反射可以访问，修改后的代码如下：

```go
type dog struct {
    LegCount int
}
```

然后根据字段名获取字段的值时，将字符串的字段首字母大写，修改后的代码如下：

```go
vLegCount := valueOfDog.FieldByName("LegCount")
```

再次运行程序，发现仍然报错：

```go
panic: reflect: reflect.Value.SetInt using unaddressable value
```

 `valueOfDog` 这个结构体实例不能被寻址，因此其字段也不能被修改。修改代码，取结构体的指针，再通过 reflect.Value 的 Elem() 方法取到值的反射值对象。修改后的完整代码如下：

```go
package main

import (
    "reflect"
    "fmt"
)

func main() {

    type dog struct {
            LegCount int
    }
    
    // 获取dog实例地址的反射值对象
    valueOfDog := reflect.ValueOf(&dog{})
    
    // 取出dog实例地址的元素
    valueOfDog = valueOfDog.Elem()
    
    // 获取legCount字段的值
    vLegCount := valueOfDog.FieldByName("LegCount")
    
    // 尝试设置legCount的值
    vLegCount.SetInt(4)
    fmt.Println(vLegCount.Int())
}

//代码输出如下：
4
```

值的修改从表面意义上叫`可寻址`，换一种说法就是值必须“可被设置”。那么，想修改变量值，一般的步骤是：

1. 取这个变量的地址或者这个变量所在的结构体已经是指针类型。
2. 使用 reflect.ValueOf 进行值包装。
3. 通过 Value.Elem() 获得指针值指向的元素值对象（Value），因为值对象（Value）内部对象为指针时，使用 set 设置时会报出宕机错误。
4. 使用 Value.Set 设置值。

## 反射结构体

### 使用反射获取结构体的成员类型

任意值通过` reflect.TypeOf() `获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 `NumField() `和` Field() `方法获得结构体成员的详细信息。

与成员获取相关的 reflect.Type 的方法如下表所示。

| 方法                                                        | 说明                                                         |
| ----------------------------------------------------------- | ------------------------------------------------------------ |
| Field(i int) StructField                                    | 根据索引返回索引对应的结构体字段的信息，当值不是结构体或索引超界时发生宕机 |
| NumField() int                                              | 返回结构体成员字段数量，当类型不是结构体或索引超界时发生宕机 |
| FieldByName(name string) (StructField, bool)                | 根据给定字符串返回字符串对应的结构体字段的信息，没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机 |
| FieldByIndex(index []int) StructField                       | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生宕机 |
| FieldByNameFunc(match func(string) bool) (StructField,bool) | 根据匹配函数匹配需要的字段，当值不是结构体或索引超界时发生宕机 |

#### 结构体字段类型

reflect.Type 的` Field() `方法返回` StructField `结构，这个结构描述结构体的成员信息，通过这个信息可以获取成员与结构体的关系，如偏移、索引、是否为匿名字段、结构体标签（StructTag）等，而且还可以通过 StructField 的 Type 字段进一步获取结构体成员的类型信息。

StructField 的结构如下：

```go
type StructField struct {
    Name string          // 字段名
    PkgPath string       // 字段路径
    Type      Type       // 字段反射类型对象
    Tag       StructTag  // 字段的结构体标签
    Offset    uintptr    // 字段在结构体中的相对偏移
    Index     []int      // Type.FieldByIndex中的返回的索引值
    Anonymous bool       // 是否为匿名字段
}
```

字段说明如下：

- Name：为字段名称。
- PkgPath：字段在结构体中的路径。
- Type：字段本身的反射类型对象，类型为 reflect.Type，可以进一步获取字段的类型信息。
- Tag：结构体标签，为结构体字段标签的额外信息，可以单独提取。
- Index：FieldByIndex 中的索引顺序。
- Anonymous：表示该字段是否为匿名字段。

#### 获取成员反射信息

下面代码中，实例化一个结构体并遍历其结构体成员，再通过 reflect.Type 的 FieldByName() 方法查找结构体中指定名称的字段，直接获取其类型信息。

反射访问结构体成员类型及信息：

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    
    // 声明一个空结构体
    type cat struct {
        
        Name string
        
        // 带有结构体tag的字段
        //Type 是 cat 的一个成员，这个成员类型后面带有一个以 ` 开始和结尾的字符串。
        //这个字符串在Go语言中被称为 Tag（标签）。
        //一般用于给字段添加自定义信息，方便其他模块根据信息进行不同功能的处理。
        Type int `json:"type" id:"100"`
    }
    
    // 创建 cat 实例，并对两个字段赋值。结构体标签属于类型信息，无须且不能赋值。
    ins := cat{Name: "mimi", Type: 1}
    
    // 获取结构体实例的反射类型对象
    typeOfCat := reflect.TypeOf(ins)
    
    // 遍历结构体所有成员
    //使用 reflect.Type 类型的 NumField() 方法获得一个结构体类型共有多少个字段。
    //如果类型不是结构体，将会触发宕机错误。
    for i := 0; i < typeOfCat.NumField(); i++ {
        
        // 获取每个成员的结构体字段类型
        // Field() 方法和 NumField 一般都是配对使用，用来实现结构体成员的遍历操作。
        fieldType := typeOfCat.Field(i)
        
        // 输出成员名和tag
        fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
    }
    
    // 通过字段名, 找到字段类型信息
    if catType, ok := typeOfCat.FieldByName("Type"); ok {
        
        //使用 StructField 中 Tag 的 Get() 方法，根据 Tag 中的名字进行信息获取。
        fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
    }
}

//代码输出如下：
name: Name  tag: ''
name: Type  tag: 'json:"type" id:"100"'
type 100
```

### 结构体标签（Struct Tag）

通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的` Tag `被称为结构体标签（StructTag）。结构体标签是对结构体字段的额外信息标签。

JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称 `ORM`）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。这些信息都是静态的，无须实例化结构体，可以通过反射获取到。

#### 结构体标签的格式

Tag 在结构体字段后方书写的格式如下：

```go
`key1:"value1" key2:"value2"`
```

结构体标签由一个或多个键值对组成；键与值使用冒号分隔，值用双引号括起来；键值对之间使用`一个空格分隔。`

#### 从结构体标签中获取值

StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：

- `func (tag StructTag) Get(key string) string`：根据 Tag 中的键获取对应的值，例如``key1:"value1" key2:"value2"``的 Tag 中，可以传入“key1”获得“value1”。

- `func (tag StructTag) Lookup(key string) (value string, ok bool)`：根据 Tag 中的键，查询值是否存在。

#### 结构体标签格式错误的问题

编写 Tag 时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，示例代码如下：

```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    
    type cat struct {
        Name string
        Type int `json: "type"  id:"100"`
    }
    
    typeOfCat := reflect.TypeOf(cat{})
    
    if catType, ok := typeOfCat.FieldByName("Type"); ok {
        fmt.Println(catType.Tag.Get("json"))
    }
}
//运行上面的代码会输出一个空字符串，并不会输出期望的 type。
```

在 json: 和 "type" 之间增加了一个空格，这种写法没有遵守结构体标签的规则，因此无法通过 Tag.Get 获取到正确的 json 对应的值。这个错误在开发中非常容易被疏忽，造成难以察觉的错误。

```go
type cat struct {
    Name string
    Type int `json:"type" id:"100"`
}
//运行结果如下：
type
```



### 结构体字段信息的获取修改

我们一般使用反射修改结构体的字段，只要有结构体的指针，我们就可以修改它的字段。

下面是一个解析结构体变量 t 的例子，用结构体的地址创建反射变量，再修改它。然后我们对它的类型设置了 typeOfT，并用调用简单的方法迭代字段。

需要注意的是，我们从结构体的类型中提取了字段的名字，但每个字段本身是正常的` reflect.Value `对象。

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    
    type T struct {
        A int
        B string
    }
    
    t := T{23, "skidoo"}
    
    s := reflect.ValueOf(&t).Elem()
    
    typeOfT := s.Type()
    
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v\n", i,
            typeOfT.Field(i).Name, f.Type(), f.Interface())
    }
}

//运行结果如下：
0: A int = 23
1: B string = skidoo
```

T 字段名之所以大写，是因为结构体中只有可导出的字段是“可设置”的。

因为 s 包含了一个可设置的反射对象，我们可以修改结构体字段：

```go
package main
import (
    "fmt"
    "reflect"
)

func main() {
    type T struct {
        A int
        B string
    }
    
    t := T{23, "skidoo"}
    
    s := reflect.ValueOf(&t).Elem()
    
    s.Field(0).SetInt(77)
    s.Field(1).SetString("Sunset Strip")
    
    fmt.Println("t is now", t)
}

//运行结果如下：
t is now {77 Sunset Strip}
```

如果我们修改了程序让 s 由 t（而不是 &t）创建，程序就会在调用 SetInt 和 SetString 的地方失败，因为 t 的字段是不可设置的。

## 通过类型信息创建实例



，代码如下：

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {

    var a int
    // 取变量a的反射类型对象
    typeOfA := reflect.TypeOf(a)
    
    // 根据反射类型对象创建类型实例
    aIns := reflect.New(typeOfA)
    
    // 输出Value的类型和种类
    fmt.Println(aIns.Type(), aIns.Kind()) //代码输出如下：*int ptr
}
```

使用` reflect.New()` 函数传入变量 a 的反射类型对象，创建这个类型的实例值，值以 reflect.Value 类型返回。这步操作等效于：new(int)，因此返回的是` *int `类型的实例。

## 通过反射调用函数

如果反射值对象（reflect.Value）中值的类型为`函数`时，可以通过 reflect.Value 调用该函数。使用反射调用函数时，需要将参数使用反射值对象的切片` []reflect.Value` 构造后传入` Call() `方法中，调用完成时，函数的返回值通过` []reflect.Value` 返回。

下面的代码声明一个加法函数，传入两个整型值，返回两个整型值的和。将函数保存到反射值对象（reflect.Value）中，然后将两个整型值构造为反射值对象的切片（[]reflect.Value），使用 Call() 方法进行调用。如果反射值对象（reflect.Value）中值的类型为`函数`时，可以通过 reflect.Value 调用该函数。使用反射调用函数时，需要将参数使用反射值对象的切片` []reflect.Value` 构造后传入` Call() `方法中，调用完成时，函数的返回值通过` []reflect.Value` 返回。

下面的代码声明一个加法函数，传入两个整型值，返回两个整型值的和。将函数保存到反射值对象（reflect.Value）中，然后将两个整型值构造为反射值对象的切片（[]reflect.Value），使用 Call() 方法进行调用。

反射调用函数：

```go
package main

import (
    "fmt"
    "reflect"
)

// 普通函数
func add(a, b int) int {
    return a + b
}

func main() {

    // 将函数包装为反射值对象
    funcValue := reflect.ValueOf(add)
    
    // 构造函数参数, 传入两个整型值
    paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
    
    // 反射调用函数
    retList := funcValue.Call(paramList)
    
    // 获取第一个返回值, 取整数值
    fmt.Println(retList[0].Int())
}
```

>  提示
>
> 反射调用函数的过程需要构造大量的 reflect.Value 和中间变量，对函数参数值进行逐一检查，还需要将调用参数复制到调用函数的参数内存中。调用完毕后，还需要将返回值转换为 reflect.Value，用户还需要从中取出调用值。因此，反射调用函数的性能问题尤为突出，不建议大量使用反射函数调用。

## 通过反射调用方法

整体与调用函数一致，额外的需要先通过对象的值反射获取对象方法的反射对象，再使用 Call() 调用，示例：

```go
type Stu struct {
  Name string
}

func (this *Stu) Fn(p1, p2 int) int {
  return p1 + p2
}

func main() {
  s := &Stu{"Hank"}
  valueS := reflect.ValueOf(s)
  method := valueS.MethodByName("Fn")
  paramList := []reflect.Value{
    reflect.ValueOf(22),
    reflect.ValueOf(20),
  }
  resultList := method.Call(paramList)
  fmt.Println(resultList[0].Int()) // 42
}
```



## inject库：依赖注入

**“依赖注入”和“控制反转”**

正常情况下，对函数或方法的调用是我们的主动直接行为，在调用某个函数之前我们需要清楚地知道被调函数的名称是什么，参数有哪些类型等等。

所谓的`控制反转`就是将这种主动行为变成间接的行为，我们不用直接调用函数或对象，而是借助框架代码进行间接的调用和初始化，这种行为称作“控制反转”，库和框架能很好的解释控制反转的概念。

`依赖注入`是实现控制反转的一种方法，如果说控制反转是一种设计思想，那么依赖注入就是这种思想的一种实现，通过注入参数或实例的方式实现控制反转。如果没有特殊说明，我们可以认为依赖注入和控制反转是一个东西。

控制反转的价值在于解耦，有了控制反转就不需要将代码写死，可以让控制反转的的框架代码读取配置，动态的构建对象

### inject 实践

`inject` 是依赖注入的Go语言实现，它能在运行时注入参数，调用方法，是 Martini 框架（Go语言中著名的 Web 框架）的基础核心。

在介绍具体实现之前，先来想一个问题，如何通过一个字符串类型的函数名来调用函数？Go语言没有 Java 中的 Class.forName 方法可以通过类名直接构造对象，所以这种方法是行不通的，能想到的方法就是使用 map 实现一个字符串到函数的映射，示例代码如下：

```go
func fl() {
    println ("fl")
}

func f2 () {
    println ("f2")
}

funcs := make(map[string] func ())
funcs ["fl"] = fl
funcs ["f2"] = fl
funcs ["fl"]()
funcs ["f2"]()
```

但是这有个缺陷，就是 map 的 Value 类型被写成 func()，不同参数和返回值的类型的函数并不能通用。将 map 的 Value 定义为` interface{} `空接口类型即可以解决该问题，但需要借助类型断言或反射来实现，通过类型断言实现等于又绕回去了，反射是一种可行的办法。

inject 包借助反射实现函数的注入调用，下面通过一个示例来看一下。

```go
package main

import (
    "fmt"
    "github.com/codegangsta/inject"
)

type S1 interface{}
type S2 interface{}

func Format(name string, company S1, level S2, age int) {
    fmt.Printf("name ＝ %s, company=%s, level=%s, age ＝ %d!\n", name, company, level, age)
}

func main() {
    
    //控制实例的创建
    inj := inject.New()
    
    //实参注入
    inj.Map("tom")
    inj.MapTo("tencent", (*S1)(nil))
    inj.MapTo("T4", (*S2)(nil))
    inj.Map(23)
    
    //函数反转调用
    inj.Invoke(Format)
}

//运行结果如下：
name ＝ tom, company=tencent, level=T4, age ＝ 23!
```

可见 inject 提供了一种注入参数调用函数的通用功能，`inject.New()` 相当于创建了一个控制实例，由其来实现对函数的注入调用。inject 包不但提供了对函数的注入，还实现了对 struct 类型的注入，示例代码如下所示：

```go
package main

import (
    "fmt"
    "github.com/codegangsta/inject"
)

type S1 interface{}
type S2 interface{}

type Staff struct {
    Name    string `inject`
    Company S1     `inject`
    Level   S2     `inject`
    Age     int    `inject`
}

func main() {

    //创建被注入实例
    s := Staff{}
    
    //控制实例的创建
    inj := inject.New()
    
    //初始化注入值
    inj.Map("tom")
    inj.MapTo("tencent", (*S1)(nil))
    inj.MapTo("T4", (*S2)(nil))
    inj.Map(23)
    
    //实现对 struct 注入
    inj.Apply(&s)
    
    //打印结果
    fmt.Printf("s ＝ %v\n", s)
}

//运行结果如下：
s ＝ {tom tencent T4 23}
```

可以看到 inject 提供了一种对结构类型的通用注入方法。至此，我们仅仅从宏观层面了解 iniect 能做什么，下面从源码实现角度来分析 inject。



### inject 原理分析

inject 包中只有 2 个文件，一个是 inject.go 文件和一个 inject_test.go 文件，这里我们只需要关注 inject.go 文件即可。

inject.go 短小精悍，包括注释和空行在内才 157 行代码，代码中定义了 4 个接口，包括一个父接口和三个子接口，如下所示：

```go
type Injector interface {
    Applicator
    Invoker
    TypeMapper
    SetParent(Injector)
}

type Applicator interface {
    Apply(interface{}) error
}

type Invoker interface {
    Invoke(interface{}) ([]reflect.Value, error)
}

type TypeMapper interface {
    Map(interface{}) TypeMapper
    MapTo(interface{}, interface{}) TypeMapper
    Get(reflect.Type) reflect.Value
}
```

Injector 接口是 Applicator、Invoker、TypeMapper 接口的父接口，所以实现了 Injector 接口的类型，也必然实现了 Applicator、Invoker 和 TypeMapper 接口：

- Applicator 接口只规定了 Apply 成员，它用于注入 struct。
- Invoker 接口只规定了 Invoke 成员，它用于执行被调用者。
- TypeMapper 接口规定了三个成员，Map 和 MapTo 都用于注入参数，但它们有不同的用法，Get 用于调用时获取被注入的参数。


另外 Injector 还规定了 SetParent 行为，它用于设置父 Injector，其实它相当于查找继承。也即通过 Get 方法在获取被注入参数时会一直追溯到 parent，这是个递归过程，直到查找到参数或为 `nil `终止。

```go
type injector struct {
    values map[reflect.Type]reflect.Value
    parent Injector
}

func InterfaceOf(value interface{}) reflect.Type {
    t := reflect.TypeOf(value)
    for t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    if t.Kind() != reflect.Interface {
        panic("Called inject.InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
    }
    return t
}

func New() Injector {
    return &injector{
        values: make(map[reflect.Type]reflect.Value),
    }
}
```

injector 是 inject 包中唯一定义的 struct，所有的操作都是基于 injector struct 来进行的，它有两个成员 values 和 parent。values 用于保存注入的参数，是一个用 reflect.Type 当键、reflect.Value 为值的 map，理解这点将有助于理解 Map 和 MapTo。

New 方法用于初始化 injector struct，并返回一个指向 injector struct 的指针，但是这个返回值被 Injector 接口包装了。

InterfaceOf 方法虽然只有几句实现代码，但它是 Injector 的核心。InterfaceOf 方法的参数必须是一个接口类型的指针，如果不是则引发 panic。InterfaceOf 方法的返回类型是 reflect.Type，大家应该还记得 injector 的成员 values 就是一个 reflect.Type 类型当键的 map。这个方法的作用其实只是获取参数的类型，而不关心它的值。

示例代码如下所示：

```go
package main

import (
    "fmt"
    "github.com/codegangsta/inject"
)

type SpecialString interface{}

func main() {
    fmt.Println(inject.InterfaceOf((*interface{})(nil)))
    fmt.Println(inject.InterfaceOf((*SpecialString)(nil)))
}

//运行结果如下：
interface {}
main.SpecialString
```

InterfaceOf 方法就是用来得到参数类型，而不关心它具体存储的是什么值。

```go
func (i *injector) Map(val interface{}) TypeMapper {
    i.values[reflect.TypeOf(val)] = reflect.ValueOf(val)
    return i
}

func (i *injector) MapTo(val interface{}, ifacePtr interface{}) TypeMapper {
    i.values[InterfaceOf(ifacePtr)] = reflect.ValueOf(val)
    return i
}

func (i *injector) Get(t reflect.Type) reflect.Value {
    val := i.values[t]
    if !val.IsValid() && i.parent != nil {
        val = i.parent.Get(t)
    }
    return val
}

func (i *injector) SetParent(parent Injector) {
    i.parent = parent
}
```

Map 和 MapTo 方法都用于注入参数，保存于 injector 的成员 values 中。这两个方法的功能完全相同，唯一的区别就是 Map 方法用参数值本身的类型当键，而 MapTo 方法有一个额外的参数可以指定特定的类型当键。但是 MapTo 方法的第二个参数 ifacePtr 必须是接口指针类型，因为最终 ifacePtr 会作为 InterfaceOf 方法的参数。

为什么需要有 MapTo 方法？因为注入的参数是存储在一个以类型为键的 map 中，可想而知，当一个函数中有一个以上的参数的类型是一样时，后执行 Map 进行注入的参数将会覆盖前一个通过 Map 注入的参数。

SetParent 方法用于给某个 Injector 指定父 Injector。Get 方法通过 reflect.Type 从 injector 的 values 成员中取出对应的值，它可能会检查是否设置了 parent，直到找到或返回无效的值，最后 Get 方法的返回值会经过 IsValid 方法的校验。

```go
package main

import (
    "fmt"
    "reflect"
    "github.com/codegangsta/inject"
)

type SpecialString interface{}

func main() {

    inj := inject.New()
    
    inj.Map("C语言中文网")
    
    inj.MapTo("Golang", (*SpecialString)(nil))
    inj.Map(20)
    
    fmt.Println("字符串是否有效？", inj.Get(reflect.TypeOf("Go语言入门教程")).IsValid())
    
    fmt.Println("特殊字符串是否有效？", inj.Get(inject.InterfaceOf((*SpecialString)(nil))).IsValid())
    
    fmt.Println("int 是否有效？", inj.Get(reflect.TypeOf(18)).IsValid())
    
    fmt.Println("[]byte 是否有效？", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
    
    
    inj2 := inject.New()
    inj2.Map([]byte("test"))
    inj.SetParent(inj2)
    fmt.Println("[]byte 是否有效？", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
}

//运行结果如下所示：
字符串是否有效？ true
特殊字符串是否有效？ true
int 是否有效？ true
[]byte 是否有效？ false
[]byte 是否有效？ true
```

通过以上例子应该知道 SetParent 是什么样的行为，是不是很像面向对象中的查找链？

```go
func (inj *injector) Invoke(f interface{}) ([]reflect.Value, error) {
    t := reflect.TypeOf(f)
    var in = make([]reflect.Value, t.NumIn()) //Panic if t is not kind of Func
    for i := 0; i < t.NumIn(); i++ {
        argType := t.In(i)
        val := inj.Get(argType)
        if !val.IsValid() {
            return nil, fmt.Errorf("Value not found for type %v", argType)
        }
        in[i] = val
    }
    return reflect.ValueOf(f).Call(in), nil
}
```

Invoke 方法用于动态执行函数，当然执行前可以通过 Map 或 MapTo 来注入参数，因为通过 Invoke 执行的函数会取出已注入的参数，然后通过 reflect 包中的 Call 方法来调用。Invoke 接收的参数 f 是一个接口类型，但是 f 的底层类型必须为 func，否则会 panic。

```go
package main
import (
    "fmt"
    "github.com/codegangsta/inject"
)
type SpecialString interface{}

func Say(name string, gender SpecialString, age int) {
    fmt.Printf("My name is %s, gender is %s, age is %d!\n", name, gender, age)
}

func main() {
    inj := inject.New()
    inj.Map("张三")
    inj.MapTo("男", (*SpecialString)(nil))
    
    inj2 := inject.New()
    inj2.Map(25)
    inj.SetParent(inj2)
    
    inj.Invoke(Say)
}

//运行结果如下：
My name is 张三, gender is 男, age is 25!
```

上面的例子如果没有定义 SpecialString 接口作为 gender 参数的类型，而把 name 和 gender 都定义为 string 类型，那么 gender 会覆盖 name 的值。

```go
func (inj *injector) Apply(val interface{}) error {
    v := reflect.ValueOf(val)
    for v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    if v.Kind() != reflect.Struct {
        return nil
    }
    t := v.Type()
    for i := 0; i < v.NumField(); i++ {
        f := v.Field(i)
        structField := t.Field(i)
        if f.CanSet() && structField.Tag == "inject" {
            ft := f.Type()
            v := inj.Get(ft)
            if !v.IsValid() {
                return fmt.Errorf("Value not found for type %v", ft)
            }
            f.Set(v)
        }
    }
    return nil
}
```

Apply 方法是用于对 struct 的字段进行注入，参数为指向底层类型为结构体的指针。可注入的前提是：字段必须是导出的（也即字段名以大写字母开头），并且此字段的 tag 设置为``inject``。

示例代码如下所示：

```go
package main
import (
    "fmt"
    "github.com/codegangsta/inject"
)
type SpecialString interface{}
type TestStruct struct {
    Name   string `inject`
    Nick   []byte
    Gender SpecialString `inject`
    uid    int           `inject`
    Age    int           `inject`
}
func main() {
    s := TestStruct{}
    inj := inject.New()
    inj.Map("张三")
    inj.MapTo("男", (*SpecialString)(nil))
    inj2 := inject.New()
    inj2.Map(26)
    inj.SetParent(inj2)
    inj.Apply(&s)
    fmt.Println("s.Name =", s.Name)
    fmt.Println("s.Gender =", s.Gender)
    fmt.Println("s.Age =", s.Age)
}

//运行结果如下：
s.Name = 张三
s.Gender = 男
s.Age = 26
```

