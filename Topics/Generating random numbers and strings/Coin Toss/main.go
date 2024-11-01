package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	// put your code here
	num := rand.Float32()
	if num < 0.5 {
		fmt.Println("TAILS")
	} else if num > 0.5 {
		fmt.Println("HEADS")
	}
}
