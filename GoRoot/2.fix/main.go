package main

import (
	"fmt"
	"time"
)

func newGoRoutine() {
	fmt.Println("Hey, I`m new GoRutine!")
}

func main() {
	go newGoRoutine()
	time.Sleep(1 * time.Second) // Немного тормозим выполнение main горутины, таким образом даем время для того, чтобы newGoRoutine успела выполниться
	fmt.Println("Main goroutine work!")
}
