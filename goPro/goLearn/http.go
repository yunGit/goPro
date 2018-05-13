// goPro project http.go
package http

import (
	"fmt"
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello")
}

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	name string
	age  int
}

func (s Struct) String() string {
	return fmt.Sprintf("name:%s, age:%d", s.name, s.age)
}
func (stru *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, stru)
}

func main() {
	fmt.Println("Hello world!")

	// HTTP server
	//	 包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求
	var h hello
	err := http.ListenAndServe("localhost:4001", h)
	if err != nil {
		log.Fatal(err)
	}
	//	 监听固定路径请求
	http.Handle("/string", String("This is a string type request!!"))
	http.Handle("/struct", &Struct{"cuicuicui", 30})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
