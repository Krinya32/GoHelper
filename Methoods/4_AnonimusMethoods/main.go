package main

import "fmt"

type University struct {
	city string
	name string
}

// 1. Данный метод явно связан только с структурой University
func (u *University) FullInfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s", u.name, u.city)
}

// 2. А в структуру Professor встроены поля структуры University (переходят и все методы тоже)
type Professor struct {
	fullName string
	age      int
	University
}

func main() {
	p := Professor{
		fullName: "Boris Bobrov",
		age:      50,
		University: University{
			city: "Moscow",
			name: "BMSTU",
		},
	}
	// 3. Вызываем метод структуры University через экземпляр профессора
	p.FullInfoUniversity()
}
