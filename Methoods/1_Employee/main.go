package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

// 1. Методы - функции привязанные к определнным структурам
// func (s struct) MethodName(parameters type) type {}
// Reciever - получатель метода

func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Salary: %d %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{
		name:     "Bob",
		position: "Senior Golang developer",
		salary:   3000,
		currency: "USD",
	}
	// 2. Вызов метода
	emp.DisplayInfo()

}
