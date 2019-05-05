package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	go compute(10)
	go compute(10)

	go func() {
        fmt.Println("Executing my Concurrent anonymouse function")
    }()
    fmt.Scanln()
}
