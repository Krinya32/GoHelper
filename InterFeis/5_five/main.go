package main

import "fmt"

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("%s and %d y.o.\n", std.name, std.age)
}

func TypeFinder(i interface{}) {
	switch v := i.(type) { //Присвоим переменнной v значение лежащее под предполагаемым интерфейсом
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case Describer: // Вывод - с типом интерфейса, можно сравниваться точно также, как и с другим типажом языка
		// Это как раз и говорит о том что интерфейсы полуабстрактные типы
		fmt.Println("I am describer!")
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	std := Student{"Alex", 15}
	TypeFinder(10)
	TypeFinder("hello")
	TypeFinder(nil)
	TypeFinder(std)
}
