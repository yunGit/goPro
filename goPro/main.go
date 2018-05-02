// goPro project main.go
package main

import (
	"fmt"
)

type vertex struct {
	X, Y int
}

var (
	// 结构体文法
	v1 = vertex{1, 2}
	v2 = vertex{X: 1}  // Y:0被省略
	v3 = vertex{}      // x:0和y:0
	z  = &vertex{1, 2} // 类型为 *vertex
)

func main() {
	// 指针，即间接引用或非直接引用
	// go没有指针运算
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)
	fmt.Println("p = ", p)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	// 结构体
	// 即一个字段的集合，使用.访问
	v := vertex{1, 3}
	v.Y = 9
	fmt.Println(v)

	// 指针访问结构体
	q := &v
	q.X = 1e9
	fmt.Println(v)

	// http://go-tour-zh.appspot.com/moretypes/5
	fmt.Println(v1, v2, v3, *z)
}
