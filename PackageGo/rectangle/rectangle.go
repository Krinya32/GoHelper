package rectangle // Т.к. пакет main один, тут мы создаем сторонний пакет

type Rectangle struct {
	A, B  int    // Начинается с большой буквы
	color string // Начинается с маленькой буквы
}

func New(newA int, newB int, newColor string) *Rectangle {
	return &Rectangle{newA, newB, newColor}
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.A + r.B)
}
