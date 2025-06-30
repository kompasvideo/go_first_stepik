package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a)
	a = append(a, 8)
	fmt.Println(a) // выведет [0 0 0 0 0 8]	fmt.Println(a)
}
