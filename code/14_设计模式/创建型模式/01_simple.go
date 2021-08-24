package main

import "fmt"

//
// go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
// NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。


//在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。

type API interface {
	Say(name string) string
}

// //NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &HiLong{}
	} else if t == 2 {
		return &HiZhong{}
	}
	return nil
}

type HiLong struct {

}

func (long *HiLong) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

type HiZhong struct {

}

func (zhong *HiZhong) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

func main() {
	api1 := NewAPI(1)
	str1 := api1.Say("long")
	fmt.Printf("str1: %v, type: %T\n", str1, api1) // str1: Hi long, type: *main.HiLong

	api2 := NewAPI(2)
	str2 := api1.Say("zhong")
	fmt.Printf("str2: %v, type: %T\n", str2, api2) // str2: Hi zhong, type: *main.HiZhong

}
