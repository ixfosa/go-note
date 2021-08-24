package main

import "fmt"

// 下面的代码有什么问题？
type X struct {
}

func (x *X) test()  {
	fmt.Println(x)
}


func main() {
	var x *X
	x.test()

	// X{} 是不可寻址的，不能直接调用方法。
	// X{}.test() // cannot take the address of X literal

	// 在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。
	xx := X{}
	xx.test()
}
