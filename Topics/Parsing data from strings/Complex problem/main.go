package main

import (
	"fmt"
	"strconv"
)

const bitSize = 128

func main() {
	var data string
	fmt.Scanln(&data)

	num, err := strconv.ParseComplex(data, bitSize)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("value: %v, type: %[1]T", num)
}
