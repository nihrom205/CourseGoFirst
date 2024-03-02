package main

import (
	"fmt"
	"strings"
)

func main() {
	var login, email string
	fmt.Scan(&login, &email)

	if len(login) < 10 || strings.Contains(login, "@") {
		fmt.Println("Некорректный логин")
	} else if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		fmt.Println("Некорректная почта")
	} else {
		fmt.Println("ОК")
	}
}
