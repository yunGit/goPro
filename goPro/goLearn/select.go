// goPro project select.go
package goLearn

import (
	"fmt"
	"strconv"
	"time"
)

// goroutine, Go运行时环境管理的轻量级线程
// go func()
// goroutine 在相同的地址空间中运行，因此访问共享内存必须进行同步
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值
// ch <- v // 将 v 送入 channel ch
// v := <- ch // 从 ch 接收，并且赋值给 v
// 与 map 与 slice 一样， channel 使用前必须创建
// ch := make(chan int)
// 默认情况下，在另一端准备好之前，发送和接收都会阻塞，使得 goroutine  可以在没有明确的锁或静态变量的情况下进行同步
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

// close
// 发送者可以 close 一个 channel 来表示再没有值会被发送了
// 接受者可以通过赋值语句的第二个参数来测试 channel 是否被关闭
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 关闭 channel
	close(c)
}

// select
func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("in fibonacci, x = ", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func receiveMsg() {
	c1 := make(chan string)
	c2 := make(chan string)
	t1 := time.Now() // get current time
	go func() {
		// receive three message
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 150)
			fmt.Println("receiveMsg, go1, i = ", i)
			c1 <- "msg 1 with index " + strconv.Itoa(i)
		}
	}()
	go func() {
		for j := 0; j < 3; j++ {
			time.Sleep(time.Millisecond * 200)
			fmt.Println("receiveMsg, go2, j = ", j)
			c2 <- "msg 2 with index " + strconv.Itoa(j)
		}
	}()

	// print two messge
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received msg :", msg1)
		case msg2 := <-c2:
			fmt.Println("received msg :", msg2)
		}
	}
	elapsed := time.Since(t1)
	fmt.Println("time elapsed ", elapsed)
}

func main() {
	fmt.Println("Hello world!")

	//	// http://tour.studygolang.com/concurrency/1
	//	go say("world")
	//	say("hello")

	//	a := []int{7, 2, 4, -4, 3, 0}
	//	c := make(chan int)
	//	go sum(a[:len(a)/2], c)
	//	go sum(a[len(a)/2:], c)
	//	x, y := <-c, <-c // 从 c 中获取
	//	fmt.Println(x, y, x+y)

	// 缓冲 channel
	// channel 可以是 _带缓冲的_，为 make 提供第二个草书作为缓冲长度来初始化一个缓冲 channel
	// ch := make(chan int, 100)
	// 向缓冲channel 发送数据的时候，只用在缓冲区满的时候才会阻塞，当缓冲区清空的时候接触阻塞
	//	d := make(chan int, 2)
	//	d <- 1
	//	d <- 2
	//	//	d <- 3
	//	fmt.Println(<-d)
	//	fmt.Println(<-d)
	//	//	fmt.Println(<-d)

	//	// range, 循环会从 e 不断的取数据直到 e 被关闭
	//	e := make(chan int, 10)
	//	go fibonacci(cap(e), e)
	//	for i := range e {
	//		fmt.Println(i)
	//	}

	// select
	// select使得一个 goroutine 在多个通讯操作上等待
	// select会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支，当多个都准备好的时候，会随机选择一个
	f := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			//			fmt.Println("in go func, i = ", i)
			fmt.Println(<-f)
		}
		quit <- 0
	}()
	fibonacci_select(f, quit)

	receiveMsg()

	for {

	}
}
