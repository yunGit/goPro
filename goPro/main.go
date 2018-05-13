// goPro project main.go
package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	X, Y float64
}

// go 没有类，但是可以在结构体类型上定义方法
// 方法的接受者出现在 func关键字和方法名之间的参数中
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 接受者为指针修改接受者原始值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
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

// 接口
type Abser interface {
	Abs() float64
}

// 隐式接口
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReaderWriter interface {
	Reader
	Writer
}

// Stringer, 存在于fmt中的接口，用来描述自己的类型
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	s := fmt.Sprint(float64(e))

	return fmt.Sprintf("Cannot Sqrt negative number:%s", s)
}
func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, ErrNegativeSqrt(n)
	}
	return 0, nil
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d,%d,%d", ip[0], ip[1], ip[2])
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	var n int
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
		n++
	}
	return n, nil
}

func main() {
	fmt.Println("Hello world!")
	// 方法
	v := Vertex{3, 4}
	// 结构体方法的调用
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// 接受者为指针的方法
	// 方法可以与命名类型或命名类型的指针关联
	// 有两个原因需要使用指针接受者，首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）
	// 其次，方法可以修改接受者指向的值
	v.Scale(5)
	fmt.Println(v, v.Abs())

	// 接口，接口类型是由一组方法定义的集合
	// 接口类型的值可以存放实现这些方法的任何值
	var a Abser
	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())
	// a = v				// !! Error !!，Only *Vertex.Abs, not defined Vertex.Abs
	// fmt.Println(a.Abs())

	// 隐式接口，隐式接口解耦了实现接口的包和定义接口的包：互不依赖
	// 因此也就无需再每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义
	var w Writer
	// os.stdout 实现了 writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")

	// Stringer
	y := Person{"arthur dent", 42}
	z := Person{"zaphod beeblebrox", 90011}
	fmt.Println(y, z)

	// Error,错误
	// 使用error值来表示错误状态
	// 与fmt.stringer类似，‘error’类型是一个内建接口
	//	type error interface {
	//		Error() string
	//	}
	// 通常函数会返回一个error值，调用的他的代码应当判断这个错误是否等于'nil'，来进行错误处理
	ii, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v \n", err)
	}
	fmt.Println("Converted integer:", ii)
	// 自定义错误值
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	ip := IPAddr{127, 0, 0, 1}
	fmt.Println(ip)

	// Reader 接口
	//	 func (T) Read(b []byte) (n int, err error)
	r := strings.NewReader("Hello World!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	// 自己定义Read
	//	myReader := MyReader{}
	//	bb := make([]byte, 8)
	//	for true {
	//		n, err := myReader.Read(bb)
	//		fmt.Printf("n = %v err = %v bb = %v\n", n, err, bb)
	//		fmt.Printf("bb[:n] = %q\n", bb[:n])
	//		if err == io.EOF {
	//			break
	//		}
	//	}
}
