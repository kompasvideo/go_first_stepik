package main

import "fmt"

func sum1(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func main() {
	s := []int{42, 33, 22, 11}
	sum1(s...)
}
