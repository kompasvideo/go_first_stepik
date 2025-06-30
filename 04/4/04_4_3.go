package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 30; i += 5 {
		sum += i
	}
	fmt.Println(sum)
}
