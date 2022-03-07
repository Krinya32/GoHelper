package main

import "fmt"

// 1. Вопрос - имеет ли для интерфейса значение, тот факт, что метод, реализованный для претендента в качестве получателя принимает значение и указатель

// 0. В Go принято называть интерфейсы с окончанием на "er"
type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Println("Student name:", std.name, "Student age:", std.age)
}

type Professor struct {
	name string
	age  int
}

func (prof *Professor) Describe() {
	fmt.Printf("Professor name: %s and professor age: %d\n", prof.name, prof.age)
}

func main() {
	var descr1 Describer
	stud1 := Student{"Vasya", 16}
	descr1 = stud1 // Student реализует интерфейс describer
	descr1.Describe()

	stud2 := &Student{"Bob", 23} // Поскольку экземпляр ссылка, разименовать его имеет право кто угодно (в том числе и компилятор)
	descr1 = stud2
	descr1.Describe()

	var descr2 Describer
	prof1 := &Professor{"Viktor", 65}
	descr2 = prof1 // Профессор реализует интерфейс Describer
	descr2.Describe()

	prof2 := Professor{"Bob", 43}
	prof2.Describe() // Здесь ссылку на &prof берет компилятор
	// Proffessor не реализует интерфейс Describer.
	// Дело в том, что интерфейс не референсный тип
	// Выливается это в следующее следствие:
	// Когда мы работаем с обычным методом, то взять референс на экземпляр может компилятор, но когда мы пытаемся сделать это через интерфейс, в нем компилятор не видит никаких ссылок
}
