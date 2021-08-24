package main

import "fmt"

// 代理模式
// 代理模式用于延迟处理操作或者在进行实际操作前后进行其它处理。

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string
	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre: "

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += " :after"

	return res
}

func main() {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	fmt.Println(res)
}

