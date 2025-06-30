package main

import "fmt"

func two(a int) (int, int) {
	return a + 5, a - 5
}

func main() {
	x, y := two(7)
	fmt.Println(x * y)
}
