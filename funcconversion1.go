package main

import "fmt"

type adder interface {
add(string) int
}
type handler func(name string) int

func (h handler) add(name string) int {
return h(name) + 10
}

func doubler(name string) int {
return len(name) * 2
}

// func tester(name int) int{
//   return name * 2
// }


//非函数类型
type myint int

//实现了adder接口
func (i myint) add(name string) int {
return len(name) + int(i)
}

func main() {
    var my handler = func(name string) int {return len(name)}
     fmt.Println(my.add("taozs")) 
     
     fmt.Println(handler(doubler).add("taozs"))
    //fmt.Println( handler(doubler(tester)).add("wk") )

     
}