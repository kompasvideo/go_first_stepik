package main

import "fmt"

func main() {
	var a [5]int
	b := [5]int{0, 2, 4, 6, 8}
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 8
	a[1] = 42
	fmt.Println(a[1])
}
