package main

import "fmt"

// 适配器模式
//适配器模式用于转换一种接口适配另一种接口。

//实际使用中Adaptee一般为接口，并且使用工厂函数生成实例。

//在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，又因为Go语言中非入侵式接口特征，
// 其实Adapter也适配Adaptee接口。


//Target 是适配的目标接口
type Target interface {
	Request() string
}

// SpecificRequest 是目标类的一个方法
type Adaptee interface {
	SpecificRequest() string
}

// AdapteeImpl 是被适配的目标类
type AdapteeImpl struct {}

//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}


func (*AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}

// NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}


//NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}


func main() {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	fmt.Println(res)
}
