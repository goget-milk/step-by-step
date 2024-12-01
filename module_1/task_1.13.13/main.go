package main

import "fmt"

/*
Дано натуральное число A > 1.Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A.
Если А не является числом Фибоначчи, выведите число -1.
*/

func main() {

	var n int
	fmt.Scan(&n)
	var a, b int = 1, 1
	i := 2
	for {
		i++
		a, b = b, a+b
		if b == n {
			fmt.Println(i)
			return
		} else if b > n {
			fmt.Println(-1)
			return
		}
	}

}
