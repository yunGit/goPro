// goPro project main.go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
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

	fmt.Println(sqrt(2), sqrt(-4))
}
