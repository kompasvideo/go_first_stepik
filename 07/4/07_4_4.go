package main

import "fmt"

func main() {
	x := map[int]int{
		8: 2,
		2: 5,
		4: 10,
		5: 4,
	}
	delete(x, 2)
	fmt.Println(x[4] - x[5])
}
