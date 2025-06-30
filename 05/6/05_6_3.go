package main

import "fmt"

func squareSumm(a int, b int) int {
	var sum int = 0
	//var i int
	for i := a; i <= b; i++ {
		sum += i * i
	}
	return sum
}

func main() {
	fmt.Println(squareSumm(2, 6))
}
