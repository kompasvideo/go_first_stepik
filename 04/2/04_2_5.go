package main

import "fmt"

func main() {
	s := 2

	switch {
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
		fallthrough
	case 3:
		fmt.Println("Три")
		fallthrough
	case 4:
		fmt.Println("Четыре")
	case 5:
		fmt.Println("Пять")
	default:
		fmt.Println("Странное число")
	}

}
