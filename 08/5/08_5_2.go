package main

import "fmt"

func f2(ch chan int) {
	ch <- 4
}

func main() {
	ch := make(chan int)
	go f2(ch)
	fmt.Println(<-ch / 2)
}
