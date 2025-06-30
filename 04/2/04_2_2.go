package main

import "fmt"

func main() {
	month := "Январь"

	switch {
	case month == "Декабрь" || month == "Январь" || month == "Февраль":
		fmt.Println("Зима")
	case month == "Март" || month == "Апрель" || month == "Май":
		fmt.Println("Весна")
	}
}
