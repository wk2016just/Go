package main

import (
	"fmt"
)

type user struct {
	id   int
	name string
}

func main() {
	wk := user{2, "wk"}

	fmt.Println(wk.id)
}
