package main

import (
	"fmt"
)

type A struct {
	a string
	b int64
	c *A
}

func (this *A) Print() {
	fmt.Println("a")
}

func main() {
	var a1 A
	var a2 *A

	a1.Print()
	a2.Print() // nil能调用Print()
}
