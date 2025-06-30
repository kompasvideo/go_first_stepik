package main

import "fmt"

func main() {
	a := [5]int{0, 2, 4, 6, 8}
	fmt.Println(a)
	var s []int = a[1:3]
	fmt.Println(s)
	fmt.Println(s[0])
}
