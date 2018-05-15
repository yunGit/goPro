// goPro project main.go
package main

import (
	"fmt"
)

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

func main() {
	fmt.Println("Hello world!")

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

	for {

	}
}
