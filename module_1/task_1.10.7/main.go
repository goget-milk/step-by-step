package main

import "fmt"

/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
*/

func main() {
	var x, p, y, sum, years int

	fmt.Scan(&x, &p, &y)
	for {
		if x >= y {
			break
		}
		sum = (x * p) / 100
		x += sum
		years++
	}
	fmt.Println(years)
}

/*
Альтернативное решение

func main() {
	var x, p, y, n int
	fmt.Scan(&x, &p, &y)
	for n = 0; x <= y; n++ {
		x += x * p / 100
	}
	fmt.Println(n)
}

*/
