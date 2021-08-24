package main

import "fmt"

// 工厂方法模式
// 工厂方法模式使用子类的方式延迟生成对象到子类中实现。
// Go中不存在继承 所以使用匿名组合来实现


// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

func (ob *OperatorBase) SetA (a int)  {
	ob.a = a
}

func (ob *OperatorBase) SetB (b int)  {
	ob.b = b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct {

}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (po *PlusOperator) result() int {
	return po.a + po.b
}

type MinusOperatorFactory struct {

}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator {
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperator  struct {
	*OperatorBase
}

func (po *MinusOperator ) result() int {
	return po.a - po.b
}

func main()  {
	//factory := PlusOperatorFactory{}
	factory := MinusOperatorFactory{}
	op := factory.Create()
	op.SetA(2)
	op.SetB(3)
	res := op.result()
	fmt.Println(res)
}