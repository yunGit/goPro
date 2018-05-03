// goPro project main.go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	fmt.Println("Hello world!")
	// 方法
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}
