package main

/*
На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой.
Если строка правильная - вывести Right иначе - вывести Wrong
*/

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	txt := []rune(text)
	if unicode.IsLetter(txt[0]) && unicode.IsUpper(txt[0]) && txt[len(txt)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
