package main

import (
	"fmt"
	"math"
	"strings"
)

/*----------------基本类型介绍--------------------*/
/*
Golang 更明确的数字类型命名，支持 Unicode，支持常用数据结构。
	类型				长度(字节)	默认值	说明
	bool			1			false
	byte			1	0		uint8
	rune			4	0		Unicode Code Point, int32
	int, uint		4或8		0	32 或 64 位
	int8, uint8		1	0		-128 ~ 127, 0 ~ 255，byte是uint8 的别名
	int16, uint16	2	0		-32768 ~ 32767, 0 ~ 65535
	int32, uint32	4	0		-21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名
	int64, uint64	8	0
	float32	4		0.0
	float64	8		0.0
	complex64		8
	complex128		16
	uintptr			4或8		以存储指针的 uint32 或 uint64 整数
	array			值类型
	struct			值类型
	string			""			UTF-8 字符串
	slice			nil			引用类型
	map				nil			引用类型
	channel			nil			引用类型
	interface		nil			接口
	function		nil			函数

	支持八进制、 六进制，以及科学记数法。标准库 math 定义了各数字类型取值范围。
		 a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
	空指针值 nil，而非C/C++ NULL。
 */


/*---------------整型---------------------*/
/*
	整型分为以下两个大类： 按长度分为：int8、int16、int32、int64对应的无符号整型：uint8、uint16、uint32、uint64
	其中，uint8就是我们熟知的 byte 型，int16对应C语言中的 short 型，int64对应C语言中的 long 型。
 */


/*-----------------浮点型-------------------*/
/*
	Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准：
	float32 的浮点数的最大范围约为3.4e38，可以使用常量定义：math.MaxFloat32。 float64 的浮点数的最大范围约为 1.8e308，
	可以使用一个常量定义：math.MaxFloat64。
 */
func main81() {
	fmt.Println("math.MaxFloat32 = ", math.MaxFloat32)  //math.MaxFloat32 =  3.4028234663852886e+38
	fmt.Println("math.MaxFloat64 = ", math.MaxFloat64)  //math.MaxFloat64 =  1.7976931348623157e+308
}

/*----------------复数--------------------*/
/*
complex64和complex128
复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
 */
func main82() {
	var comp complex64 = 1 + 2i
	fmt.Println("comp = ", comp) //comp =  (1+2i)
}


/*-----------------布尔值-------------------*/
/*
Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

注意：
    布尔类型变量的默认值为false。
    Go 语言中不允许将整型强制转换为布尔型.
    布尔型无法参与数值运算，也无法与其他类型进行转换。
 */
func main83() {
	var a = true
	fmt.Println("a = ", a)  //a =  true
	//var b = 1 + a  //invalid operation: 1 + a (mismatched types int and bool)
	//fmt.Println("b = ", b)
}


/*-----------------字符串-------------------*/
/*
Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 
Go 语言里的字符串的内部实现使用 UTF-8 编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：
 */
func main84()  {
	name := "ixfosa"
	sex := "man"
	println("name = ", name, "sex = ", sex)
}

/*------------------- 字符串转义符-----------------*/
/*
Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。
	转义	含义
	\r	回车符（返回行首）
	\n	换行符（直接跳到下一行的同列位置）
	\t	制表符
	\'	单引号
	\"	双引号
	\	反斜杠

 */
func main85() {
	fmt.Println(" \"c:\\pprof\\main.exe\"") //"c:\pprof\main.exe"
}

/*-----------------多行字符串-------------------*/
/*
	Go语言中要定义一个多行字符串时，就必须使用反引号字符：
 */
func main86() {
	str := ` row1
	row2
	row3
`
	//反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
fmt.Println(str)
}

/*----------------字符串的常用操作--------------------*/
/*
	方法					介绍
	len(str)			求长度
	+或fmt.Sprintf		拼接字符串
	strings.Split		分割
	strings.Contains	判断是否包含
	strings.HasPrefix,strings.HasSuffix		前缀/后缀判断
	strings.Index(),strings.LastIndex()		子串出现的位置
	strings.Join(a[]string, sep string)		join操作
 */
func main87()  {
	hello := "hello"
	name := "ixfosa"
	str := "fo"
	imgName := "fo.png"

	fmt.Println("len(str) ", len(str))

	fmt.Println("strings.Split ", strings.Split(name, ""))
	fmt.Println("strings.Contains ", strings.Contains(name, str))

	fmt.Println("strings.HasPrefix ", strings.HasPrefix(imgName, "png"))
	fmt.Println("strings.HasSuffix ", strings.HasSuffix(imgName, "png"))

	fmt.Println("strings.Index() ", strings.Index(name, str))
	fmt.Println("strings.LastIndex() ", strings.LastIndex(name, str))

	helloIxfosa := fmt.Sprintf("%s world", hello)
	fmt.Println("helloIxfosa = ", helloIxfosa)

	/*len(str)  2
	strings.Split  [i x f o s a]
	strings.Contains  true
	strings.HasPrefix  false
	strings.HasSuffix  true
	strings.Index()  2
	strings.LastIndex()  2
	helloIxfosa =  hello world*/
}

/*----------------byte和rune类型--------------------*/
/*
组成每个字符串的元素叫做 “字符” ，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
	var a := '中'
	var b := 'x'

Go 语言的字符有以下两种：
    uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
    rune类型，代表一个 UTF-8字符。

当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。
rune类型实际是一个int32。 Go 使用了特殊的 rune 类型来处理 Unicode，
让基于 Unicode的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾
 */
func main88() {
	str := "ixfosa-夏佛送"

	//遍历字符串
	for i := 0; i < len(str); i++ {  //byte
		fmt.Printf("%v(%c)  ", str[i], str[i])
	}
	fmt.Println()

	for idx, v := range str {  //byte
		fmt.Printf("%v-%v(%c)  ", idx, v, v)
	}

	//105(i)  120(x)  102(f)  111(o)  115(s)  97(a)  45(-)  229(å)  164(¤)  143()  228(ä)  189(½)  155()  233(é)  128()  129()
	//0-105(i)  1-120(x)  2-102(f)  3-111(o)  4-115(s)  5-97(a)  6-45(-)  7-22799(夏)  10-20315(佛)  13-36865(送)

	/*因为UTF8编码下一个中文汉字由3~4个字节组成，
	所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

	字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
	字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
	rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。*/
}


/*-----------------修改字符串-------------------*/
/*
	要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
 */
func main89() {
	str := "ixfosa-夏佛送"
	var runeArr []rune = []rune(str)
	for k, v := range runeArr {
		fmt.Print(k , "-", string(v), "  ") //0-i  1-x  2-f  3-o  4-s  5-a  6--  7-夏  8-佛  9-送
	}

	fmt.Println()

	s1 := "hello"
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println("byteS1 = ", string(byteS1))

	runeStr := []rune(runeArr)
	runeStr[7] = ' '
	fmt.Println("runeStr = ", string(runeStr))

	//byteS1 =  Hello
	//runeStr =  ixfosa- 佛送
}

/*------------------类型转换------------------*/
/*
Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

强制类型转换的基本语法如下：
    T(表达式)
其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

比如计算直角三角形的斜边长时使用math包的Sqrt()函数，
该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
 */
func main() {
	a, b := 8, 8

	c := math.Sqrt(float64(a*b))

	fmt.Println(c) //8
}