package main

import "fmt"

func main() {
	x := 12
	res := 0
	switch {
	case x > 12:
		res += 1
	case x > 5 && x < 15:
		res += 2
		fallthrough
	case x > 12:
		res += 3
	case x < 20:
		res += 4
	default:
		res = 13
	}
	fmt.Println(res)
}
