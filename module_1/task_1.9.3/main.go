package main

import "fmt"

/*
  Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
*/

func main() {

	var num int
	fmt.Scan(&num)
	switch {
	case num < 10 && num >= 0:
		fmt.Println(num)
	case 100 > num && num >= 10:
		fmt.Println(num / 10)
	case 1000 > num && num >= 100:
		fmt.Println(num / 100)
	case 10000 > num && num >= 1000:
		fmt.Println(num / 1000)
	case num >= 10000:
		fmt.Println(num / 10000)
	}

}

/*
  Альтернативное решение

  func main() {
	  var a int
  	fmt.Scan(&a)
  	fmt.Println(getFirstNumber(a))
  }

  func getFirstNumber(a int) int {
	  if a < 10 {
	  	return a
  	} else {
	   	return getFirstNumber(a / 10)
  	}
  }

*/
