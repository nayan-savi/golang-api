package main

import (
	"fmt"
)

// Vertex type
type Vertex struct {
	Latitude, Longitude float64
}

func main() {
	var data = map[string]Vertex{
		"tumkur": Vertex{Latitude: 13.3379, Longitude: 77.1173},
	}
	fmt.Println(data)

	var places = make(map[string]Vertex)
	places["bangalore"] = Vertex{Latitude: 12.9716, Longitude: 77.5946}

	fmt.Println(places)

}
