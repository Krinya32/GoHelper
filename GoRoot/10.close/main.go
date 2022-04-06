package main

import "fmt"

//1. Закрытие каналов и итерирование
// Со стороны получателя можно использовать синтаксис
// val, ok - значение помещенное в канал отправителя
// ok - true/false в зависимости от того открыт канал или уже закрыт отправителя
// Если канал закрыт, то в vаl будет помещено zeroValue для типа канала

func generator(ch chan int) {
	for i := 0; i < 25; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go generator(ch)
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Recieved from channel", val)
	}

	// Конструкцию можно упростить и использовать
	// for val:= range ch{
	// fmt.Println("Recieved from channel", val)}
}
