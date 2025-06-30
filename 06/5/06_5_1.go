package main

import "fmt"

type Contact struct {
	name string
	age  int
}

func (x Contact) welcome() {
	fmt.Println(x.name)
	fmt.Println(x.age)
}

func main() {
	x := Contact{"Андрей", 33}
	p := &x
	fmt.Println(p.age)
	p = &Contact{"Андрей", 33}
	fmt.Println(p.age)
	x.welcome()
	p.welcome()
}
