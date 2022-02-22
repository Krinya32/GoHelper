package main

import "fmt"

// 1. Структура - заименованные набор полей (состояний), определяющий новый тип данных

// 2. Определение структуры(явное определение)
type Student struct {
	firstname string
	lastname  string
	age       int
}

// 3. Если имеется ряд состояний однго типа можно сделать так
type AnitherStudent struct {
	firstname, lastname, groupname string
	age, course                    int
}

// 11. Структура с анонимными полями
type Human struct {
	firstname string
	lastname  string
	string
	int
	bool
}

func PrintStudent(std Student) {
	fmt.Println("+++++++++++")
	fmt.Println("Firstname:", std.firstname)
	fmt.Println("Lastname:", std.lastname)
	fmt.Println("Age:", std.age)
}
func main() {

	// 4. Создание представителей структуры
	stud1 := Student{
		firstname: "Fedya",
		lastname:  "Petrov",
		age:       21,
	}
	PrintStudent(stud1)

	stud2 := Student{"Petya", "Ivanov", 19} // Порядок указания свойств, такой же как в структуре
	PrintStudent(stud2)

	// 5. Что если не все поля структуры пределить
	stud3 := Student{
		firstname: "Vasya",
	}
	PrintStudent(stud3)

	// 6. Анонимные структуры (структура без имени)
	anonStudent := struct {
		age           int
		groupId       int
		proffesorName string
	}{
		age:           23,
		groupId:       2,
		proffesorName: "Alexeev",
	}
	fmt.Println("Anonstudent:", anonStudent)

	// 7. Доступ к состояниям
	studVova := Student{"Vova", "Ivanov", 19}
	fmt.Println("firstname", studVova.firstname)
	fmt.Println("lastname", studVova.lastname)
	fmt.Println("age:", studVova.age)
	studVova.age += 2
	fmt.Println("NewAge:", studVova.age)

	// 8. Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)

	// 9. Указатели на экземпляры структур
	studPointer := &Student{
		firstname: "Igor",
		lastname:  "Sidorov",
		age:       22,
	}
	fmt.Println("Value studPointer", studPointer)
	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value afterpointerModFly", studPointer)
	studPointerNew := new(Student)
	fmt.Println(studPointerNew)
	// 10. Работа с доступом к полям структур через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age)
	fmt.Println("age via .age:", studPointer.age)
	// Неявно происходит разыменование указателя studpointer и запрос соотв. поля

	// 12. Создание экземпляра с анонимными полями структуры
	human := &Human{
		firstname: "Bob",
		lastname:  "Johnson",
		string:    "Additional info",
		int:       -1,
		bool:      true,
	}
	fmt.Println(human)
	fmt.Println("Anon field string", human.string)
}
