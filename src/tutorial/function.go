package main

import (
	"fmt"
	"time"
)

func getFullName(firstName, lastName string) string {
	return firstName + " "+ lastName
}

func getNameAndAge(firstName, lastName string, year int) (string, int) {
	return firstName + " "+ lastName, time.Now().Year() - year - 1
}

func main() {
	fullName := getFullName("Nayan", "Kumar")
	fmt.Println(fullName)

	name, age := getNameAndAge("Nayan", "Kumar", 1990)
	fmt.Println(name, age)
	fmt.Println(time.Now().Year())
}