package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Понедельник")
		fallthrough
	case 2:
		fmt.Println("Вторник")
		fallthrough
	case 3:
		fmt.Println("Среда")
		fallthrough
	case 4:
		fmt.Println("Четверг")
		fallthrough
	case 5:
		fmt.Println("Пятница")
	default:
		fmt.Println("Неправильный день")
	}

	/* выведет
	   Среда
	   Четверг
	   Пятница
	*/
}
