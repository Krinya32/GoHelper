package main

import "fmt"

// 1. Обьявляем интерфейс декларирующий поведенческий паттерн (набор методов под реализацию)
type Worker interface {
	Work()
}

// 2. Данная структура претендент на удовлетворение интерфейса Worker

type Employee struct {
	name string
	age  int
}

// 3. Реализуем метод Work(), чтобы структура Employee удовлетворяла интерфейсу

func (e Employee) Work() {
	fmt.Println("Now Employee with name", e.name, "is working!")
}

// 4. А давайте создадим функцию, которая будет принимать аргументы типа Worker и что то с ними делать?

func Describer(w Worker) {
	fmt.Printf("Interface with type %T and Value %v\n", w, w)
}

// 6. Объявление структуры. Данная структура - второй претендентна удовлетворение интерфейса Worker

type Student struct {
	name         string
	courseNumber int
}

func (s Student) WantToEat() {
	fmt.Println("Student with name", s.name, "want to eat")
}
func (s Student) Work() {
	fmt.Println("Student with name", s.name, "is working")
}

func main() {
	// 5. Создаем экземпляр Employee
	emp := Employee{"Bob", 34}
	var workerEmployee Worker = emp // Присваиваем сотрудника в переменную типа Worker
	workerEmployee.Work()
	Describer(workerEmployee) // В результате видим, что тип интерфейса определяется структурой, его реализующей, а значение интерфейса - это соответственно значение состояний структуры

	// 7. Создадим экземпляр Student
	std := Student{"Mike", 19}
	var workerStudent Worker = std
	workerStudent.Work()
	Describer(workerStudent) // Принял внутренний тип Student, а значение - равно значению полей экземпляра

	// 8. Создадим набор тех, кто умеет Work()
	var workers []Worker = []Worker{workerStudent, workerEmployee}
	for _, worker := range workers {
		Describer(worker) // Данная функция вызывается у разных экземпляров, благодаря тому, что для ее вызова, экземпляру нужно удовлетворить некоторому контракту, если структура экз. поддерживает данных паттерн, то у экз. можно вызвать все необходимые для этого методы
	}

}
