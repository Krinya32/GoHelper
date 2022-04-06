package main

import (
	"fmt"
	"time"
)

//1. Запустим несколько горутин и посмотрим, как они бьются за ресурсы

//2. Определяем первую горутину
func printEvenNumbers() {
	for i := 1000; i < 1020; i += 2 {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

//3.  Определяем вторую горутину
func printOddNumbers() {
	for i := 1; i < 20; i += 2 {
		time.Sleep(450 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

//4. Определяем main горутину
func main() {
	go printEvenNumbers()               // сразу идем дальше, запуск фукции будет происходить в отдельной горутине
	go printOddNumbers()                // также идем дальше, запуск функции будет происходить потом
	time.Sleep(5000 * time.Millisecond) // тормозим основную горутину, чтобы остальные успели что-то сделать
	fmt.Println("Main goroutine is died")
}

// 5.  Таким образом горутины работают следующим образом : (нарисуем прямоугольник с палочками!)
