package main // package name

import (
	"fmt"
)

func main() {
	//Простейший вывод на консоль
	fmt.Println("Hello world")
	fmt.Println("Second line")
	//Функция print
	fmt.Print("One")
	fmt.Print("Two")
	//Форматированный вывод: Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)
	////////////////////////////////////////////////
	///////////////////////////////////////////////
	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assigment", age)

	//Декларирование и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is:", height)

	//В чем *полустрогость* типизации. Можно опускать тип переменной
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T/n", uid)
	// Декларирование и инициализация переменных одного типа (множественный случай)
	var firstVar, secondVar int = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar)
	//Декларирование блока переменных
	var (
		personName string = "Bob"
		personAge  int    = 42
		personUid  int
	)

	fmt.Printf("Name:%s Age:%d UID:%d\n", personName, personAge, personUid)
	// Немного странного
	var a, b = 30, "Vova" // в одну строку не получится без указаний явного типажа, компилятор сделает автоматиччески
	fmt.Println(a, b)

	//Немного хорошего. Повторное декларирование переменной приводит в ошибке компиляции
	// var a = 200

	//Короткая декларация (короткое объявление)
	count := 10
	fmt.Println("Count:", count)
	count++
	fmt.Println("Count", count)

	//Множественное присваивание через :=
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	// 	aArg, bArg := 30, 40
	// 	fmt.Println(aArg, bArg)

	//Исключение из правила
	bArg, cArg := 300, 400 // Требует как минимум одну новую переменную
	fmt.Println(aArg, cArg, bArg)

}
