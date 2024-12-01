package main

import "fmt"

/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
*/

func main() {

	var num string
	fmt.Scan(&num)
	n := []rune(num)
	for i := 0; i < len(n); i++ {
		fmt.Print(int(n[i]-'0') * int(n[i]-'0'))
	}
}

/*
Альтернативное решение

func main() {
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		c -= '0'
		fmt.Print(c * c)
	}
}
*/

/*
Альтернативное решение

func main() {
  var n int
  fmt.Scan(&n)
  for n != 0 {
    defer fmt.Print((n%10)*(n%10))
    n/=10
  }
}
*/
