package main

import "fmt"

//1. Для отправки данных в канал а (chan int) используют синтаксис
// a <- dataint

//2. Для получения данных из канала используем синтаксис
// dataint := <-a

//3. Отправка и получения данных из канала - блокирующая операция!
// Это означает, что если данные отправлены в канал, то выполнение текущей программы останавливается до тех пор, пока стороны из этого канала кто то не считает данные.

//Аналогично и в обратную сторону. Если кто-то считает из канала, то выполнение текущей программы (горутины) останавливается до тех пор, пока кто-то в этот канал не отправит данные

//4. Пример использования каналов

func newGoRoutine(done chan bool) {
	fmt.Println("Hey, I`m Gorutine!")
	done <- true
}

func main() {
	done := make(chan bool)
	go newGoRoutine(done)
	<-done // В этой точке выполнение main горутины останавливается до тех пор, пока в канал кто-нибудь не запишет данные!
	fmt.Println("Main goroutine work!")
}
