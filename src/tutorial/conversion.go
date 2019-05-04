package main

import (
	"fmt"
)

func main() {
	var men uint8
	men = 5
	var women uint16
	women = 16

	var people int
	people = int(men) + int(women)
	fmt.Println(people)

	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
	// it returns true
	fmt.Println(maxFloat32 + 10)      // 16777216
	fmt.Println(maxFloat32 + 2000000) // 16777216

	var t complex64
	fmt.Println(t)

	var b bool
	fmt.Println(b)

	var myName string
	myName = "Nayan Kumar"
	fmt.Println(myName)

	const meaningOfLife int = 42

	fmt.Println(meaningOfLife)
}
