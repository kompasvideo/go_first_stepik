package main

import (
	"fmt"
	"time"
)

func out2(from, to int) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
}
func main() {
	go out2(0, 5)
	go out2(6, 10)
	time.Sleep(5000 * time.Millisecond)
}
