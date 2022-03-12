package main

import "fmt"

//1. Создадим чуть более полезную программу, которая будет  делать следующие действия:
// Берем число, например 456
// Подсчитаем сумму квадратов цифр и сумму кубов, а затем сложим полученные результаты
// Делать будем следующим образом:
// * main gorutine получает число и вызывает 2 другие горутины, по итогу, получив от них результаты, она просто их сложит и выведет в консоль
// ** squareGoRoutine - будет запущена main и подсчитает сумму квадратов всех цифр, результат положит в канал
// ** cubesGoRoutine - будет запущена main и подсчитает сумму кубов всех цифр, результат положит в канал

func squareGoRoutine(number int, squareChan chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareChan <- sum
}

func cubesGoRoutine(number int, cubeChan chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10 //456%10 -> 6 // 45%10 -> 5 // ...
		sum += digit * digit * digit
		number /= 10
	}
	cubeChan <- sum
}

func main() {
	number := 456
	squareChan := make(chan int)
	cubeChan := make(chan int)
	go squareGoRoutine(number, squareChan)
	go cubesGoRoutine(number, cubeChan)
	squareSum, cubeSum := <-squareChan, <-cubeChan
	fmt.Println("Total result is:", squareSum+cubeSum)
}
