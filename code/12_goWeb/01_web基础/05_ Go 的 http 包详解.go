package main

func main() {
	
}
/*
Go 的 http 包详解
	Go 的 http 有两个核心功能：Conn、ServeMux

Conn 的 goroutine
与我们一般编写的 http 服务器不同，Go 为了实现高并发和高性能，使用了 goroutines 来处理 Conn 的读写事件，这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。这是 Go 高效的保证。

Go 在等待客户端请求里面是这样写的：
	c, err := srv.newConn(rw)
	if err != nil {
		continue
	}
	go c.serve()
这里我们可以看到客户端的每次请求都会创建一个 Conn，这个 Conn 里面保存了该次请求的信息，然后再传递到对应的 handler，
该 handler 中便可以读取到相应的 header 信息，这样保证了每个请求的独立性。


ServeMux 的自定义#
	type ServeMux struct {
		mu sync.RWMutex   // 锁，由于请求涉及到并发处理，因此这里需要一个锁机制
		m  map[string]muxEntry  // 路由规则，一个 string 对应一个 mux 实体，这里的 string 就是注册的路由表达式
		hosts bool // 是否在任意的规则中带有 host 信息
	}

下面看一下 muxEntry
	type muxEntry struct {
		explicit bool   // 是否精确匹配
		h        Handler // 这个路由表达式对应哪个 handler
		pattern  string  // 匹配字符串
	}
接着看一下 Handler 的定义
	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)  // 路由实现器
	}
*/