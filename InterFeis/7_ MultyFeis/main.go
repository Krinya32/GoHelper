// 0. Интерфейсы ( с точки зрения ООП) - увеличивают уровень абстракции вашего кода
// За счет увеличения уровня абстракции - можно решать много сторонних проблем, связанных с поддержкой
// понимане и реюзабельность кода

//1. Что делать, если хочется скрестить 2 интерфейса и создать единый уровень абстракции в коде?
package main

import "fmt"

type PerimeterCalculater interface {
	Perimeter() int
}

type AreaCalculate interface {
	Area() int
}

// 2. Соберем новый интерфейс из двух старых
type FigureFeatureCalculator interface {
	PerimeterCalculater // Встраиваем интерфейсы
	AreaCalculate
}

type Rectangle struct {
	a, b  int
	color string
}

// 2. Реализуем интерфейс PerimeterCalculate
func (r Rectangle) Perimeter() int {
	return 2 * (r.a + r.b)
}

// 3. Реализуем второй интерфейс
func (r Rectangle) Area() int {
	return r.a * r.b
}

func main() {
	var pCalc PerimeterCalculater
	var aCalc AreaCalculate
	rect := Rectangle{10, 20, "green"}
	pCalc = rect // Структура Rectangle удовлетворяет интерфейсу
	pCalc.Perimeter()
	aCalc = rect // Структура Rectangle удовлетворяет интерфейсу
	fmt.Println("Perimeter:", pCalc.Perimeter())
	fmt.Println("Area:", aCalc.Area())

	var ffCalc FigureFeatureCalculator
	ffCalc = rect // Структура удовлетворяет интерфейсу figurefeatureCalculator
	fmt.Println("Figure Area", ffCalc.Area())
	fmt.Println("Figure Perimeter", ffCalc.Perimeter())
}
