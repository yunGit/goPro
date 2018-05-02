// goPro project main.go
package main

import (
	"fmt"
	"math/rand"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Hello World!")
	v := 42 // change me!
	f := 3.14
	g := 0.875 + 0.5i
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)

	const constbool = true
	fmt.Println("constbool = ?", constbool)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println("rand = ", rand.Intn(20))
}
