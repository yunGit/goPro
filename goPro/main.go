// goPro project main.go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// go 没有类，但是可以在结构体类型上定义方法
// 方法的接受者出现在 func关键字和方法名之间的参数中
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 可以对保重的 任意 类型定义任意方法，而不仅仅针对结构体
// 但是，不能对来自其他包的类型或基础类型定义方法
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	fmt.Println("Hello world!")
	// 方法
	v := &Vertex{3, 4}
	// 结构体方法的调用
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// next http://go-tour-zh.appspot.com/methods/2
	// 接受者为指针的方法
	// 方法可以与命名类型或命名类型的指针关联
	// 有两个原因需要使用指针接受者，首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）
	// 其次，方法可以修改接受者指向的值

}
