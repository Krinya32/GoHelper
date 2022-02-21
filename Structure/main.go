package main

import "fmt"

/// 1.Указатели на массивы. Почему так делать не надо (ИСПОЛЬЗУЙТЕ СЛАЙСЫ)
func mutation(arr *[3]int) {
	// (*arr)[1] = 909
	// (*arr)[2] = 10000
	// Можно написать итак т.к. ГО сам разименует указатель на массив (из-за того, что функция принимает *arr)
	arr[1] = 909
	arr[2] = 10000
}

// 2. Используйте слайсы - это идеоматично с точки зрения ГО
func mutationSlice(sls []int) {
	sls[1] = 909
	sls[2] = 10000
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("arr before mutation", arr)
	mutation(&arr)
	fmt.Println("arr after mutation", arr)

	newArr := [3]int{1, 2, 4}
	fmt.Println("Before mutation Newarr:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("After mutation NewArr:", newArr)
}
