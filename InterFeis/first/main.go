package main

import (
	"fmt"
	"unicode/utf8"
)

// 1. описание интерфейса( описание того, что должен уметь претендент)
type BigWord interface {
	IsBig() bool
}

// 2. Наш претендент, которого нажо научить делать IsBig() bool
type MySuperString string

// 3. Реализация IsBig() у претендента MySuperString
func (mss MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MySuperString("asdqwwgrsvtrhrbdf")
	var interFaceSample BigWord // Обьявление переменной, типа интерфейса BigWord
	interFaceSample = sample    // Присваивание значение для переменной возможно, потому что для Sample / удовлетворяет интерфейсу
	fmt.Println("IsBig", interFaceSample.IsBig())

	// 4. Попытка присвоить значение с типажом, неудовлетворяющему интерфейсу

	// var interFaceBadSAmple BigWord
	// interFaceBadSAmple = "wqe13123" // тип string не имеет реализации метода IsBig, поэтому не удовлетворяет интерфейсу
}
