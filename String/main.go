package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "hello world"
	fmt.Println(name)

	//1. Строка - это байтовый слайс со своими особенностями при
	// обращении  к нижележащему массиву
	word := "Тестовая строка"
	fmt.Printf("String:%s\n", word)
	// Какие значения байт сейчас находятся в слайсе word?
	fmt.Printf("Bytes:  ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) // %х - формат представления 16-ти ричного байта
	}
	fmt.Println()
	// Каким образом получать доступ к отдельно стоящим символам
	fmt.Printf("Characters:  ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) // %с - формат представления символв
	}
	fmt.Println()
	// 2. Строки в GO хранятся как наборы UTF-8 символов. Каждый символ
	// может занимать больше чем 1 байт

	// 3. Руна (Rune) - это стандартный встроенный тип в GO, позволяющий
	// хранить единый неделимый UTF cимвол ВНЕ ЗАВИСИМОСТИ ОТ ТОГО сколько байт он занимает
	fmt.Printf("Runes: ")
	runeSlice := []rune(word) // Преобразование слайса байт к слайсу рун []byte{sliceRune}
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c", runeSlice[i])
	}
	fmt.Println()
	// Итерирование по строке с использование рун
	for id, runeVal := range word { // id - это индекс байта, с КОТОРОГО НАЧИНАЕТСЯ РУНА runeVal
		fmt.Printf("%c stats at position %d\n", runeVal, id)
	}
	// []byte{10,20,30,40,20,30,40,50}
	// //0 "value"
	// //4 "another value"

	//5. Создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43} // Исходное представление байтов
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDecimalBytesSlice := []byte{100, 101, 102, 103} // Синтаксический сахар - можно использовать десятичное представление байтов
	myDecimalByteStr := string(myDecimalBytesSlice)
	fmt.Println(myDecimalByteStr)

	// 6. Создание строки из слайса рун
	// Руны как hex
	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFormRune := string(runeHexSlice)
	fmt.Println("form Runes (hex):", myStrFormRune)
	// Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 'y', 'a'} // - ' ' таким образом обозначается руна
	myStrfromRuneLiteral := string(runeLiteralSlice)
	fmt.Println("from Runes(literals)", myStrfromRuneLiteral)
	fmt.Printf("%s and %T\n", myStrfromRuneLiteral, myStrfromRuneLiteral)

	// 7. Длина и емкость строки
	// Длина len() количетсво байт с слайсе
	fmt.Println("Length of Вася", len("Вася"), "Bytes")
	// Длина RuneCounter - количвество [рун]
	fmt.Println("Length of Вася", utf8.RuneCountInString("Вася"), "Runes")
	// Вычисление емкости строки бессмыслекнно - так как строка - базовый тип
	fmt.Println(cap([]rune("Вася")))

	// 8. Сравнение строк == и !=. Начиная с go 1.6
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	// 9. Конкатенация
	word3 := word1 + word2
	fmt.Println(word3)

	// 10. Строитель строк (Java -> StringBuilder)
	firstname := "Alex"
	secondName := "Jonhson"
	fullName := fmt.Sprintf("%s #### %s", firstname, secondName)
	fmt.Println("Fullname:", fullName)

	// 11. Строки неизменяемые
	// fullName[0] = "Q" strings are immutable

	// 12. А Слайсы изменяемые
	fullNameSlices := []rune(fullName)
	fullNameSlices[0] = 'F'
	fullName = string(fullNameSlices)
	fmt.Println("String mutation:", fullName)
	//13. Сравнение рун
	if 'Q' == 'Q' {
		fmt.Println("Runes equal")
	} else {
		fmt.Println("Runes  not equal")
	}

	// 14. Где живут полезные методы со строками?
	// import "strings"

}
