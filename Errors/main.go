// В Go существуют два механизма сигнализирования аномального поведения
// 1-й механизм это ошибки Error() - ЯВЛЯЕТСЯ КАНОНИЧНЫМ ИСПОЛНЕНИЕМ НА СЛУЧАЙ НЕНОРМАЛЬНОГО ПОВЕДЕНИЯ
// 2- механизм - это паника (не лучший вариант, т.к. приводит к аварийному завершению, и в принципе мог быть ошибкой)

package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
)

func funcWithError(a int) (string, error) {
	if a%2 == 0 {
		return "Even", nil
	}
	return "", errors.New("THIS IS ERROR")
}

func PanicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Panic satisfied", r)
		debug.PrintStack()
	}
}

func panicExample(firstname *string, lastname *string) {
	defer PanicRecover() // Даже в случае возникновения паники, первым делом будут вызваны все deferred функции
	if firstname == nil {
		panic("RuntimeError: firstname can not be nil")
	}

	if lastname == nil {
		panic("Runtime error: lastname can not be nil")
	}
	fmt.Println("Fullname:", *firstname, *lastname)
}

func main() {
	numLiteral := "12"
	num, err := strconv.Atoi(numLiteral)
	if err != nil {
		log.Fatal("can not convert this value to int", err)
		return
	}
	fmt.Println("Convert done", num)

	ans, err := funcWithError(5)
	if err != nil {
		log.Fatal("not use odd values", err)
		return
	}
	fmt.Println(ans)

	var name = "Bob"
	panicExample(&name, nil)
}
