package main

import "fmt"

func main() {
	a := 23
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 11
	fmt.Println(*b)
}
