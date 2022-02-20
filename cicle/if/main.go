package main

import (
	"fmt"
	"strings"
)

func main() {
	// Классический условный оператор
	// if condition {
	// 	//body
	// }
	// Условный оператор с блоком else
	// if condition {
	// 	//body
	// } else {
	// 	//body
	// }
	var value int
	fmt.Scan(&value)
	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}

	// if condition1 {

	// } else if condition2 {

	// } else if ... {

	// } else {

	// }

	var color string
	fmt.Scan(&color)
	if strings.Compare(color, "green") == 0 {
		fmt.Println("color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	//Good Инициализация в блоке условного оператора
	// Блок присваивания - только :=
	// Инициализируемая переменная видна ТОЛЬКО в области видимого условного оператора ( в телаах if ,else if,)
	// Но не за его пределами
	if num := 10; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	// Ущербно
	//*
	// var age int = 10
	// if age > 7 {
	// 	fmt.Println("Go to school")
	// } //По факту сюда подставляется ; компилятором и дальнейший код не имеет связи с предыдущим if
	// else {
	// 	fmt.Println("AnotherCase")
	// }
	//*

	// НЕ ИДЕОМАТИЧНО
	if width := 100; width < 100 {
		fmt.Println("width > 100")
	} else {
		fmt.Println("width < 100")
	}
	// Странное правило номер 1: В GO стараются избегать блоков ELSE
	// Идеоматичность
	if height := 100; height > 100 {
		fmt.Println("height > 100")
		return
	}
	fmt.Println("height < 100")
}
