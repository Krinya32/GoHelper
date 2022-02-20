package main

import (
	"fmt"
)

func main() {
	// 1. Map -  это набор пар ключ:значение. Инициализация пустого мапа
	mapper := make(map[string]int)
	fmt.Println("Empty map:", mapper)

	// 2. Добавление пар в существвубщую mapy
	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Mapper after adding pairs:", mapper)

	// 3. Инициализация мапы с указанием пар
	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1000,
	}
	newMapper["Jo"] = 3000
	fmt.Println("New mapper:", newMapper)

	// 4. Что может быть ключем в маре
	// 4.1 ВАЖНО: Мапа НЕ УПОРЯДОЧЕНА В GO
	// 4.2 КЛЮЧОМ в мапе может быть только сравнимый тип (==, !=)

	// 5. Нулевое значение для map
	// var mapZeroValue map[string]int // mapZeroValue == nil
	// mapZeroValue["Alice"] = 12

	// 6. Получение элементов из мап
	// 6.1 Получение элемента который представлен в мапе
	testPerson := "Alice"
	fmt.Println("Salary of testPerson:", newMapper[testPerson])

	// 6.2 Получение элемента который не представлен в мапе
	testPerson = "Derek"
	fmt.Println("Salary of new testPerson:", newMapper[testPerson]) // при обращении по несуществующему ключу -  новая пара не добавляется
	fmt.Println(newMapper)

	// 7. Проверка вхождения ключа

	employee := map[string]int{
		"Den":   0,
		"Alice": 0,
		"Bob":   0,
	}
	fmt.Println("Den and value:", employee["Den"])
	fmt.Println("Jo and value:", employee["Jo"])
	// 7.1 При обращении по ключу -  возвращается два значения
	if value, ok := employee["Den"]; ok {
		fmt.Println("Den and value:", value)
	} else {
		fmt.Println("Den does not exist in map")
	}

	if value, ok := employee["Jo"]; ok {
		fmt.Println("Jo and value:", value)
	} else {
		fmt.Println("Jo does not exist in map")
	}

	// 8. Перебор элементов мапы
	fmt.Println("=================")
	for key, value := range employee {
		fmt.Printf("%s and value %d\n", key, value)
	}
	// 9. Как удалять пары
	// 9.1 Удаление существующей пары
	fmt.Println("Before deleting:", employee)
	delete(employee, "Den")
	fmt.Println("After deleting:", employee)

	// 9.2  Удаление несуществующей пары
	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna") // ОЧЕНЬ ДОРОГАЯ ОПЕРАЦИЯ
	}

	fmt.Println("After second deleting:", employee)

	// 10. Количество пар == длина мапы
	fmt.Println("Pair amount in map", len(employee))

	// 11. Мапа как и слайс (сслылочный тип)
	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}

	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Println("words map:", words)
	fmt.Println("newWords map:", newWords)

	// 12. Сравнение мапов
	// 12.1 Cравнение массивов (массив можно использовать в качестве ключа к мапе)
	if [3]int{1, 2, 3} == [3]int{3, 4, 5} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	// 12.2 Сравнение слайсов. Слайсы из за того что тип ссылочный, можно сравнить на равенство,только с nil
	// if []int{1,2,3} == []int{1,2,3} {

	// }

	// 12.3 Сравнение мап. Мапа из за того , что тип ссылочный можно сравнивать только с nil
	aMap := map[string]int{
		"a": 1,
	}
	var bMap map[string]int

	if aMap == nil {
		fmt.Println("Zero value map")
	}

	if bMap == nil {
		fmt.Println("Zero value of map bMap")
	}

	// 13. Грустное следствие
	// Если мапа/слайс являются составляющей одной структуры, структура автоматически несравнима
}
