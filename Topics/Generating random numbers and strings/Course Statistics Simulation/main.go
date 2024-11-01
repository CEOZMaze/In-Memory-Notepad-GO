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
	gradeBook := make([]float64, 10)
	totalGrade := 0.0
	for i, _ := range gradeBook {
		grade := rand.Float64() * 100
		gradeBook[i] = grade

		totalGrade += gradeBook[i]

	}
	averageGrade := totalGrade / float64(len(gradeBook))
	fmt.Printf("%.2f\n", averageGrade)
}
