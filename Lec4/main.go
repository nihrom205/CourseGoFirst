package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	// boolean operator
	aBoolean, bBoolean := true, true
	fmt.Println("AND", aBoolean && bBoolean)
	fmt.Println("OR", aBoolean || bBoolean)
	fmt.Println("NOT", !aBoolean)

	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	var a int = 32
	b := 96
	fmt.Println("a + b = ", a+b)

	var aInt int8 = 55
	// вывод через %T форматирование
	fmt.Printf("Type is %T\n", aInt)

	//узнаем, сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	//эксперимент
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	// эксперимент 2
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	// эксеримент 3		приведение типов
	var third64 int64 = 161233414
	var fourth int = 345678
	fmt.Println(third64 + int64(fourth))

	// float
	// float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of %T\n", floatFirst, floatSecond)
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub", sub)
	fmt.Printf("Sum: %.3f and Sub: %.3f\n", sum, sub)

	// Numeric, Complex
	c1 := complex(5, 7)
	c2 := complex(5, 7)
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	// String. Строка - это набор БАЙТ
	// name := "Fedya"
	name := "Федя"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string :", name, len(name)) // функция len() возвращает количество элементов в наборе
	fmt.Println("Amoount of chars:", name, utf8.RuneCountInString(name))

	// Rune - руна. Это один utf-ный символ.

	// Поиск подстроки в строке
	//totalString, subString := "ABCDEFG", "CDE"
	totalString, subString := "ABCDEFG", "cde"
	fmt.Println(strings.Contains(totalString, subString))

	// rune -> alias int32
	var sampleRune rune
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)

	var anotherRune rune = 'Q' // для инициализации руны используйте ''
	fmt.Printf("Rune as char %c\n", anotherRune)

	var numberrRune rune = 2345679
	fmt.Printf("Rune as char %c\n", numberrRune)

	var charRune rune = 'Ж'
	fmt.Printf("Rune as char %c\n", charRune)

	// "A" < "abcd"
	// 0 if a == b, -1 if a < b, and +1 if a > b
	fmt.Println(strings.Compare("a", "a"))	// 0
	fmt.Println(strings.Compare("a", "b"))	// -1
	fmt.Println(strings.Compare("abcd", "b"))	// -1
	fmt.Println(strings.Compare("b", "a"))	// +1
	fmt.Println(strings.Compare("b", "abcd"))	// +1

	var by byte // alias uint8
	fmt.Println("Byte: ", by)
}
