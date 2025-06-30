package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		if i%5 == 0 {
			break // если число кратно 5, то выходим из цикла
		}
		fmt.Println(i)
	}
	/* выведет:
	   2
	   2
	   3
	   4
	*/
}
