package main

import "fmt"

func sum(s ...string) {
	sums := 0
	for _, v := range s {
		switch v {
		case "w":
			sums += 3
		case "d":
			sums++
		}
	}
	fmt.Println(sums)
}

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	res := results[0:]
	var s string
	fmt.Scan(&s)
	res = append(results, s)
	sum(res...)
}
