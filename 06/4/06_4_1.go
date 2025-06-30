package main

import "fmt"

type Contact struct {
	name string
	age  int
}

func main() {
	x := Contact{"Андрей", 33}
	p := &x
	fmt.Println(p.age)
	p = &Contact{"Андрей", 33}
}
