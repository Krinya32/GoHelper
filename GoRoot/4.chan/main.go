package main

import "fmt"

//1. Каналы - средство для коммуникации между горутинами
// Каналы можно рассматривать, как соединительные трубы, через которые горутины между собой общаются (аналогично так, как вода течет по трубe, данные передаются через каналы)

//2. Объявление канала
// Каналы по умолчанию имеют zeroValue - nil. Поэтому их создаем через функцию make.

func main() {
	var a chan int // объявляем канал, через который будут передаваться данные типа int
	if a == nil {
		fmt.Println("channel is nil, Let`s define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
}
