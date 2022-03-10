package main

import (
	"log"
	"testing"
)

//1. Файл с модульными тестами принято называть
// * <script_name>_test.go
// * <package_name>_test.go

//2. Для того, чтобы тестировать функции (методы, структуры, интерфейсы) принято следующее
// На каждый юнит создаем по своей тестирующей функции (Test)
// Принято такую функцию начинать со слова Test..

func TestAdd(t *testing.T) {
	//1. 1-й тест-case
	if res := Add(10, 20); res != 30 {
		log.Fatalf("Add(10,20) fail, expected %d, got %d\n", 30, res) // log.Fatal провоцирует завершение работы кода
	}
	// 2-й тест-case
	if res := Add(30, 30); res != 60 {
		log.Fatalf("Add(30,30) fail, expected %d, got %d\n", 60, res)
	}
}

func TestSub(t *testing.T) {}

func TestMult(t *testing.T) {}

//3. Для запуска тестов используем команду *go test
// Детально: go test -v

//4. Сoverage (покрытие) - показывает сколько процентов от ычего кода покрыто модульными тестами
// go test -cover -v
// 75-80%  - этого бывает более, чем достаточно!

//5. Все, что начинается со слова test, будет запущено командой go test
//
