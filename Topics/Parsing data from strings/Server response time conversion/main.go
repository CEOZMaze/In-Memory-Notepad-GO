package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// DO NOT delete or modify the code block below!
	var logEntry string
	fmt.Scanln(&logEntry)
	logEntry = logEntry[:len(logEntry)-2]

	// Write the additional code to convert the `responseTime` to float64:
	responseTime, err := strconv.ParseFloat(logEntry, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the response time to seconds below:
	responseTime = responseTime / 1000.0

	// Print the converted response time in seconds with two decimal points.
	fmt.Printf("Response time: %.2f seconds", responseTime)
}
