package main

import "fmt"

var s = "глобальная переменная"

func scope2() {
	fmt.Println(s)
}

func main() {
	fmt.Println(s)
	scope2()
}
