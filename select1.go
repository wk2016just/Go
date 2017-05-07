package main

import (
	"fmt"
)

func main() {
	var c2, c3 chan string
	var i1, i2 string
	c1 := make(chan string)

	go func() { c1 <- "nihao" }()

	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}

	}
}
