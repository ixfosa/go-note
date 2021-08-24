package main

func main() {
	
}
/*
web 工作方式的几个概念
	Request：用户请求的信息，用来解析用户的请求信息，包括 post、get、cookie、url 等信息

	Response：服务器需要反馈给客户端的信息

	Conn：用户的每次请求链接

	Handler：处理请求和生成返回信息的处理逻辑
*/

/*
http 包执行流程

	1.创建 Listen Socket, 监听指定的端口，等待客户端请求到来。

	2.Listen Socket 接受客户端的请求，得到 Client Socket, 接下来通过 Client Socket 与客户端通信。

	3.处理客户端的请求，首先从 Client Socket 读取 HTTP 请求的协议头，如果是 POST 方法，
	还可能要读取客户端提交的数据，然后交给相应的 handler 处理请求，handler 处理完毕准备好客户端需要的数据，
	通过 Client Socket 写给客户端。


这整个的过程里面我们只要了解清楚下面三个问题，也就知道 Go 是如何让 Web 运行起来了
	如何监听端口？
	如何接收客户端请求？
	如何分配 handler？

前面小节的代码里面我们可以看到，Go 是通过一个函数 ListenAndServe 来处理这些事情的，这个底层其实这样处理的：
初始化一个 server 对象，然后调用了 net.Listen("tcp", addr)，也就是底层用 TCP 协议搭建了一个服务，然后监控我们设置的端口。

下面代码来自 Go 的 http 包的源码，通过下面的代码我们可以看到整个的 http 处理过程：
	func (srv *Server) Serve(l net.Listener) error {
		defer l.Close()
		var tempDelay time.Duration // how long to sleep on accept failure
		for {
			rw, e := l.Accept()
			if e != nil {
				if ne, ok := e.(net.Error); ok && ne.Temporary() {
					if tempDelay == 0 {
						tempDelay = 5 * time.Millisecond
					} else {
						tempDelay *= 2
					}
					if max := 1 * time.Second; tempDelay > max {
						tempDelay = max
					}
					log.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
					time.Sleep(tempDelay)
					continue
				}
				return e
			}
			tempDelay = 0
			c, err := srv.newConn(rw)
			if err != nil {
				continue
			}
			go c.serve()
		}
	}
监控之后如何接收客户端的请求呢？上面代码执行监控端口之后，调用了 srv.Serve(net.Listener) 函数，这个函数就是处理接收客户端的请求信息。
这个函数里面起了一个 for{}，首先通过 Listener 接收请求，其次创建一个 Conn，最后单独开了一个 goroutine，
把这个请求的数据当做参数扔给这个 conn 去服务：go c.serve()。这个就是高并发体现了，用户的每一次请求都是在一个新的 goroutine 去服务，
相互不影响。

那么如何具体分配到相应的函数来处理请求呢？conn 首先会解析 request:c.readRequest(),
然后获取相应的 handler:handler := c.server.Handler，也就是我们刚才在调用函数 ListenAndServe 时候的第二个参数，
我们前面例子传递的是 nil，也就是为空，那么默认获取 handler = DefaultServeMux, 那么这个变量用来做什么的呢？
对，这个变量就是一个路由器，它用来匹配 url 跳转到其相应的 handle 函数，那么这个我们有设置过吗？
有，我们调用的代码里面第一句不是调用了 http.HandleFunc("/", sayhelloName) 嘛。
这个作用就是注册了请求 / 的路由规则，当请求 uri 为 "/"，路由就会转到函数 sayhelloName，DefaultServeMux 会调用 ServeHTTP 方法，
这个方法内部其实就是调用 sayhelloName 本身，最后通过写入 response 的信息反馈到客户端。

*/