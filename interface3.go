package main

import (
	"fmt"
)

type Person struct {
	name string
}

type Animal struct {
	aname string
}

func (p Person) eat() {
	fmt.Println(p.name + " person eat")
}

func (a Animal) eat() {
	fmt.Println(a.aname + " animal eat")
}

type Eat interface {
	eat()
}

type EEat interface {
	eat()
}

func main() {
	var a Person = Person{"wk"}
	var b Animal = Animal{"hellomumu"}
	a.eat()
	b.eat()
	//a = b 只有接口才可以互相赋值
	var c Eat = Person{"wk"}
	var d EEat = Animal{"hellomumu"}
	c.eat()
	d.eat()
	c = d
	c.eat()

}
