package main

import "fmt"

/*
Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/

func main() {
	var n, max, count int

	for {
		fmt.Scan(&n)
		if n <= 0 {
			break
		}
		if n > max {
			max = n
			count = 1
		} else if max == n {
			count++
		}
	}
	fmt.Println(count)
}

/*
Альтернативное решение
func main(){
  var a, max, count int

  for fmt.Scan(&a); a != 0 ; fmt.Scan(&a){
    switch {
      case a > max:
      max, count = a, 1
      case a == max:
      count++
    }
  }
  fmt.Println(count)
}
*/
