package main

import "fmt"

func DoSomething(pretendent interface{}) {
	switch pretendent.(type) { // Пытаемся извлечь нижележащий кейс
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	default:
		fmt.Println("Unknown type! But it is working!")
	}
}

func main() {
	DoSomething(10)
	DoSomething("Hello World!")
	DoSomething(func(a, b int) int { return a + b })
}
