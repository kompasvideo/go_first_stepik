package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["James"] = 42
	m["Amy"] = 24

	fmt.Println(m["James"]) // выведет 42
}
