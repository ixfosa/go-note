package main

import (
	"fmt"
	"sync"
)

// 单例模式

// 使用懒惰模式的单例模式，使用双重检查加锁保证线程安全

// Singleton 是单例模式类
type Singleton struct {

}

var singleton *Singleton
var once sync.Once

// GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
func main() {
	ins1  := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		fmt.Println("instance is not equal")
	} else {
		fmt.Println("instance is equal")
	}


	ParallelSingleton()
}

const parCount = 100
func ParallelSingleton()  {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*Singleton{}

	for i := 0; i < parCount; i++ {
		go func(idx int) {
			instances[idx] = GetInstance()
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			fmt.Println("instance is not equal")
		} else {
			fmt.Println("instance is equal")
		}
 	}
}