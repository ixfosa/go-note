package main

import "fmt"

/*
关于类型转化，下面选项正确的是？
	A.
		type MyInt int
		var i int = 1
		var j MyInt = i

	B.
		type MyInt int
		var i int = 1
		var j MyInt = (MyInt)i

	C.
		type MyInt int
		var i int = 1
		var j MyInt = MyInt(i)

	D.
		type MyInt int
		var i int = 1
		var j MyInt = i.(MyInt)
参考答案及解析：C。知识点：强制类型转化。



2.关于switch语句，下面说法正确的有?
	A. 条件表达式必须为常量或者整数；
	B. 单个case中，可以出现多个结果选项；
	C. 需要用break来明确退出一个case；
	D. 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；
参考答案及解析：BD。
 */

// 如果 Add() 函数的调用代码为：

func main() {
	var a Integer = 1
	var b Integer = 2

	var i interface{} = &a

	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
}

type Integer int

func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

// 参考答案及解析：AC。

/*
	A.
		type Integer int
		func (a Integer) Add(b Integer) Integer {
				return a + b
		}

	B.
		type Integer int
		func (a Integer) Add(b *Integer) Integer {
				return a + *b
		}

	C.
		type Integer int
		func (a *Integer) Add(b Integer) Integer {
				return *a + b
		}

	D.
		type Integer int
		func (a *Integer) Add(b *Integer) Integer {
				return *a + *b
		}
 */