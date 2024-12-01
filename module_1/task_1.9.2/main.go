package main

import "fmt"

/*
  По данному трехзначному числу определите, все ли его цифры различны.
*/

func main() {
	var num int
	fmt.Scan(&num)

	f := num / 100
	s := (num % 100) / 10
	t := num % 10

	if f != s && f != t && s != t {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
