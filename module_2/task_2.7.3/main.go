package main

import "fmt"

/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
*/

func main() {

	var num string
	fmt.Scan(&num)
	n := []rune(num)
	max := 0
	for i := 0; i < len(n); i++ {
		if int(n[i]-'0') > max {
			max = int(n[i] - '0')
		}
	}
	fmt.Println(max)

}

/*
Альтернативное решение

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	var max int32
	for _, d := range s {
		if d > max {
			max = d
		}
	}
	fmt.Println(string(max))
}

*/

/*
Альтернативное решение

func main() {
	var s string
	fmt.Scan(&s)
	sl:=strings.Split(s,"")
	sort.Strings(sl)
	fmt.Print(sl[len(s)-1])
}
*/
