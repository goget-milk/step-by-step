package main

import "fmt"

/*
  Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
*/

func main() {
	// var n int

	// fmt.Scan(&n)

	n := 2010

	fmt.Println(n % 100 / 10)
	fmt.Println("Hello! World")

}

