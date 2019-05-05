package main

import (
	"fmt"
)

func main() {
	var myint int8
	for i := 0; i < 129; i++ {
		myint++
	}

	fmt.Println(myint)
}
