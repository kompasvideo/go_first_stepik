package main

import "fmt"

//type Contact struct {
//	name string
//	age  int
//}

func main() {
	x := Contact{"Андрей", 33}
	x.age = 27
	fmt.Println(x.age)  // выведет 27
	fmt.Println(x.name) // выведет Андрей
}
