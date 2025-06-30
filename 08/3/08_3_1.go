package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch <- 8
	value := <-ch
	fmt.Println(value)
}
