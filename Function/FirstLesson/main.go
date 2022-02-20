package main // package name
import "fmt"

// 1.  Явная функция, это локально определенныйц блок кода, имеющий имя (ЯВНОЕ ОПРЕДЕЛЕНИЕ)
//  Функцию необходимо определить. Функцию необходимо вызвать
// 2. Сигнатура функции и их определение
// func functionName(Params type) typeReturnValue {
// 	//body
// }

func main() {
	fmt.Println("Hello world")
	// 4. Вызов простейшей функции
	res := add(10, 20)
	fmt.Println("Result:", res)
	fmt.Println("Result of mult:", mult(1, 2, 3, 4))
	newPer, _ := rectangleParametres(10, 10) // высчитать один периметр, без площади
	per, area := rectangleParametres(10.5, 10.5)
	fmt.Println("Area", area)
	fmt.Println("Per", per)
	fmt.Println("NewPer", newPer)
	secondPer, secondArea := namedReturn(10, 20)
	fmt.Println(secondArea, secondPer)
	emptyReturn(10)

	helloVariadic(10, 20, 30, 40, 50, 60, 60)
	helloVariadic(10, 20)
	helloVariadic(10)
	helloVariadic()

	someStrings(2, 10, "Bob", "Alex", "Vita")

	sum1 := sumVariadic(10, 20, 40, 50, 60)
	sliceNumber := []int{10, 20, 30}
	sum2 := sumVariadic(sliceNumber...)
	fmt.Println(sum1, sum2)

	// 12. Анонимная функция
	anon := func(a, b int) int {
		return a + b
	}

	fmt.Println("Anon:", anon(20, 30))
	fmt.Println("Bigfunction:", bigFunction(10, 20))

}

// 13. Анонимная функция внутри явной
func bigFunction(aArg, bArg int) int {
	return func(a, b int) int { return a + b + 1 }(aArg, bArg)
}

// 3. Простейшая функция (определить функцию можно как до момекнта его вызова функции main, так и в любом месте пакета, главное чтобы она была определена где нибудь)
func add(a int, b int) int {
	result := a + b
	return result
}

// 4. Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

// 5. Возврат больше чем одного значения (returnType1, returnType2.....)
func rectangleParametres(length, width float64) (float64, float64) {
	var perimeter = 2 * (length + width)
	var area = length * width
	return perimeter, area
}

// 6. Именованный возврат значений
func namedReturn(a, b int) (perimeter int32, area int) {
	perimeter = int32(2 * (a + b))
	area = a * b
	return // Не нужно указывать воззвращаемые переменные
}

// 7. При вызове оператора return функция прекращает свое выполнение и возвращает что-то

func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}

	if a == b {
		return a, true
	}

	return 0, false
}

// 8. Что вернется в случае если return в принципе не указан
func emptyReturn(a int) {
	fmt.Println(a)
}

// 9. Valiadic parameters (Континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Printf("Value %v and type:%T\n", a, a)
}

// 10. Смешение параметров с variadic
// 10.1 Континуальный параметр всегда самый последний
// 10.2 Variadic parametres на всю функцию один (для удобочитаемости)
func someStrings(a, b int, words ...string) {
	fmt.Println("Parameter:", a)
	fmt.Println("Parameter:", b)
	var result string
	for _, word := range words {
		result += word
	}
	fmt.Println("Result concat:", result)
}

// 11. Передача слайса или использование Variadic parametres?
func sumVariadic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}
