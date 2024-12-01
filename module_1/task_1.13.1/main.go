package main

import "fmt"

/*
Дано трехзначное число. Найдите сумму его цифр.
*/

func main() {

	var num, sum int
	fmt.Scan(&num)
	for i := 100; i >= 1; i /= 10 {
		sum += num / i % 10
	}
	fmt.Println(sum)

}

/*
Альтернативное решение

func main() {
  var a int
  fmt.Scan(&a)
  sum := 0
  for a > 0 {
    sum += a % 10
    a /= 10
  }
  fmt.Println(sum)
}
*/

/*
Альтернативное решение

func main(){
  var n,sum int

  for fmt.Scan(&n); n>0; n/=10{
    sum+=n%10
  }

  fmt.Print(sum)
}
*/
