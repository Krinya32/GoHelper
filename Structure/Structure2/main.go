package main

import "fmt"

// 1. Вложенные структуры. Это использование одной структуры, как тип поля в другой структуре
type Univercity struct {
	yearBased int
	infoshort string
	infolong  string
}

type Student struct {
	firstname  string
	lastname   string
	university Univercity
}

// 4. Встроенные структуры (когда мы добавляем поля одной  структуры к другой)

type Professor struct {
	firstname string
	lastname  string
	age       int
	greatwork string
	// papers     []string  - добавление этого поля делает структуру несравнимой (мапа, слайс)
	Univercity // В этом месте происходит добавление всех полей структуры Uni в Prоfessor
}

func main() {
	// 2. Создание экземпляров структур с вложением
	stud := Student{
		firstname: "Fedya",
		lastname:  "Petrov",
		university: Univercity{
			yearBased: 1994,
			infoshort: "Cool Uiversity",
			infolong:  "Very cool University",
		},
	}

	// 3. Получение доступа к вложенным полям структур

	fmt.Println("firstname:", stud.firstname)
	fmt.Println("Lastname:", stud.lastname)
	fmt.Println("Year Based Uni:", stud.university.yearBased)
	fmt.Println("Longinfo:", stud.university.infolong)
	// 5. Создание экземпляра со встраиванием структур
	prof := Professor{
		firstname: "Anatoly",
		lastname:  "Smirnov",
		age:       125,
		greatwork: "Ultimate C programming",
		Univercity: Univercity{
			yearBased: 1963,
			infoshort: "short info",
		},
	}
	// 6. Обращение к полям или состоянием со встроенной стртуктурой
	fmt.Println("firstname:", prof.firstname)
	fmt.Println("yearbased", prof.yearBased)
	fmt.Println("infoshort", prof.infoshort)

	// 7. Сравнение экземпляров ==
	// При сравнении экземпляров происходит сравнение всех их полей друг с другом
	profLeft := Professor{}
	profRight := Professor{}

	fmt.Println(profLeft == profRight)
	// 8. Если ихотябы одно из полей структур несравнимо, то и вся структура несравнима
}
