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

	b := 345
	a := &b
	c := &b
	*c += 100
	*a++
	fmt.Println(b)
	// 7. Указательная арифметика ОТСУТСТСВУЕТ ПОЛНОСТЬЮ
	// Адрес одной ячейки - вы можете через этот адрес продвинуться в другие ячейки
	// cpp_like := "cpp"
	// cpp_pointer := &cpp_like
	// cpp_pointer++

	// 8. Передача указателей в функции
	// Колосальный прирост производительности,, за счет того что передается не значение, которое должно копироваться, а  передается лишь адрес в памяти, за которым уже хранитья значение
	sample := 1
	// samplePointer := &sample
	fmt.Println("Origin value of sample", sample)
	changeParam(&sample)

	fmt.Println("After sample changing is", sample)

	// 9. Возврат поинтера из функции (В С++ результат такого механизма - не определен впринципе)
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("ptr1: %T and adress %v and value %v\nptr2: %T and adress %v and value %v", ptr1, ptr1, *ptr1, ptr2, ptr2, *ptr2)
}

// 9.1 Инициализация функции вовзращающей указатель
func returnPointer() *int {
	var numeric int = 321
	return &numeric // В момент вовзрата GO перемещает данную переменную в кучу
}

// 8.1 Определение функции принимающей параметр, как указатель
func changeParam(val *int) {
	*val += 100
}
