package main

import "fmt"

func welcome() {
	fmt.Println("Добро пожаловать")
}

func main() {
	defer welcome()
	fmt.Println("Привет!")
}
