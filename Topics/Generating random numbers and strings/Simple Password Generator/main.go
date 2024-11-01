package main

import (
	"fmt"
	"math/rand"
)

func main() {
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	number := "0123456789"
	specialSymbol := "!?$&%#"

	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)

	// put your code here
	charSets := []string{upperCase, lowerCase, number, specialSymbol}
	outString := []uint8{}
	for i := range charSets {
		randIndex := rand.Intn(len(charSets[i]))
		outString = append(outString, charSets[i][randIndex])

	}
	fmt.Println(string(outString))
}
