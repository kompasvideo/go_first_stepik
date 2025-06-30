package main

import "fmt"

func main() {
	x := 2

	switch x {
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
	case 3:
		fmt.Println("Три")
	default:
		fmt.Println(x)
	}
}
