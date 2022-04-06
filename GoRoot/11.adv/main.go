package main

import "fmt"

//1. Попробуем решить задачу с подсчетом суммы квадратов и кубов полученного числа с использование закрытия каналов

func digits(number int, dchn1 chan int) {
	for number != 0 {
		digit := number % 10
		dchn1 <- digit
		number /= 10
	}
	close(dchn1)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch) // Своя горутина для генерации цифр
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
	close(squareop)
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch) // Своя горутина для генерации цифр
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
	close(cubeop)
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
