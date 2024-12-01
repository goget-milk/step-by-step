package main

import "fmt"

/*
Дано трехзначное число. Переверните его, а затем выведите.
*/

func main() {

	var num int
	fmt.Scan(&num)
	result := 0
	for num > 0 {
		result *= 10
		result += num % 10
		num /= 10
	}
	fmt.Println(result)

}

/*
Альтернативное решение

func main(){
  var n int
  fmt.Scan(&n)
  for ;n > 0; n /= 10{
    fmt.Print(n % 10)
  }
}
*/
