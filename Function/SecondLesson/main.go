package main

import "fmt"

// 1. ЯВНЫЕ ФУНКЦИИ. Любая функция в Go является экземпляром 1 уровня (функцию можно присваивать в переменную, можно передавать в качестве параметра и возвращать из других функций)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

// 2. Возврат функции в качестве значения
func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

// 3. Функция как параметр в другой функции
func receiveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200
	return f(intVarA, intVarB)
}

func main() {
	var command string
	command = "substraction"
	res := calcAndReturnValidFunc(command)
	fmt.Println("Result with command:", command, "value", res(10, 20))
	fmt.Println(res(30, 40))

	// 4. Тип функции
	fmt.Printf("Type of res:%T\n", res)
	fmt.Printf("Type of res:%T\n", calcAndReturnValidFunc)
	// 5. Тип функции в GO определяется как входными параметрами, так и выходными.
	fmt.Println("receiveFuncAndReturnValue", receiveFuncAndReturnValue(add))

	fmt.Println(receiveFuncAndReturnValue(func(a, b int) int {
		return a*a + b*b + 2*a*b
	}))

}
