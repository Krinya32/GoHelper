package main

import (
	"fmt"
	"math"
)

// 3. В чем преимущество методов над функциями?
// Во первых: наличие методов улучшает консистентность кода, так как напрямую влияет на его понимание
// Во вторых: в ГО запрещается создавать функции с одинаковыми названиями, а методы разрешены

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

// 4. Создадим функцию и метод Perimeter  для структуры circle
func Perimeter(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

// 5. Создадими функцию и метод Perimeter для структуры Rectangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.width + r.length)
}

// При обьявлении функции Perimeter происходит ошибка
// 6. В Go разрешено создавать методы с одинаковыми именами в пределах одной структуры. Главное чтобы получатель метода вразных структурах где реализован метод со схожим именем отличался

func main() {
	c := Circle{10.5}
	fmt.Println("Call function:", Perimeter(c))
	fmt.Println("Call methood:", c.Perimeter())

	r := Rectangle{10, 20}
	fmt.Println("Call methood Rectangle", r.Perimeter())
}
