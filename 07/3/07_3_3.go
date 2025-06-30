package main

import "fmt"

func main() {
	x := "hello"
	for _, c := range x {
		fmt.Println(c)
	}

	// выведет
	// 104
	// 101
	// 108
	// 108
	// 111

	for _, c := range x {
		fmt.Printf("%c \n", c)
	}

	// выведет
	// h
	// e
	// l
	// l
	// o
}
