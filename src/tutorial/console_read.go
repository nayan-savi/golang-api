package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	if strings.Contains(text, "hi") {
		fmt.Println("Element is present")
	}

	str := strings.Fields("one two three   four")
	fmt.Println(str[0])
}
