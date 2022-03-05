package main

import "fmt"

type Rectangle struct {
	length, width int
}

// 1. Реализуем функцию и мтерод для подсчета периметра треугольника
// ВАЖНО: все параметры передаем как копии

// 4. Данную функцию можно вызватть ТОЛЬКО у копии прямоугольника (но не у его ссылки)
func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

// 5. При выове этого метода неважно, будет ли он вызываться у экземпляра или у его ссылки
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// 6. Допустим будет метод который меняет значение состояния структуры на новое, но этот метод ValueReciever  в качестве копии
func (r Rectangle) SetLength(newLength int) {
	r.length = newLength
}

func main() {
	// 2. Cоздаем экземпляр прямоугольника
	rectangleAsValue := Rectangle{10, 20}
	fmt.Println("Call function for RectangleASValue:", Perimeter(rectangleAsValue))
	fmt.Println("Call methood for RectangleAsValue:", rectangleAsValue.Perimeter())

	// 3. Создадим указатель на прямоугольник
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("Call methood for rectangleAsPointer", rectangleAsPointer.Perimeter())
	// Peimeter(rectangleAsPointer) - илюстрация к пункту 5
	// 7. Вызываем метод SetLength
	fmt.Println("Before call methood Setlength", rectangleAsValue)
	rectangleAsValue.SetLength(9999)
	fmt.Println("After call methood SetLength", rectangleAsValue)

	// 8. Вызываем метод SetLength у ссылки на RectangleAsValue
	rectangleAsPointer.SetLength(9999999)
	fmt.Println("After call methood SetLength for &rectangle", *rectangleAsPointer)

}
