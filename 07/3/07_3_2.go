package main

import "fmt"

func main() {
	nums := make([]int, 5)
	nums[1] = 2
	nums[2] = 3
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
