package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Rectangle struct {
	length float64
	width  float64
}
type Student struct {
	Person
	class string
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am  %d years old", p.name, p.age)
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
