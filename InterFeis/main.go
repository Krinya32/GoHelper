package main

import "fmt"

// 0 Структура - явнодекларированный заименованный набор СОСТОЯНИЙ.
// Структура исходя из своего описания отвечает на вопрос: из чего я должен состоять
// чтобы считаться с тем ТИПОМ, который декларируется структурой
// Структура описывает ПАТТЕР СОСТОЯНИЯ

// 1. Интерфейсы - это явнодекларированный набор сигнатур ПОВЕДЕНИЙ (чаще всего в виде набора методов), удовлетворив который, можно считаться типом, который декларирует интерфейс
// Интерфейсы в Go декларируют ПОЛУ-АБСТРАКТНЫЙ тип.
// ОТвечает на вопрос - что я должен УМЕТЬ ДЕЛАТЬ, чтобы считаться с тем ТИПОМ, который декларирует интерфейс?
// Интерфейс в контектсе Go - описывает ПАТТЕРН ПОВЕДЕНИЯ

// 2. Объявление интерфейса
type FigureBuilder interface {
	// Набор сигнатур (методов) - которые необхдимо реализовать в труктуре претенденте
	// Во-первых, у претендента должен быть метод Area() возвращающий int
	Area() int
	// Во-вторых, у претендента должен быть Perimeter() возвращающий int
	Perimeter() int
}

// 3. Декларируем претендента
// 3.1 Первый претендент - это прямоугольник
// У него есть два метода - Area() и Perimeter()
// Когда эти методы реализованы, говорят,что RECTANGLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕЙСА FigureBuilder
// RECTANGLЕ РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// 2. Второй претендент - это окружность
type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {
	// 4. Создадим 3 экземпляра
	r1 := Rectangle{10, 20}
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1, 1}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}

	// 5. Посчитаем общую площадь этих двух фигур
	rectangles := []Rectangle{r1, r2, r3}
	totalSumAreaOfRectangler := 0
	for _, rect := range rectangles {
		totalSumAreaOfRectangler += rect.Area()
	}

	circles := []Circle{c1, c2, c3}
	totalSumAreaOfCircle := 0
	for _, circ := range circles {
		totalSumAreaOfCircle += circ.Area()
	}

	fmt.Println("Total area is:", totalSumAreaOfRectangler+totalSumAreaOfCircle)

	// 6. Теперь воспользуемся интерфейсом указанным выше
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3} // Объявляю слайс экземпляров, удовлетворяющих интерфейсу FigureBuider
	// Сдругой стороны, кажется, что это слайс каких то определенных типов
	// 7. Посчитаем общую площадь
	total := 0
	for _, fig := range figures {
		total += fig.Area()
	}
	fmt.Println("Total area:", total)
	// 8. Пояснение - так как каждый экземпляр удовлетворяет интерфейсу figureBuilder (объявляющий полу-абстрактный тип)
	// у каждого из слайса figures можно 100% вызвать метод Area() (который вернет int), или Perimeter() который тоже 100% вернет int
}
