package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

// 1. Метод, в котором получатель копируется и в его теле происходит работа с локальной копией
func (e Employee) SetName(newName string) {
	e.name = newName
}

// 2. Метод, в котором получатель передается по ссылке ( в теле метода работаем с ссылкой на экземпляр)
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

// 4. Используйте методы с PointerReeiver'ом в ситуациях когда:
// 1) Изменения в работе метода над экземпляром должны быть видны на вызывающей стороне
// 2) Когда экземпляр достаточно "увесистый" т.е. копировать его долго с точки зрения расходов ресурсов
// 3) С ними может работать любой вид экземпляра

func main() {
	e := Employee{"Alex", 3000}
	fmt.Println("Employee before changes", e)
	e.SetName("Bob")
	fmt.Println("After Setname call:", e)
	(&e).SetSalary(4500) // 3. Вызов метода у ссылки на сотрудника
	// 5. Аналогично явному вызову (&e).SetSalary(999)
	fmt.Println("After SetSalary call:", e)
}
