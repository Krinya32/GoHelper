package main

import "fmt"

//1. Defer - оператор отложенного вызова функции (метода)

//2. Создадим отложенную функцию.

//3. С входными параметрами
func CheckDBCloseConnection(a int) {
	fmt.Println("Check Started...... Value numIp is deffered:", a)
	fmt.Println("Check done Connection refused!!!!")
}

//4. Если функция принимает входные параметры и данная функция используется в связке с defer то:
// Параметры расчитываются в момент передачи их в функцию
// А вызов функции с уже давно рассчитанным параметром осуществляетсяв момент завершения вышележащей функции

// 5. В какой момент defer вызывается

func OuterFunc() int {
	defer fmt.Println("I`m deferred function!")
	fmt.Println("Outerfunc started")
	var result = 10
	fmt.Println("Outerfunc finished. Ready to return value!")
	return result
}

func main() {
	defer fmt.Println("Step 1 defer")
	defer fmt.Println("Step 2 defer")
	defer fmt.Println("Step 3 defer")
	defer fmt.Println("Step 4 defer")
	defer fmt.Println("Step 5 defer")

	numIp := 10
	p := &numIp
	//defer CheckDBCloseConnection(numIp) // Defer означает, что данная функция будет вызвана при завершении main( с параметром, значение которого рассчитывает на момент 21 строки)
	*p++
	fmt.Println("Value num ip in main()", numIp)
	fmt.Println("Main function started")
	fmt.Println("Main function Ended")
	fmt.Println("Value from OuterFunc on main side is", OuterFunc())
}
