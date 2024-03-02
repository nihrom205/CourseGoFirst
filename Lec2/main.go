package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("two line")

	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")

	// форматированный вывод
	fmt.Printf("\nHello, my name is %s\n", "Bob")
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)


	// декларирование переменных
	var age int
	fmt.Println("my age, is", age)
	age = 32
	fmt.Println("my age, is", age)

	// декларирование и инициализация пользовательских значений
	var height int = 180
	fmt.Println("my height, is", height)


	// полустрогая типизация
	// компилятор угадает тип
	var uuid = 12345
	fmt.Println("my uuid, is", uuid)
	fmt.Printf("тип: %T\n", uuid)

	// декаларирование и инициалзация переменных одного типа
	var firstVar, secondVar int = 20, 30
	fmt.Printf("firstVar:%d, secondVar:%d", firstVar, secondVar)

	// декларирование блока переменных
	var (
		personName string = "Bob"
		personAge = 33
		personUID int
	)
	fmt.Printf("Name: %s\nAge: %d\nUid: %d\n", personName, personAge, personUID)


	// немного странного
	var a, b = 20, "Vova"
	fmt.Println(a, b)

	// немного хорошено. повторное декларирование переменной приводит к ошибке компиляции
	//var a = 200

	// короткоя декларация (объявление)
	count := 10
	fmt.Println("Count: ", count)
	count++
	fmt.Println("Count: ", count)

	// множественное присваивание через :=
	aArg, bArg := 10, 20
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	// aArg, bArg := 30, 40
	// fmt.Println(aArg, bArg) // no new variables on left side of :=

	// исключение из этого правила
	bArg, cArg := 300, 400
	fmt.Println(aArg, bArg, cArg)	//30 300 400


	// пример
	width, length := 20.5, 30.2
	fmt.Printf("Min dimensinal of rectangle is : %.3f", math.Min(width, length))
}