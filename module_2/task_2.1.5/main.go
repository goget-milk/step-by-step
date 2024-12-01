/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму.
Пример вызова функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)
Результат: 2, 1
*/

package main

import "fmt"

func sumInt(n ...int) (c, s int) {
	for _, e := range n {
		c++
		s += e
	}
	return
}

func main() {
	count, sum := sumInt(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(count, sum)
}
