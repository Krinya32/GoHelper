package main

import "fmt"

// 0. Почему интерфейс полуабстрактный тип в Go
// 1. Создадим интерфейс ездилка
type Rider interface {
	Ride()
	Gas()
	Stop()
}

func main() {
	// 2. Создаю экземпляр ездилки
	var r Rider   // ZeroValue - nil, ZeroType - nil
	if r == nil { // Попробуем сравнить с nil
		fmt.Printf("r is nil, Value %v and type %T\n", r, r)
	}

	r.Ride() // Здесь будет паника, так как у экземпляра интерфейса,можно вызвать метод Ride()
	// Но так как значение которое там лежит - nil - получается nil.Ride()
	// Мораль - если код падает с memory-dereferacing - это попытка обратится к экземпляру интерфейса без присвоенного претендента.

}
