/*
Дана последовательность, состоящая из целых чисел. 
Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.
Example:
Input:
5
1 2 3 -1 -4
Output:
3
*/

package main

import "fmt"

func main() {
	
	var n, num, count int
	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num > 0 {
			count++
		}
	}
	
	fmt.Println(count)

}