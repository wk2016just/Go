package main

import "fmt"

func main() {
	func() int {
		var i = 1
		fmt.Println(i)
		return i
	}()

	func(b int) {
		fmt.Println(b)
	}(33)

	a := func() {
		fmt.Println("a")
	}
	a()
}
