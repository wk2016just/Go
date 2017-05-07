package main

import (
	"fmt"
)

type user struct {
	id int
}

func g(any interface{}) int {
	return any.(user).id
}

func si(i, j int) interface{} {
	if i == j {
		return "equal"
	}
	return i + j
}

func main() {
 var u user = user{1}
  i := g(u)
  fmt.Println(i)
  
  fmt.Println(si(2,2))
}
