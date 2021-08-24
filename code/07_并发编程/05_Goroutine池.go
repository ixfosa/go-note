package main

import (
	"fmt"
	"math/rand"
)

/****************** Goroutine池*********************/
/*
	本质上是生产者消费者模型
	可以有效控制goroutine数量，防止暴涨

需求：
	计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
	随机生成数字进行计算
 */

type Job struct {
	Id int
	RandNum int  // 需要计算的随机数
}

type Result struct {
	job *Job   // 这里必须传对象实例
	sum int   // 求和
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result)  {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// // 随机数接过来
				r_num :=  job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for sum != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
					r := &Result{
						job: job,
						sum: sum,
					}
					resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

func main() {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)

	// 2.结果管道
	resultChan := make(chan *Result, 128)

	// 3.创建工作池
	createPool(64, jobChan, resultChan)

	// 4.开个打印的协程
	go func( resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)

	// 循环创建job，输入到管道
	var id int = 1
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id: id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}
