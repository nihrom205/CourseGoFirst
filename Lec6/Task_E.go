package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var pass1, pass2 string
	for {
		fmt.Scan(&pass1, &pass2)

		if utf8.RuneCountInString(pass1) < 8 {
			fmt.Println("Слишком короткий пароль!")
		} else if strings.Contains(pass1, "123") || strings.Contains(pass1, "qwe") {
			fmt.Println("Слишком простой пароль!")
		} else if pass1 != pass2 {
			fmt.Println("Введенные пароли различаются!")
		} else {
			fmt.Println("Ну наконец-то!")
			break
		}
	}
}
