package main

import "fmt"

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	ch <- result
}
func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i * i
		}
	}
	ch <- result
}

func main() {
	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)

	select {
	case x := <-evenCh:
		fmt.Println(" 1 - ")
		fmt.Println(x)
		fmt.Println()
	case y := <-sqCh:
		fmt.Println(" 2 - ")
		fmt.Println(y)
		fmt.Println()
	}
}
