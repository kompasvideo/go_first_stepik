package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Print(a, b, c, d)
}

// подаем на вход:
// 4 5
// 6
// 7
// получаем вывод: 4 5 0 0
// считалась первая строка в первый и второй аргументы,
// остальные строки не считались
