package main

import "fmt"

// 外观模式
// API 为 facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。

//facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。

// API is facade interface of facade package
type API interface {
	Test() string
}

type AModuleAPI interface {
	TestA() string
}

type BModuleAPI interface {
	TestB() string
}

// facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}


type aModuleImpl struct {

}

type bModuleImpl struct {

}


func (*aModuleImpl) TestA() string  {
	return "A module running"
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// ewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}


func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAPI() API {
	return &apiImpl{
		a: &aModuleImpl{},
		b: &bModuleImpl{},
	}
}




func main() {
	api := NewAPI()
	ret  := api.Test()
	fmt.Println(ret)
}
