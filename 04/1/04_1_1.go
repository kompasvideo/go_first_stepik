package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		fmt.Println("a больше чем b")
	} else if a == b {
		fmt.Println("a равно b")
	} else {
		fmt.Println("a меньше чем b")
	}
}
