package main

import "fmt"

// 1. Методы для стандартных типов
// в Go встроено куча приметивов (int,int32,string,bool)
// Что если очень хочется дописать к стандартному типу какой то метод

// 2. Наивная попытка. Это не выполнимо, Компилятор запрещает добавление новых методов к существующим базовым типам
// func (a int) IsEven() bool {
// 	if a%2 == 0 {
// 		return true
// 	}
// 	return false
// }

// 3.  Но мне очень хочется, что делать
// Создайте новый тип - ваш int и делайте с ним что хотите.
// Для создание нового типа используют контсрукцию
type MySuperDuperInt int

func (msdi *MySuperDuperInt) IsEven() bool {
	if *msdi%2 == 0 {
		return true
	}
	return false
}

func main() {
	num1 := MySuperDuperInt(202)
	num2 := MySuperDuperInt(1)
	fmt.Println(num1.IsEven())
	fmt.Println(num2.IsEven())
	// 4. Внутренние преобразования
	// var num3 MySuperDuperInt = MySuperDuperInt(10)
	// var num4 int = int(num3)
}
