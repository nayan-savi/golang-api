package main

import (
	"fmt"
) 

// Vehicle interface
type Vehicle interface {
	startEngine()
}

type TwoWheeler struct {
	Name string
}

type FourWheeler struct {
	Name string
}

func (vehicle TwoWheeler) startEngine() {
	fmt.Printf("%s engine is started ...\n", vehicle.Name)
}

func (vehicle FourWheeler) startEngine() {
	fmt.Printf("%s engine is started ...\n", vehicle.Name)
}

func main() {
	var vehicle1 TwoWheeler
	vehicle1.Name = "Two"
	vehicle1.startEngine()

	var vehicle2 FourWheeler
	vehicle2.Name = "Four"
	vehicle2.startEngine()
}