package main

import "fmt"

/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности,
во второй строке -- n чисел, входящих в данную последовательность.
*/


func main() {
	var n int
	fmt.Scan(&n)
	var x int
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x/100 == 0 && x >= 10 && x%8 == 0 {
			sum += x
		}
	}
	
}