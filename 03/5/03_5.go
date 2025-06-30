package main

import "fmt"

func main() {
	a := 7
	b := a * 2
	b /= 4
	a += 3
	fmt.Println(5 <= 7 && 5 != 7)
	fmt.Println(7 >= 4 || 5 == 9)
	fmt.Println(!(7 >= 4) && 5 == 5)
	fmt.Println(!(11 > 3 || 7 != 4))
	fmt.Println(10 > 7 && 5 < 6 || 5 == 6)
}
