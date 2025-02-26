package main

import "fmt"

type A struct {
	year int
}

func (a A) Greet() {
	fmt.Println("Hello, I'm a year", a.year)
}

type B struct {
	A
}

//func (b B) Greet() {
//	fmt.Println("Welcome, I'm a year", b.year)
//}

func main() {
	var a A
	a.year = 2016
	var b B
	b.year = 2024
	a.Greet()
	b.Greet()
}
