package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита.
На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/

func main() {

	var pass string
	fmt.Scan(&pass)
	if utf8.RuneCountInString(pass) >= 5 {
		for _, ch := range pass {
			if checkSymbol(ch) {
				continue
			} else {
				fmt.Println("Wrong password")
				return
			}
		}
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

}

func checkSymbol(ch rune) bool {
	if unicode.IsDigit(ch) || (unicode.IsLetter(ch) && unicode.Is(unicode.Latin, ch)) {
		return true
	}
	return false
}

/*
Альтернативное решение

func main() {
	var pass string
	fmt.Scan(&pass)

	if isPasswordValid(pass) {
		fmt.Printf("Ok")
	} else {
		fmt.Printf("Wrong password")
	}
}

func isPasswordValid(password string) bool {
	if len([]rune(password)) < 5 {
		return false
	}

	for _, char := range password {
		if !unicode.IsDigit(char) && !unicode.Is(unicode.Latin, char) {
			return false
		}
	}

	return true
}
*/
