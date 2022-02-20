package main // package name

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean
	var firstBoolean bool
	fmt.Println(firstBoolean)
	//Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT", !aBoolean)
	// fmt.Println(aBoolean + bBoolean) // Логические операции выполнимы только на д логическими типами

	// Numerics. Integers
	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum a+b:", a+b)
	// Вывод типа через %Т форматирования
	fmt.Printf("Type is %T\n", a)
	// Узнаем сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))

	//Эксперимент. При использовании короткого объявления - тип для целого числа - int платформа-зависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	//Эксперимент 2. Используйте явное приведение типов при необходимости если уверены что не произойдет коллиззии
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	//Эксперимент 3. Если проводятся арифметические операции над int b intX, то обязательно нужно использлвать  механизм приведение, т.к. int != int64
	var third64 int64 = 15123123
	var fourInt int = 156234
	fmt.Println(third64 + int64(fourInt))

	// + - * / %

	//Аналогичный образом устроены uint8, uint16, uint32, uint64, uint
	// Numerics. Floot
	//float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of %T and type of %T\n", floatFirst, floatSecond)
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	fmt.Printf("Sum %.3f and Sub %.3f\n", sum, sub)

	// Numerics. Coplex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	//Strings. Строка - это набор БАЙТ. 1 символ кириллицы занимает 2 байт
	name := "Федя"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)
	fmt.Println("Length of:", name, len(name)) // Функция Len() возвращает количество элементов в наборе
	fmt.Println("Amount of chars:", utf8.RuneCountInString(name))
	//Rune - 1 utf-ный символ.

	// Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "asd"
	fmt.Println(strings.Contains(totalString, subString))
	//rune -> alias int32
	var samleRune rune
	var anotherRune rune = 'Q' //Для инициализации руны символьным значением - используйте ''
	var thirdRune rune = 234
	fmt.Println(samleRune)
	fmt.Printf("Rune is char %c\n", samleRune)
	fmt.Printf("Rune is char %c\n", anotherRune) // %c флаг формата для вывода символьных значений
	fmt.Printf("Rune is char %c\n", thirdRune)
	// "A" < "abcd"
	fmt.Println(strings.Compare("abcd", "efgh")) // -1 if first < second, 0 if first == second, 1 if first > second

	var aByte byte // alias int8
	fmt.Println("aByte", aByte)

}
