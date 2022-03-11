package main

import "fmt"

//1. Дaнная функция будет запущена в качестве горутины.
func newGoRoutine() {
	fmt.Println("Hey, I`m new GoRutine!")
}

//2. Функция main - на самом деле тоже горутина
// Оссобеность в том, что если эта горутина завершается - все остальные запущенные убиваются
func main() {
	go newGoRoutine() // в этот момент происходит формирование запроса на вызов функции в отдельной горутины
	// соответственно код основной горутины main продолжает сразу же выполняться и не ждет завершения newGoRoutine()
	fmt.Println("Main goroutine work!")
	//Запустим код и....
}
