/*
На ввод подаются пять целых чисел, которые записываются в массив.
Нужно найти и вывести максимальное число в этом массиве.
Example:
Input:
2
3
56
45
21
Output:
56
*/

package main

import "fmt"

func main() {
	
	array := [5]int{}
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	max := array[0]
	
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
	
}