package main

import "fmt"

func test(v ...int) {
	res := 3
	for _, x := range v {
		res += x
	}
	fmt.Println(res)
}

func main() {
	v := []int{2, 4, 5}
	test(v...)
}
