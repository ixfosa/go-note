package main

import (
	"fmt"
	"time"
)

// 速率限制 是一个重要的控制服务资源利用和质量的途径。Go 通过 Go 协程、通道和打点器优美的支持了速率限制。
func main() {
	// 假设我们想限制我们接收请求的处理，我们将这些请求发送给一个相同的通道。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	// 通过在每次请求前阻塞 limiter 通道的一个接收，我们限制自己每 200ms 执行一次请求。
	for req := range requests {
		<- limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制我们可以通过通道缓冲来实现。
	// 这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制。
	burstyLimiter := make(chan time.Time, 3)

	// 想将通道填充需要临时改变3次的值，做好准备。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200 ms 将添加一个新的值到 burstyLimiter中，直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟超过 5 个的接入请求。它们中刚开始的 3 个将由于受 burstyLimiter 的“脉冲”影响。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
