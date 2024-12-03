/*
Дана строка, содержащая только английские буквы (большие и маленькие).
Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.Replace(str, "", "*", -1)
	str = strings.Trim(str, "*")
	fmt.Println(str)
}

/*
Альтернативное решение

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(strings.Join(strings.Split(s, ""), "*"))
}
*/

/*
  Альтернативное решение

func main() {
	var s string
	fmt.Scan(&s)
	split := strings.Split(s, "")
	join := strings.Join(split, "*")
	fmt.Println(join)
}
*/
