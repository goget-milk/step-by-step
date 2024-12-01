/*
Напишите программу, принимающая на вход число N(N≥4), а затем N целых чисел-элементов среза. 
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
Example:
Input: 
5
41 -231 24 49 6
Output: 49
*/


package main

import "fmt"

func main() { 
	
	var n, num int
	b := make([]int, 0, n)
	
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		b = append(b, num)
	}
	
	fmt.Println(b[3])	
	
}