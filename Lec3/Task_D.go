package main

import "fmt"

func main() {
	var (
		firstName string
		lastName string
		age int
	)

	fmt.Scan(&firstName, &lastName, &age)
	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS", firstName, lastName, age)
}
