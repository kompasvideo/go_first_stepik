package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Print(a, b, c, d)
}

// подаем на вход
// 4 5
// 6
// 7
// получаем вывод: 4 5 6 7
// все отлично считалось, разбивка по пробелам и переносам строк
