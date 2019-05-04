package main

import (
	"fmt"
)

// Person type
// type Person struct {
// 	name string
// 	age  int
// }

func main() {
	// Person type
	type Person struct {
		name string
		age  int
	}
	var person Person
	person.name = "nayan"
	person.age = 27
	fmt.Println(person)

	newPerson := Person{name: "kumar", age: 28}
	newPerson.age = 23
	fmt.Println(newPerson)
}
