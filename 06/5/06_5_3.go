package main

import "fmt"

type Contact2 struct {
	name string
	age  int
}

func (ptr *Contact2) increase(val int) {
	ptr.age += val
}

func main() {
	x := Contact2{"Андрей", 33}
	x.increase(5)
	fmt.Println(x.age)
}
