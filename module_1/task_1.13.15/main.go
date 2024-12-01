package main

import "fmt"

/*
Вводятся натуральное число и цифра, которую нужно удалить.
*/

func main() {

	var num, n int
	fmt.Scan(&num, &n)
	numbers := []int{}
	for i := 100000000; i >= 1; i /= 10 {
		if num/i == 0 {
			continue
		}
		if num/i%10 == n {
			continue
		}
		numbers = append(numbers, num/i%10)
	}
	for _, n := range numbers {
		fmt.Print(n)
	}

}

/*
Альтернативное решение

func main() {
  var a, b int
  fmt.Scan(&a, &b)
  result := 0
  for i := 1; a > 0; {
    if a % 10 != b {
      result += a % 10 * i
      i *= 10
    }
    a /= 10
  }
  fmt.Println(result)
}
*/

/*
Альтернативное решение

func main() {
	var num, n int
	fmt.Scan(&num, &n)
	for ; num > 0; num /= 10 {
		if num % 10 != n {
			defer fmt.Print(num % 10)
		}
	}
}
*/
