package main

import "fmt"

/*
  Определите является ли билет счастливым. Счастливым считается билет,
  в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
*/

func main() {
	var num int
	fmt.Scan(&num)

	firstHalf := num / 1000
	lastHalf := num % 1000

	ff := firstHalf / 100
	fs := (firstHalf % 100) / 10
	ft := firstHalf % 10

	sum1 := ff + fs + ft

	lf := lastHalf / 100
	ls := (lastHalf % 100) / 10
	lt := lastHalf % 10

	sum2 := lf + ls + lt

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*
Альтернативное решение

func main() {
	var a int
	fmt.Scan(&a)
	b := a / 1000
	c := a % 1000
	bsum := (b % 10) + (b / 100) + (b / 10 % 10)
	csum := (c % 10) + (c / 100) + (c / 10 % 10)
	if(bsum == csum){
		fmt.Print("YES")
	}else {
		fmt.Print("NO")
	}
}
*/
