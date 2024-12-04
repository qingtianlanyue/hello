package demo

import "fmt"

// var num = 100

// const Mode = 1

type Person struct {
	name string
	Age  int
}

func Add(x, y int) int {
	return x + y
}

func say() {
	var myName = "Tom"
	fmt.Println(myName)
}
