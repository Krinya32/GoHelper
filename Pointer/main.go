package main

import "fmt"

// 1. Указатели - переменная хранящая в памяти значения адрес памяти другой переменной

func main() {

	// 2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable // &... - операция взятия адреса в памяти
	// Выше у нас создан указатель на переменную variable
	// B pointer лежит 18293xcd000132  - это место в памяти где хранится int значение 30
	fmt.Printf("Type of pointer %T\n", pointer)
	fmt.Printf("Value of pointer %v\n", pointer)

	// 3. А какое нулевое значение для указателя
	var zeroPointer *int //zeroValue имеет значение nil (Указатель вникуда)
	fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer)
	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("After initialization type %T and value %v\n", zeroPointer, zeroPointer)
	}

	// 4. Разыменование указателя т.е. получение значения: *pointer - возвращает значение, хранимое по адресу
	var numericValue int = 32
	var pointerToNumeric *int = &numericValue

	fmt.Printf("value is numericValue is %v\nAdress us %v\n", *pointerToNumeric, pointerToNumeric)

	// 5. Cоздание указателей на нулевые значения типов
	// var zeroVar int
	// var zeroPoint *int = &zeroVar
	zeroPoint := new(int) // создает под капотом zeroValue для int и возвращает адрес, где этот 0 хранится
	fmt.Printf("Value is *zeropointer %v\n Adress is %v\n", *zeroPoint, zeroPoint)

	// 6. Изменение значения хранимого по адресу через указатель
	zeroPointerToInt := new(int)
	fmt.Printf("Adress is %v\n and Value in zeroPointToInt is:%v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Adress is %v\n and Value in zeroPointToInt is:%v\n", zeroPointerToInt, *zeroPointerToInt)

}
