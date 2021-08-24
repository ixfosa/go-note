package main

import (
	"fmt"
)

// 创建者模式
// 一个复杂对象的构建分离成多个简单对象的构建组合


// Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct Product
func (d *Director) Construct ()  {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Builder1 struct {
	result string
}

func (b1 *Builder1) Part1 ()  {
	b1.result += "1"
}

func (b1 *Builder1) Part2 ()  {
	b1.result += "2"
}

func (b1 *Builder1) Part3 ()  {
	b1.result += "3"
}

func (b1 *Builder1) GetRes() string{
	return b1.result
}

type Builder2 struct {
	result int
}

func (b2 *Builder2) Part1 ()  {
	b2.result += 1
}

func (b2 *Builder2) Part2 ()  {
	b2.result += 2
}

func (b2 *Builder2) Part3 ()  {
	b2.result += 1
}

func (b2 *Builder2) GetRes() int{
	return b2.result
}

func main() {
	builder1 := &Builder1{}
	director1 := NewDirector(builder1)
	director1.Construct()
	res1:= builder1.GetRes()
	fmt.Println("res1: ", res1)


	builder2 := &Builder2{}
	director2 := NewDirector(builder2)
	director2.Construct()
	res2:= builder2.GetRes()
	fmt.Println("res2: ", res2)
}