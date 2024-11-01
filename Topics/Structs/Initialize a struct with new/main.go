package main

import (
	"fmt"
	"reflect"
)

type Pokemon struct {
	Name   string
	Number int
	Level  int
}

// Initialize the `pikachu` struct using the `new()` function below
func main() {
	pikachu := new(Pokemon)

	isCreatedWithNew(pikachu) // DO NOT delete this line!

	// DO NOT delete the code lines below!
	_ = fmt.Print
	_ = reflect.Struct
}
