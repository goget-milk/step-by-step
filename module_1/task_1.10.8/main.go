package main

import "fmt"

/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
*/

func main() {
	var num1, num2 int = 564, 8954
	result1 := 0
	result2 := 0
	// fmt.Scan(&num1, &num2)
	for num1 > 0 {
		result1 *= 10
		result1 += num1 % 10
		num1 /= 10
	}
	for num2 > 0 {
		result2 *= 10
		result2 += num2 % 10
		num2 /= 10
	}
	for i := 0; i < 4; i++ {
		n1 := result1 % 10
		if n1 == 0 {
			break
		}
		tmp := result2
		for j := 0; j < 4; j++ {
			n2 := tmp % 10
			if n1 == n2 {
				fmt.Print(n1, " ")
			}
			tmp /= 10
		}
		result1 /= 10
	}
}

/*
Альтернативное решение

func main() {
  var a, b, an, bn int
  fmt.Scan(&a, &b)

  for i := 1000; i >= 1 && a > 0; i = i/10 {
    if a/i == 0 {
      continue
    }
    an = a / i % 10
    for i := 1; i < 10000 && b > 0; i = i*10 {
      bn = b/i % 10
      if bn == an {
        fmt.Print(an," ")
        break
      }
    }
  }
}
*/

/*
Альтернаивное решение

func matchDigits(x, y int){
  if x > 10{
    matchDigits(x / 10, y)
    matchDigits(x % 10, y)
  } else if y > 10{
    matchDigits(x, y / 10)
    matchDigits(x, y % 10)
  } else if (x == y) {
    fmt.Print(x, " ")
  }
}

func main(){
  var a, b int

  fmt.Scan(&a, &b)
  matchDigits(a, b)
}
*/
