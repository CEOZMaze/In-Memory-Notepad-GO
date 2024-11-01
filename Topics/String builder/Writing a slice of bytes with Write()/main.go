package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	// call the Write() function two times with the slice of bytes containing the phrases:
	builder.Write([]byte("Aye, carumba!"))

	builder.Write([]byte("Eat my shorts!"))

	// print the built string below
	fmt.Println(builder.String())
}
