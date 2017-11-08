package main

import (
	"runtime"
	"fmt"
	"time"
)

func main() {
	// goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，
	// Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。
	// 也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。
	// goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。
	// go hello(a, b, c)

	demoGoroutine()

	demoChannel()

	demoBufferedChannel()

	demoRangeClose()

	demoSelect()

	demoExpire()
}


func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		fmt.Println(s)
	}
}

func demoGoroutine() {
	// 但在Go 1.5以前调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，需要在我们的程序中显式调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。
	// GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。
	//  默认情况下，在Go 1.5将标识并发系统线程个数的runtime.GOMAXPROCS的初始值由1改为了运行环境的CPU核数。

	fmt.Println(runtime.NumCPU())
	go say("world") //开一个新的Goroutines执行
	say("hello") //当前Goroutines执行
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	// channel通过操作符<-来接收和发送数据
	// 发送total到channel c.
	c <- total  // send total to c
}

func demoChannel() {
	// goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢，Go提供了一个很好的通信机制channel。
	//  channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。
	// 定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	// 从c中接收数据，并赋值给x, y
	x, y := <-c, <-c  // receive from c
	// 默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。
	// 所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
	// 其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。


	fmt.Println(x, y, x + y)
}

func demoBufferedChannel()  {
	c := make(chan int, 2)//修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	// 改成1报，fatal error: all goroutines are asleep - deadlock!
}


func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	// 生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。
	// 如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。
	// 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
	close(c)
}

func demoRangeClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		// 如果存在多个channel的时候,通过select可以监听channel上的数据流动。
		// select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。
		// 在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func demoSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func demoExpire() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <- c:
				println(v)
			// 有时候会出现goroutine阻塞的情况,可以利用select来设置超时
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<- o
}