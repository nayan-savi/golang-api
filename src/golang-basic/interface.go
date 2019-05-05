package main

import (
	"fmt"
)

func myFunc(a interface{}) {
	fmt.Printf("Type :%T Value: %v\n", a, a)
}

func main() {

	var myage int
	myage = 25
	myFunc(myage)

	var myfloat float64
	myfloat = 45.566
	myFunc(myfloat)
	
	
}