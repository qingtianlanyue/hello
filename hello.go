package main

import (
	"fmt"

	"github.com/qingtianlanyue/hello/demo"
)

var x int8 = 10

const pi = 3.14

func init() {
	fmt.Println("x:", x)
	fmt.Println("pi:", pi)
	sayHi()
}

func sayHi() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("你好，世界！")
	fmt.Println(demo.Add(1, 2))
}
