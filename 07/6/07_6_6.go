package main

import "fmt"

func myPrint(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

func main() {
	myPrint(6, 7, 8)
}
