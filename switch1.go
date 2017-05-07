package main

import "fmt"

func main() {
	var x int = 1
	whatAmI := func(i interface{}) {
		switch i.(type) { //t初始化为i的type，这里whatami是匿名函数
		case bool, uint: //case , 可以多个表达式
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type\n")
		}
	}
	whatAmI(x)

}
