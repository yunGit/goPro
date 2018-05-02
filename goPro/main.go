// goPro project main.go
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// 牛顿法开根号
	z := 1.0
	zOld := 0.0
	for math.Pow(z-zOld, 2) >= 0.0000000000000001 {
		zOld = z
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}
	return fmt.Sprintf("%g", z)
}

func pow(x, n, lim float64) float64 {
	// v 的作用域在if语句块之内
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	sum := 0
	// for
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//	// 死循环
	//	for {
	//	}

	// if - for
	fmt.Println(sqrt(2), sqrt(-4))

	// if - else
	fmt.Println(
		pow(1.414, 2, 10),
		pow(3, 3, 20),
	)

	// switch
	fmt.Printf("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s\n", os)
	}

	// no condition switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// next http://go-tour-zh.appspot.com/flowcontrol/12
}
