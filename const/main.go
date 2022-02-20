package main

import (
	"fmt"
	"math"
)

//I. Константы неизменяемые переменные, которые служат:
// 1) для более строгого понимания кода,
// 2)для того чтобы случайно не поменять значение
// 3) Для удобных преобразований

const (
	MAIN_PORT string = "8001"
)

func main() {
	//2. Объявление одной контсанты
	const a = 10
	fmt.Println(a)
	//3. Объявление блока констант c областью видимости внутри функции main
	const (
		ipAdress = "127.127.00.03"
		port     = "8000"
		dbName   = "postgres"
	)
	fmt.Println("ip", ipAdress)
	fmt.Println(checkPortIsValid())

	//4. Константу нельзя поменять в ходе работы проограммы
	//5. Значение констант должны быть известны на момент компиляции
	var sqrt = math.Sqrt(25)
	// const sqrt = math.Sqrt(25) // Нельзя присовить константу что-лиюо, что является реузьтатом вызова функции, метода
	fmt.Println("Var sqrt:", sqrt)

	//6. Типизированные и нетипизированные константы
	const ADMIN_EMAIL string = "admin@admin.com" // Указание типа - повышение читабельности кода

	//7. Нетипизированные константы
	// При использовании нетипизированных констант, мы разрешаем компилятору
	// использовать явное приведение типов в момент присвоение значений констант в переменные
	const NUMERIC = 10
	var numInt8 int8 = NUMERIC
	var numInt32 int32 = NUMERIC
	var numInt64 int64 = NUMERIC
	var numFloat32 float32 = NUMERIC
	var numComplex complex64 = NUMERIC

	fmt.Printf("numint8 value %v type %T\n", numInt8, numInt8)
	fmt.Printf("%v+%v is %v\n", numInt8, NUMERIC, numInt8+NUMERIC)
	fmt.Printf("numint32 value %v type %T\n", numInt32, numInt32)
	fmt.Printf("numint64 value %v type %T\n", numInt64, numInt64)
	fmt.Printf("numFloat32 value %v type %T\n", numFloat32, numFloat32)
	fmt.Printf("numComplex value %v type %T\n", numComplex, numComplex)
	//8. Константы в GO зашиваются в момент компиляции RUNTIME программы и не выбрасываются до ее окончания
}

func checkPortIsValid() bool {
	if MAIN_PORT == "8002" {
		return true
	}
	return false
}
