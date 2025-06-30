package main

import "fmt"

func scope() {
	var s = "локальная переменная"
}

func main() {
	fmt.Println(s)
}

// будет выведена ошибка:
// undefined: s
