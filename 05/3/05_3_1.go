package main

import "fmt"

func square(num int) int {
	return num * num
}

func main() {
	res := square(4)
	fmt.Println(res) // выведет: 16
}
