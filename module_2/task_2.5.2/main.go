package main

import (
	"fmt"
	"reflect"
)

/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.
*/

func main() {

	var txt string
	fmt.Scan(&txt)
	textArr := []rune(txt)
	if reflect.DeepEqual(textArr, reverse(textArr)) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}

func reverse(txt []rune) []rune {
	var revTxt []rune
	for i := range txt {
		revTxt = append(revTxt, txt[len(txt)-1-i])
	}
	return revTxt
}

/*
Альтернативное решение

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	str := in.Text()
	arr := []rune(str)
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			fmt.Println("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}

*/

/*
Альтернативное решение

func main() {
  str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  text := []rune(str)
  for i, j := range text{
    if text[len(text)-i-1] != j{
      fmt.Print("Нет")
      return
    }
  }
  fmt.Print("Палиндром")
}

*/
