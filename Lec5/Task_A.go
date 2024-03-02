package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var text string
	fmt.Scan(&text)

	countChar := utf8.RuneCountInString(text)
	rub := countChar * 23 / 100
	kop := countChar * 23 % 100
	if rub == 0 {
		fmt.Printf("%d коп.", kop)
		return
	}
	fmt.Printf("%d р. %d коп.", rub, kop)
}