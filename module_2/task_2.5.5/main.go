package main

/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
*/

import (
	"fmt"
	"strings"
)

func main() {

	var s string
	fmt.Scan(&s)
	for _, ch := range s {
		if strings.Count(s, string(ch)) > 1 {
			s = strings.Replace(s, string(ch), "", -1)
		}
	}
	fmt.Println(s)

}

/*
Альтернативное решение

func main() {
  var a string
  fmt.Scan(&a)

  for _, ch := range a {
    if strings.Count(a, string(ch)) == 1 {
      fmt.Print(string(ch))
  }
}
}

*/
