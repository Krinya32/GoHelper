package main

import (
	"fmt"
)

func main() {
	//1. Слайсы (они же срезы)
	// Слайс - динамическая обвязка над массивом
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2]   // Слайс инициализируется пустыми квадратными скобками
	fmt.Println("Slice[0,2]:", startSlice) // Создали слайс, основываясь на уже существующем массиве

	//2. Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSlice)
	//3. Изменение элементов среза
	originalArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originalArr[1:4] //Срезы - это набор ссылок на элементы нижележащего массива
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginarArr:", originalArr)
	fmt.Println("firstSlice", firstSlice)

	//4. Один массив и два производных среза
	fSlice := originalArr[:]
	sSlice := originalArr[2:5]

	fmt.Println("Before modification: Arr", originalArr, "fSlice:", fSlice, "sSlice:", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After modification: Arr", originalArr, "fSlice:", fSlice, "sSlice:", sSlice)

	//5. Срез как встроенный тип
	// type slice struct {
	// 	Length int
	// 	Capacity int
	// 	Zeroelement *byte
	// }

	//6. Длина и емкость слайса
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "Length", len(wordsSlice), "Capacity", cap(wordsSlice))
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "Length", len(wordsSlice), "Capacity", cap(wordsSlice))

	// Capacity (cap) или емкость слайса  - это значение показывающее сколько элементов
	//впринципе можно добавить в срез БЕЗ ВЫДЕЛЕНИЯ ДОПОЛНИТЕЛЬНОЙ ПАМЯТИ ПОД НИЖЕЛЕЖАЩИЙ МАССИВ

	// Допустим у нас есть срез на 3 элемента (uнициализировали без явного указания массива)
	// Компилятор при создании среза СНАЧАЛА создал массив ровно на 3 элемента, после этого компилятор вернул адрес, где массив живет
	// Срез запомнил этот адрес и теперь ссылается на него
	// Потом
	// Начинаем деформировать слайс (увеличиваем длину/увеличим количество элементов)
	// Проблема в нижележащем массиве (на котором основан слайс) все 3 ячейки. Что делать?
	// Компилятор ищет в памяти место для массива 3*2 (в общем случае n*2, где n - первоначальный размер)
	// После того как место найдено (в нашем случае найдено место для 6 элементов), в это место копируются
	// старые 3 элемента на свои позиции. На 4 позицию мы добавляем новый элемент
	// После этого компилятор возвращает слайсу новый адрес памяти где находится массив под 6 переменных

	//Емкость всегда будет изменяться как n*2
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%5 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}
	//Важно: после выделения памяти под новый массив, ссылки под старый массив будут перенеcены в новый
	//Пример
	numArr := []int{1, 2, 4}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)

	// 7. Как создавать слайсы наиболее эффективно
	// make() - функция позволяющая более детально создавать срезы
	sl := make([]int, 10, 15)
	// []int - тип коллекции
	// 10 - длина
	// 15 - емкость
	// Cначала инициализируется arr = [15]int
	// Затем делается срез arr[0:10]
	// После чего он возвращается
	fmt.Println(sl)

	// 8. Добавление элементов в срез
	myWords := []string{"one", "two", "three"}
	fmt.Println("My words:", myWords)
	anotherSlive := []string{"four", "five", "six"}
	myWords = append(myWords, anotherSlive...)
	myWords = append(myWords, "seven")
	fmt.Println("My words:", myWords)

	//9. Многомернеый срез
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println(mSlice)

}
