// goPro project complexType.go
package goLearn

import (
	"fmt"
	"math"
	"strings"
	//	"code.google.com/p/go-tour/wc"
)

type vertex struct {
	X, Y float64
}

var (
	// 结构体文法
	v1 = vertex{1, 2}
	v2 = vertex{X: 1}  // Y:0被省略
	v3 = vertex{}      // x:0和y:0
	z  = &vertex{1, 2} // 类型为 *vertex
)

// map
var m = map[string]vertex{
	"Google": vertex{
		556.3, -113.43,
	},
}

func printSlice(s string, x []int) {
	// len 长度，指已经被赋值过的最大下标+1
	// cap 容量，指切片目前可容纳的最多元素个数
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	var a = strings.Fields(s)
	for i := 0; i < len(a); i++ {
		v := m[a[i]]
		m[a[i]] = v + 1
	}
	return m
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Println("x = ", x, ", sum = ", sum)
		return sum
	}
}

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

	// 结构体文法
	fmt.Println(v1, v2, v3, *z)

	// 数组
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// slice
	// 一个slice会指向一个序列的值，并且包含了长度信息
	// []T是一个元素类型为T的slice，与引用的区别即[]内为空
	p1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p1 == ", p1)

	for i := 0; i < len(p1); i++ {
		fmt.Printf("p1[%d] == %d\n", i, p1[i])
	}

	// slice切片
	fmt.Println("p1[1:4] == ", p1[1:4])
	// 	省略下标表示从0开始
	fmt.Println("p1[:3] == ", p1[:3])
	// 	省略上标表示到 len(s)结束
	fmt.Println("p1[4:] == ", p1[4:])

	// 构造slice
	a1 := make([]int, 5)
	printSlice("a1", a1)
	b := make([]int, 0, 5)
	printSlice("b", b)
	// 数组或切片的引用来初始化切片
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	printSlice("p1", p1)

	// nil slice，slice 的零值是 nil
	// 一个nil的slice长度和容量是0
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	// 向slice添加元素
	// append ， 结果是个包含原slice 所有元素加上新添加元素的slice
	// append works on nil slice
	z = append(z, 0)
	printSlice("z", z)
	// the slice grows as needed
	z = append(z, 1)
	printSlice("z", z)
	// we can add more than one element at a time
	z = append(z, 2, 3, 4)
	printSlice("z", z)

	// for 循环的range格式可以对 slice 或者 map进行迭代循环
	for i, v := range z {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// map
	// 映射键到值，使用之前必须用make而不是new来创建；值为nil的map是空的，并且不能赋值
	//	m = make(map[string]vertex)
	// map文法
	// map文法跟结构体文法相似，不过必须有键名
	m["Bell Labs"] = vertex{
		40.6, -73.3,
	}
	fmt.Println(m)
	// 插入或修改map
	m["answer"] = vertex{45, -43}
	fmt.Println("The value:", m["answer"])
	// delete删除元素
	delete(m, "Bell Labs")
	fmt.Println("The value:", m)

	// 检测"Bell Labs“是否在m中，是则ok为true，否false；v1或为map的零值
	v1, ok := m["Bell Labs"]
	fmt.Println("The value:", v1, "Present?", ok)

	// 测试 http://go-tour-zh.appspot.com/moretypes/19
	//	wc.Test(WordCount)

	// 函数值，函数也是值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	// 函数的闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
