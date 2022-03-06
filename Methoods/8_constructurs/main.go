package main

import "fmt"

// 1. Создадим структуру Rectangle
type Rectangle struct {
	length, width int
}

// 2. Создадим контструктор для Rectangle
// Для имен контсрукторов в Go договорились использовать функцию с следующим названием
// * New() если данный конструктор на файл один (в файле присутствует описание только одной структуры)
// * NewStructName> - если в файле присутствует еще какие - то структуры

// 3. В Go принято возвращать из конструктора не сам экземпляр, а ссылку на него

func New(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

// 4. Добавим два метода
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := New(10, 20)
	fmt.Printf("Type as %T and value %v\n", r, r)
	fmt.Println("Perimeter", r.Perimeter())
	fmt.Println("Area", r.Area())
}

type Circle struct {
	radius float64
}

func NewCircle(newRadius float64) *Circle {
	return &Circle{newRadius}
}
