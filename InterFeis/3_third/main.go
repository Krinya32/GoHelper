package main

import "fmt"

// 1. А что если создать интерфейс, в котором в впринципе нет никаких требований к поведению?

type Empty interface {
}

// 2. Вопрос - кто удовлетворяет к этому интерфейсу? Если интерфейс пустой, то ему удовлетворяет вообще кто угодно.

// 3. Реализуем функцию, которая может принимать кого угодно
func Describer(pretendent interface{}) {
	fmt.Printf("Type %T and vlue %v\n", pretendent, pretendent)
}

type Student struct {
	name string
}

// 5. Type ASsertion
func SemiGeneric(pretendent interface{}) {
	val, ok := pretendent.(int) // Пытаюсь проверить, что претендент - типа int
	fmt.Println("Value", val, "Ok?", ok)
}

func main() {
	strSample := "Hello World"
	// 4. Передача параметра без присваивания в промежуточную перееменную
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{"Vova"})

	SemiGeneric(10)
	SemiGeneric(Student{"Fedya"})
	SemiGeneric("Hello world")
}
