package main
import (
    "math"
    "fmt"
)

type Shape interface{
    area() float64
}    

type Circle struct{
    x,y,radius float64 //结构体也可以连着定义
}

type Rectangle struct{
    width,height float64
}

type Square struct{
    width int64
}

func(rect Rectangle) area() float64{
    return rect.width * rect.height
}

func(circle Circle) area() float64{
    return circle.radius * circle.radius * math.Pi
}

func(square Square) area() int64{
    return square.width * square.width 
}

func getArea(shape Shape) float64{
    return shape.area()
}

func main() {
  
  circle := Circle{x:0,y:0,radius:4}
  rect := Rectangle{width:10,height:5} 
  square := Square{10}
  
  fmt.Printf("circle area:%f\n",getArea(circle))
  fmt.Printf("rect area:%f",getArea(rect))
  fmt.Printf("square area:%d",getArea(square))
  
}