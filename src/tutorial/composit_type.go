package main

import (
	"fmt"
) 

func main() {
	// array
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	fmt.Println(days)

	var fruits [2]string
	fruits[0] = "apple"
	fruits[1] = "mango"
	fmt.Println(fruits)

	// slice
	weekdays := days[0:5]
	fmt.Println(weekdays)
}
