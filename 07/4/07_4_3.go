package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 42,
		"Amy":   24}

	delete(m, "James")
	m["Bob"] = 8
	fmt.Println(m["James"]) // выведет 0
	fmt.Println(m["Bob"])   // выведет 0
}
