package main

import (
	"fmt"
	"strings"
)

func main() {
	str := strings.Fields("one two three   four")
	fmt.Println(str)

	fmt.Println(strings.Count("nayan", "a"))
	fmt.Println(strings.Count("nayan", ""))

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}
