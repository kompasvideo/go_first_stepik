package main

import "fmt"

func main() {
	sum := 1
	for i := 5; i <= 8; i++ {
		sum += i
	}
	fmt.Println(sum)
}
