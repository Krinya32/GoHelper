package main

import (
	"fmt"
)

func main() {
	//Array Массивы основы
	//1. Определение массива (Структура данных)
	// Массив статичен
	// Создадим массив под хранение 5  целочисленных элементов
	var arr [5]int // При инициализации массива важно передать информацию -сколько элементов в нем будет
	fmt.Println("This is my arr:", arr)
	//2. Определение элеметнов массива (после предварительной инициализации)
	// Необходимо обьратиться к элементу массива через синтаксис arr[i] = elem
	arr[0] = 10
	arr[1] = 20
	arr[3] = -500
	arr[4] = 720
	fmt.Println("After elements init", arr)
	//3. Определение массива с указанием элементов на месте
	// Если при инициализации количество элементов меньше номенальной длины массива то недостающие элементы инициализируются нулями
	newArr := [5]int{10, 20, 30, 40}
	fmt.Println("Short declaration ana init", newArr)
	//4. Создание массива через инициализацию переменных
	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Length:", len(arrWithValues))
	arrWithValues[0] = 1000
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Length:", len(arrWithValues))
	//5. Массив это набор ЗНАЧЕНИЙ. То есть при всех манипуляциях массив копируется (жестко,  на уровне компилятора)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 1000
	fmt.Println("First arr:", first, "Second arr:", second)
	//6. Массив и его размер это две составляющих одного типа. (Размер массива - часть типа)
	// var aArr [5]int
	// var bArr [6]int
	// aArr[0]=100
	// bArr=aArr
	//7. Итерирование по массиву (Перебор элементов)
	floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	}
	//8. Итерирование по массиву через оператор range
	var sum float64
	for id, val := range floatArr {
		fmt.Printf("%d element of arr %.2f\n", id, val)
		sum += val
	}
	fmt.Println("Total sum:", sum)
	//9. Игнорирование id в range based for цикле
	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}
	for id := range floatArr {
		fmt.Printf("%d value WO id\n", id)
	}
	// 10. Многомерные массивы
	words := [2][2]string{
		{"Bob", "Alice"},
		{"Victor", "Jo"},
	}
	fmt.Println("Multidimensional arr:", words)
	//11. Итерирование по многомерному массиву
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}

	//12. Общение со срезом == как массив, только без указания начального среза
	//Массив статичная вещь! Срез в отличии от массива динамический
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10
	slice[1] = 200
	slice = append(slice, 200) // Добавление нового элемента
	for i, v := range slice {
		fmt.Println(i, v)
	}

	emptySlice := []int{}
	emptySlice = append(emptySlice, 200)
}
